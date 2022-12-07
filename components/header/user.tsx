import React, { MouseEvent, useContext } from 'react'
import { useRouter } from 'next/router'
import { UserState } from '../../lib/context'
import { Dropdown, Header } from '@navikt/ds-react-internal'
import { PeopleInCircle } from '@navikt/ds-icons'


export const backendHost = () => {
  return process.env.NODE_ENV === 'development' ? 'http://localhost:8080' : ''
}

export default function User() {
  const userInfo = useContext(UserState)

  const router = useRouter()

  return userInfo ? (
    <Dropdown>
      <Header.Button className="whitespace-nowrap hidden md:block" as={Dropdown.Toggle}>
        {userInfo.name}
      </Header.Button>
      <Header.Button className="block md:hidden" as={Dropdown.Toggle}>
        <PeopleInCircle />
      </Header.Button>
      <Dropdown.Menu className="w-fit">
        <Dropdown.Menu.GroupedList>
          <Dropdown.Menu.GroupedList.Item>
            <a
              className={'text-base'}
              onClick={() => {
                router.push({ pathname: '/user/products' })
              }}
            >
              Mine produkter
            </a>
          </Dropdown.Menu.GroupedList.Item>
          <Dropdown.Menu.GroupedList.Item>
            <a
              className={'text-base'}
              onClick={() => {
                router.push({ pathname: '/user/stories' })
              }}
            >
              Mine fortellinger
            </a>
          </Dropdown.Menu.GroupedList.Item>
          <Dropdown.Menu.GroupedList.Item>
            <a
              className={'text-base'}
              onClick={() => {
                router.push({ pathname: '/user/requests' })
              }}
            >
              Mine tilgangssøknader
            </a>
          </Dropdown.Menu.GroupedList.Item>
          <Dropdown.Menu.GroupedList.Item>
            <a
              className={'text-base'}
              onClick={() => {
                router.push({ pathname: '/user/access' })
              }}
            >
              Mine tilganger
            </a>
          </Dropdown.Menu.GroupedList.Item>
        </Dropdown.Menu.GroupedList>
        <Dropdown.Menu.Divider />
        <Dropdown.Menu.GroupedList>
          <Dropdown.Menu.GroupedList.Item>
            <a className={'text-base'} href={`${backendHost()}/api/logout`}>
              Logg ut
            </a>
          </Dropdown.Menu.GroupedList.Item>
        </Dropdown.Menu.GroupedList>
      </Dropdown.Menu>
    </Dropdown>
  ) : (
    <Header.Button
      className={'h-full'}
      onClick={async () =>
        await router.push(
          `${backendHost()}/api/login?redirect_uri=${encodeURIComponent(
            router.asPath
          )}`
        )
      }
      key="logg-inn"
    >
      Logg inn
    </Header.Button>
  )
}
