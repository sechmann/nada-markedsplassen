import React, { useContext, useState } from 'react'
import Link from 'next/link'
import styled from 'styled-components'
import User from './user'
import HeaderLogo from '../lib/icons/headerLogo'
import SearchBox from '../search/searchBox'
import { SearchState } from '../../pages/_app'
import {useRouter} from "next/router";

const HeaderBar = styled.header`
  display: flex;
  margin: 40px 0;
  justify-content: space-between;
`
export interface UserData {
  name: string,
  teams: string[]
}

export default function Header() {
  const [userData, setUserData] = useState({})
  const searchState = useContext(SearchState)

  const router = useRouter()
    console.log(router.pathname)
  return (
    <HeaderBar role="banner">
      <Link href="/">
        <div>
          <HeaderLogo />
        </div>
      </Link>

      { router.pathname !== '/' && <SearchBox /> }
      <User user={userData} />
    </HeaderBar>
  )
}
