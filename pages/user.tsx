import styled from 'styled-components'
import * as React from 'react'
import { useContext, useState } from 'react'
import { UserState } from '../lib/context'
import TopBar from '../components/lib/topBar'
import { Name } from '../components/lib/detailTypography'
import { Tab, Tabs } from '@mui/material'
import TabPanel from '../components/lib/tabPanel'
import { MetadataTable } from '../components/user/metadataTable'
import UserProductResultLink from '../components/user/userProductResult'
import UserAccessableProduct from '../components/user/userProductAccess'
import Head from 'next/head'

const StyledTabPanel = styled(TabPanel)`
  > div {
    padding-left: 0px;
    padding-right: 0px;
  }
`

export const UserProductLink = () => {
  const userState = useContext(UserState)
  const [activeTab, setActiveTab] = useState(0)
  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue)
  }

  if (!userState)
    return (
      <div>
        <h1>Du må være logget inn!</h1>
        <p>Bruk login-knappen øverst.</p>
      </div>
    )

  return (
    <div>
      <Head>
        <title>Brukerside</title>
      </Head>
      <TopBar type={'User'}>
        <Name>{userState.name}</Name>
      </TopBar>
      {userState.groups && <MetadataTable user={userState} />}
      <Tabs
        value={activeTab}
        onChange={handleChange}
        variant="standard"
        scrollButtons="auto"
        aria-label="auto tabs example"
      >
        <Tab label="Mine produkter og samlinger" value={0} />
        <Tab label="Mine tilganger" value={1} />
      </Tabs>
      <StyledTabPanel index={0} value={activeTab}>
        <UserProductResultLink />
      </StyledTabPanel>
      <StyledTabPanel index={1} value={activeTab}>
        <UserAccessableProduct />
      </StyledTabPanel>
    </div>
  )
}

export default UserProductLink
