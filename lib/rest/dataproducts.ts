import { useEffect, useState } from "react";
import { getDataproductUrl } from "./restApi";

const getDataproduct = async (id: string) => {
    const url = getDataproductUrl(id);
    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    }
    return fetch(url, options)
}

export const useGetDataproduct = (id: string)=>{
    const [dataproduct, setDataproduct] = useState<any>(null)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState(null)

    useEffect(()=>{
        if(!id) return
        getDataproduct(id).then((res)=> res.json())
        .then((dataproduct)=>
        {
            setError(null)
            setDataproduct(dataproduct)
        })
        .catch((err)=>{
            setError(err)
            setDataproduct(null)            
        }).finally(()=>{
            setLoading(false)
        })
    }, [id])

    return {dataproduct, loading, error}
}