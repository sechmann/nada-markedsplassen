import { useEffect, useState } from "react";
import { AccessRequestsWrapper, GrantAccessData, NewAccessRequestDTO, UpdateAccessRequestDTO } from "./generatedDto";
import { deleteTemplate, fetchTemplate, postTemplate, putTemplate } from "./request";
import { buildPath } from "./apiUrl";

const accessRequestsPath = buildPath('accessRequests')
const buildFetchAccessRequestUrl = (datasetId: string) => accessRequestsPath()({datasetId: datasetId})
const buildCreateAccessRequestUrl = () => accessRequestsPath('new')()
const buildDeleteAccessRequestUrl = (id: string) => accessRequestsPath(id)()
const buildUpdateAccessRequestUrl = (id: string) => accessRequestsPath(id)()

const processAccessRequestsPath = buildPath('accessRequests/process')
const buildApproveAccessRequestUrl = (accessRequestId: string) => processAccessRequestsPath(accessRequestId)({action: 'approve'})
const buildDenyAccessRequestUrl = (accessRequestId: string, reason: string) => processAccessRequestsPath(accessRequestId)({action: 'deny', reason: reason})

const accessPath = buildPath('accesses')
const buildGrantAccessUrl = () => accessPath('grant')()
const buildRevokeAccessUrl = (accessId: string) => accessPath('revoke')({accessId: accessId})

export enum SubjectType {
    Group = 'group',
    ServiceAccount = 'serviceAccount',
    User = 'user'
}

export const fetchAccessRequests = async (datasetId: string) => 
    fetchTemplate(buildFetchAccessRequestUrl(datasetId))

export const createAccessRequest = async (newAccessRequest: NewAccessRequestDTO) => 
    postTemplate(buildCreateAccessRequestUrl(), newAccessRequest)

export const deleteAccessRequest = async (id: string) =>
    deleteTemplate(buildDeleteAccessRequestUrl(id))

export const updateAccessRequest = async (updateAccessRequest: UpdateAccessRequestDTO) => 
    putTemplate(buildUpdateAccessRequestUrl(updateAccessRequest.id), updateAccessRequest)

export const apporveAccessRequest = async (accessRequestId: string) => 
    postTemplate(buildApproveAccessRequestUrl(accessRequestId))

export const denyAccessRequest = async (accessRequestId: string, reason: string) => 
    postTemplate(buildDenyAccessRequestUrl(accessRequestId, reason))

export const grantDatasetAccess = async (grantAccess: GrantAccessData) => 
    postTemplate(buildGrantAccessUrl(), grantAccess)

export const revokeDatasetAccess = async (accessId: string) => 
    postTemplate(buildRevokeAccessUrl(accessId))

export const useFetchAccessRequestsForDataset = (datasetId: string)=>{
    const [data, setData] = useState<AccessRequestsWrapper| null>(null)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState(null)


    useEffect(()=>{
        if(!datasetId) return
        fetchAccessRequests(datasetId)
        .then((data)=>
        {
            setError(null)
            setData(data)
        })
        .catch((err)=>{
            setError(err)
            setData(null)            
        }).finally(()=>{
            setLoading(false)
        })
    }, [datasetId])

    return {data: data?.accessRequests, loading, error}
}