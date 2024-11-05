import { Alert, CopyButton, Link } from "@navikt/ds-react"
import { useRouter } from "next/router"
import { buildLoginUrl } from "../../lib/rest/login"

export interface ErrorStripeProps {
    error?: any
}

const defaultMessage = new Map([
    [403, "Forespørselen er forbudt."],
    [404, "Fant ikke ressursen på serveren."],
    [500, "Noe gikk galt på serveren."],
    [502, "Noe gikk galt på serveren."],
    [503, "Noe gikk galt på serveren."],
    [504, "Noe gikk galt på serveren."],
])

const LoginStripe = () => {
    const router = useRouter()
    return <Alert variant="info" className="w-full">
        <div className="pl-2">Du må <Link href="#" onClick={async () =>
          await router.push(buildLoginUrl(router.asPath))
        }>logge inn</Link> for å fortsette</div>
    </Alert>
}

const ErrorStripe = ({ error }: ErrorStripeProps) => {
    const message = /*error?.message ||*/ defaultMessage.get(error?.status) || "Noe gikk galt"
    const id = error?.id || undefined

    if(error?.status === 401) {
        return <LoginStripe />
    }
    
    return error && <Alert variant="error" className="w-full">
        <div className="pl-2">{message}</div>
        <p className="items-end">Du kan prøve igjen eller kontakte
            <Link className="pl-2 pr-2" href="https://navikt.slack.com/archives/CGRMQHT50" target="_blank">#nada på Slack</Link>
            for støtte hvis feilen vedvarer.
        </p>
        {id && <div className="flex">
            Vennligst legg ved Feil-ID <p className="flex pl-2 pr-2 text-nav-red">{id}<CopyButton copyText={id} size="small"></CopyButton></p> i meldingen når du kontakter oss.
        </div>}
    </Alert>
}

export default ErrorStripe