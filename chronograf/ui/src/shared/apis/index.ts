import AJAX from 'src/utils/ajax'
import {Source, Service, NewService, QueryConfig} from 'src/types'

export const getSources = async (): Promise<Source[]> => {
  try {
    const {data} = await AJAX({
      url: '/v2/sources',
    })

    return data.sources
  } catch (error) {
    throw error
  }
}

export const getSource = async (url: string): Promise<Source> => {
  try {
    const {data: source} = await AJAX({
      url,
    })

    return source
  } catch (error) {
    throw error
  }
}

export const createSource = async (
  url: string,
  attributes: Partial<Source>
): Promise<Source> => {
  try {
    const {data: source} = await AJAX({
      url,
      method: 'POST',
      data: attributes,
    })

    return source
  } catch (error) {
    throw error
  }
}

export const updateSource = async (
  newSource: Partial<Source>
): Promise<Source> => {
  try {
    const {data: source} = await AJAX({
      url: newSource.links.self,
      method: 'PATCH',
      data: newSource,
    })

    return source
  } catch (error) {
    throw error
  }
}

export function deleteSource(source) {
  return AJAX({
    url: source.links.self,
    method: 'DELETE',
  })
}

export const getQueryConfigAndStatus = async (
  url,
  queries
): Promise<AnalyzeQueriesObject[]> => {
  try {
    const {data} = await AJAX({
      url,
      method: 'POST',
      data: {queries},
    })

    return data.queries
  } catch (error) {
    console.error(error)
    throw error
  }
}

interface AnalyzeQueriesObject {
  id: string
  query: string
  duration: string
  queryConfig?: QueryConfig
}

export const analyzeQueries = async (
  url: string,
  queries: Array<{query: string}>
): Promise<AnalyzeQueriesObject[]> => {
  try {
    const {data} = await AJAX({
      url,
      method: 'POST',
      data: {queries},
    })

    return data.queries
  } catch (error) {
    console.error(error)
    throw error
  }
}

export const getServices = async (url: string): Promise<Service[]> => {
  try {
    const {data} = await AJAX({
      url,
      method: 'GET',
    })

    return data.services
  } catch (error) {
    console.error(error)
    throw error
  }
}

export const getService = async (
  url: string,
  serviceID: string
): Promise<Service> => {
  try {
    const {data} = await AJAX({
      url: `${url}/${serviceID}`,
      method: 'GET',
    })

    return data
  } catch (error) {
    console.error(error)
    throw error
  }
}

export const createService = async (
  source: Source,
  {
    url,
    name = 'My FluxD',
    type,
    username,
    password,
    insecureSkipVerify,
    metadata,
  }: NewService
): Promise<Service> => {
  try {
    const {data} = await AJAX({
      url: source.links.services,
      method: 'POST',
      data: {url, name, type, username, password, insecureSkipVerify, metadata},
    })

    return data
  } catch (error) {
    console.error(error)
    throw error
  }
}

export const updateService = async (service: Service): Promise<Service> => {
  try {
    const {data} = await AJAX({
      url: service.links.self,
      method: 'PATCH',
      data: service,
    })

    return data
  } catch (error) {
    console.error(error)
    throw error
  }
}

export const deleteService = async (service: Service): Promise<void> => {
  try {
    await AJAX({
      url: service.links.self,
      method: 'DELETE',
    })
  } catch (error) {
    console.error(error)
    throw error
  }
}
