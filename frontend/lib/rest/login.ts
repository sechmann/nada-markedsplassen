import { buildUrl } from "./apiUrl";

const loginPath= buildUrl('login')
export const buildLoginUrl = (redirect_uri:string) => loginPath()({redirect_uri: redirect_uri})