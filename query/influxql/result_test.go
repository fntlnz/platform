package influxql_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"time"

	"github.com/influxdata/platform/query"
	"github.com/influxdata/platform/query/execute"
	"github.com/influxdata/platform/query/execute/executetest"
	"github.com/influxdata/platform/query/influxql"
)

func TestMultiResultEncoder_Encode(t *testing.T) {
	for _, tt := range []struct {
		name string
		in   query.ResultIterator
		out  string
	}{
		{
			name: "Default",
			in: query.NewSliceResultIterator(
				[]query.Result{&executetest.Result{
					Nm: "0",
					Blks: []*executetest.Block{{
						KeyCols: []string{"_measurement", "host"},
						ColMeta: []query.ColMeta{
							{Label: "time", Type: query.TTime},
							{Label: "_measurement", Type: query.TString},
							{Label: "host", Type: query.TString},
							{Label: "value", Type: query.TFloat},
						},
						Data: [][]interface{}{
							{ts("2018-05-24T09:00:00Z"), "m0", "server01", float64(2)},
						},
					}},
				}},
			),
			out: `{"results":[{"statement_id":0,"series":[{"name":"m0","tags":{"host":"server01"},"columns":["time","value"],"values":[["2018-05-24T09:00:00Z",2]]}]}]}`,
		},
		{
			name: "Error",
			in:   &resultErrorIterator{Error: "expected"},
			out:  `{"error":"expected"}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			enc := influxql.NewMultiResultEncoder()
			if err := enc.Encode(&buf, tt.in); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if got, exp := strings.TrimSpace(buf.String()), tt.out; got != exp {
				t.Fatalf("unexpected output:\nexp=%s\ngot=%s", exp, got)
			}
		})
	}
}

type resultErrorIterator struct {
	Error string
}

func (*resultErrorIterator) Cancel()            {}
func (*resultErrorIterator) More() bool         { return false }
func (*resultErrorIterator) Next() query.Result { panic("no results") }

func (ri *resultErrorIterator) Err() error {
	return errors.New(ri.Error)
}

func mustParseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

// ts takes an RFC3339 time string and returns an execute.Time from it using the unix timestamp.
func ts(s string) execute.Time {
	return execute.Time(mustParseTime(s).UnixNano())
}
