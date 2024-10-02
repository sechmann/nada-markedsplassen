import { SearchOptions } from "./search"

const isServer = typeof window === 'undefined'

//Currying function for building urls
const buildUrl = (baseUrl: string)=>(path: string) => (pathParam?: string)=> (queryParams?: Record<string, string>) => 
  `${baseUrl}${path}${pathParam? `/${pathParam}`: ''}${queryParams? `?${new URLSearchParams(queryParams).toString()}`: ''}`

const localBaseUrl = buildUrl('http://localhost:8080/api')
const asServerBaseUrl = buildUrl('http://nada-backend/api')
const asClientBaseUrl = buildUrl('/api')

const baseUrl = ()=>{
    if (process.env.NEXT_PUBLIC_ENV === 'development') {
      return localBaseUrl
    }
    return isServer ? asServerBaseUrl : asClientBaseUrl
}

const dataproductUrl = baseUrl()('/dataproducts')
export const getDataproductUrl = (id: string) => dataproductUrl(id)()
export const createDataproductUrl = () => dataproductUrl('new')()
export const updateDataproductUrl = (id: string) => dataproductUrl(id)()
export const deleteDataproductUrl = (id: string) => dataproductUrl(id)()


const datasetUrl = baseUrl()('/datasets')
export const getDatasetUrl = (id: string) => datasetUrl(id)()
export const mapDatasetToServicesUrl = (datasetId: string) => `${datasetUrl(datasetId)()}/map`
export const createDatasetUrl = () => datasetUrl('new')()
export const deleteDatasetUrl = (id: string) => datasetUrl(id)()
export const updateDatasetUrl = (id: string) => datasetUrl(id)()
export const getAccessiblePseudoDatasetsUrl = () =>  datasetUrl('pseudo/accessible')()

const storyUrl = baseUrl()('/stories')
export const createStoryUrl = () => storyUrl('new')()
export const updateStoryUrl = (id: string) => storyUrl(id)()
export const deleteStoryUrl = (id: string) => storyUrl(id)()

const joinableViewUrl = baseUrl()('/pseudo/joinable')
export const getJoinableViewUrl = (id: string) => joinableViewUrl(id)()
export const createJoinableViewsUrl = () =>   joinableViewUrl('new')()
export const getJoinableViewsForUserUrl = () => joinableViewUrl()()


export const apiUrl = () => {
  if (process.env.NEXT_PUBLIC_ENV === 'development') {
    return 'http://localhost:8080/api'
  }
  return isServer ? 'http://nada-backend/api' : '/api'
}




export const getProductAreasUrl = () => `${apiUrl()}/productareas`
export const getProductAreaUrl = (id: string) => `${apiUrl()}/productareas/${id}`
export const fetchUserDataUrl = () => `${apiUrl()}/userData`

export const isValidSlackChannelUrl = (channel: string) => `${apiUrl()}/slack/isValid?channel=${channel}`

export const getInsightProductUrl = (id: string) => `${apiUrl()}/insightProducts/${id}`
export const createInsightProductUrl = () => `${apiUrl()}/insightProducts/new`
export const updateInsightProductUrl = (id: string) => `${apiUrl()}/insightProducts/${id}`
export const deleteInsightProductUrl = (id: string) => `${apiUrl()}/insightProducts/${id}`

export const fetchKeywordsUrl = () => `${apiUrl()}/keywords`
export const updateKeywordsUrl = () => `${apiUrl()}/keywords`

export const fetchAccessRequestUrl = (datasetId: string) => `${apiUrl()}/accessRequests?datasetId=${datasetId}`
export const fetchBQDatasetsUrl = (projectId: string) => `${apiUrl()}/bigquery/datasets?projectId=${projectId}`
export const fetchBQTablesUrl = (projectId: string, datasetId: string) => `${apiUrl()}/bigquery/tables?projectId=${projectId}&datasetId=${datasetId}`
export const fetchBQColumnsUrl = (projectId: string, datasetId: string, tableId: string) => `${apiUrl()}/bigquery/columns?projectId=${projectId}&datasetId=${datasetId}&tableId=${tableId}`
export const fetchStoryMetadataURL = (id: string) => `${apiUrl()}/stories/${id}`
export const searchTeamKatalogenUrl = (gcpGroups?: string[]) => {
  const parameters = gcpGroups?.length ? gcpGroups.map(group => `gcpGroups=${encodeURIComponent(group)}`).join('&') : ''
  const query = parameters ? `?${parameters}` : ''
  return `${apiUrl()}/teamkatalogen${query}`
}
export const searchPollyUrl = (query?: string) => {
  return `${apiUrl()}/polly?query=${query}`
}

export const createAccessRequestUrl = () => `${apiUrl()}/accessRequests/new`
export const deleteAccessRequestUrl = (id: string) => `${apiUrl()}/accessRequests/${id}`
export const updateAccessRequestUrl = (id: string) => `${apiUrl()}/accessRequests/${id}`
export const approveAccessRequestUrl = (accessRequestId: string) => `${apiUrl()}/accessRequests/process/${accessRequestId}?action=approve`
export const denyAccessRequestUrl = (accessRequestId: string, reason: string) => `${apiUrl()}/accessRequests/process/${accessRequestId}?action=deny&reason=${reason}`
export const grantAccessUrl = () => `${apiUrl()}/accesses/grant`
export const revokeAccessUrl = (accessId: string) => `${apiUrl()}/accesses/revoke?id=${accessId}`

export const getWorkstationURL = () => `${apiUrl()}/workstations/`
export const ensureWorkstationURL = () => `${apiUrl()}/workstations/`
export const startWorkstationURL = () => `${apiUrl()}/workstations/start`
export const stopWorkstationURL = () => `${apiUrl()}/workstations/stop`

export const fetchTemplate = (url: string) => fetch(url, {
  method: 'GET',
  credentials: 'include',
  headers: {
    'Content-Type': 'application/json',
  },
}).then(res => {
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  return res
})

export const postTemplate = (url: string, body?: any) => fetch(url, {
  method: 'POST',
  credentials: 'include',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(body),
}).then(async res => {
  if (!res.ok) {
    const errorMessage = await res.text()
    throw new Error(`${res.statusText}${errorMessage&&":"}${errorMessage}`)
  }
  return res
})

export const putTemplate = (url: string, body?: any) => fetch(url, {
  method: 'PUT',
  credentials: 'include',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(body),
}).then(res => {
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  return res
})

export const deleteTemplate = (url: string, body?: any) => fetch(url, {
  method: 'DELETE',
  credentials: 'include',
  headers: {
    'Content-Type': 'application/json',
  },
}).then(res => {
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  return res
})

export const searchUrl = (options: SearchOptions) => {
  let queryParams: string[] = [];

  // Helper function to add array-based options
  const addArrayOptions = (optionArray: string[] | undefined, paramName: string) => {
    if (optionArray && optionArray.length) {
      queryParams.push(`${paramName}=${optionArray.reduce((s, p) => `${s ? `${s},` : ""}${encodeURIComponent(p)}`)}`)
    }
  };

  // Adding array-based options
  addArrayOptions(options.keywords, 'keywords');
  addArrayOptions(options.groups, 'groups');
  addArrayOptions(options.teamIDs, 'teamIDs');
  addArrayOptions(options.services, 'services');
  addArrayOptions(options.types, 'types');

  // Adding single-value options
  if (options.text) queryParams.push(`text=${encodeURIComponent(options.text)}`);
  if (options.limit !== undefined) queryParams.push(`limit=${options.limit}`);
  if (options.offset !== undefined) queryParams.push(`offset=${options.offset}`);

  const query = queryParams.length ? `?${queryParams.join('&')}` : '';
  return `${apiUrl()}/search${query}`;
};

