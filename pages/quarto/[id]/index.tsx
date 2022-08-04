import {useRouter} from "next/router";
import {useQuartoQuery} from "../../../lib/schema/graphql";
import ErrorMessage from "../../../components/lib/error";
import LoaderSpinner from "../../../components/lib/spinner";
import * as React from "react";

const QuartoPage = () => {
    const router = useRouter()
    const id = router.query.id as string

    const query = useQuartoQuery({ variables: { id } })

    if (query.error) return <ErrorMessage error={query.error}/>
    if (query.loading || !query.data) return <LoaderSpinner/>

    const quarto = query.data.quarto.content

    return <div dangerouslySetInnerHTML={{__html: quarto}}/>
}

export default QuartoPage