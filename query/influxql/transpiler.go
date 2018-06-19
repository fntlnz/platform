package influxql

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/influxdata/influxql"
	"github.com/influxdata/platform/query"
	"github.com/influxdata/platform/query/functions"
)

// Transpiler converts InfluxQL queries into a query spec.
type Transpiler struct{}

func NewTranspiler() *Transpiler {
	return new(Transpiler)
}

func (t *Transpiler) Transpile(ctx context.Context, txt string, config interface{}) (*query.Spec, error) {
	// Parse the text of the query.
	q, err := influxql.ParseQuery(txt)
	if err != nil {
		return nil, err
	}

	cfg, _ := config.(*Config)
	transpiler := newTranspilerState(cfg)
	for i, s := range q.Statements {
		stmt, ok := s.(*influxql.SelectStatement)
		if !ok {
			// TODO(jsternberg): Support meta queries.
			return nil, errors.New("only supports select statements")
		} else if err := transpiler.Transpile(ctx, i, stmt); err != nil {
			return nil, err
		}
	}
	return transpiler.spec, nil
}

type transpilerState struct {
	id     int
	stmt   *influxql.SelectStatement
	config Config
	spec   *query.Spec
	nextID map[string]int
	now    time.Time
}

func newTranspilerState(config *Config) *transpilerState {
	state := &transpilerState{
		spec:   &query.Spec{},
		nextID: make(map[string]int),
		now:    time.Now(),
	}
	if config != nil {
		state.config = *config
	}
	return state
}

func (t *transpilerState) Transpile(ctx context.Context, id int, stmt *influxql.SelectStatement) error {
	// Clone the select statement and omit the time from the list of column names.
	t.stmt = stmt.Clone()
	t.stmt.OmitTime = true
	t.id = id

	groups := identifyGroups(t.stmt)
	if len(groups) == 0 {
		return errors.New("no fields")
	}

	cursors := make([]cursor, 0, len(groups))
	for _, gr := range groups {
		cur, err := gr.createCursor(t)
		if err != nil {
			return err
		}
		cursors = append(cursors, cur)
	}

	// Join the cursors together on the measurement name.
	// TODO(jsternberg): This needs to join on all remaining partition keys.
	cur := Join(t, cursors, []string{"_measurement"}, nil)

	// Map each of the fields into another cursor. This evaluates any lingering expressions.
	cur, err := t.mapFields(cur)
	if err != nil {
		return err
	}

	// Yield the cursor from the last cursor to a stream with the name of the statement id.
	// TODO(jsternberg): Include the statement id in the transpiler state when we create
	// the state so we can yield to something other than zero.
	t.op("yield", &functions.YieldOpSpec{Name: strconv.Itoa(t.id)}, cur.ID())
	return nil
}

func (t *transpilerState) mapType(ref *influxql.VarRef) influxql.DataType {
	// TODO(jsternberg): Actually evaluate the type against the schema.
	return influxql.Tag
}

func (t *transpilerState) op(name string, spec query.OperationSpec, parents ...query.OperationID) query.OperationID {
	op := query.Operation{
		ID:   query.OperationID(fmt.Sprintf("%s%d", name, t.nextID[name])),
		Spec: spec,
	}
	t.spec.Operations = append(t.spec.Operations, &op)
	for _, pid := range parents {
		t.spec.Edges = append(t.spec.Edges, query.Edge{
			Parent: pid,
			Child:  op.ID,
		})
	}
	t.nextID[name]++
	return op.ID
}
