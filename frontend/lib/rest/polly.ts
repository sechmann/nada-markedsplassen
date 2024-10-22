import { QueryPolly } from "./generatedDto"
import { fetchTemplate, HttpError } from "./request"
import { buildUrl } from "./apiUrl"
import { useQuery } from "react-query"

const pollyPath = buildUrl('polly')
const buildSearchPollyUrl = (query?: string) => pollyPath()({query: query || ''})

const searchPolly = async (query?: string) => 
    fetchTemplate(buildSearchPollyUrl(query))


export const useSearchPolly = (query?: string) => useQuery<QueryPolly[], HttpError>(['polly', query], ()=>searchPolly(query))
