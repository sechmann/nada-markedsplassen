import * as React from 'react'
import { Select } from '@navikt/ds-react'
import { useContext } from 'react'
import { UserState } from '../../lib/context'
import { UserInfoDetailsQuery } from '../../lib/schema/graphql'

type TeamSelectorProps = {
  register: any
  errors: any
}
export const TeamSelector = ({ register, errors }: TeamSelectorProps) => {
  const user = useContext<UserInfoDetailsQuery['userInfo'] | undefined>(
    UserState
  )
  const teams = [...new Set(user?.gcpProjects.map((p) => p.group.name))]

  return (
    <Select
      label="Team"
      {...register('group')}
      error={errors.owner?.group?.message}
    >
      <option value="">Velg team</option>
      <option value="test@nav.no" key="test">test</option>
      {teams.map((team, i) => (
        <option
          value={user?.groups.filter((g) => g.name === team)[0].email}
          key={team}
        >
          {team}
        </option>
      ))}
    </Select>
  )
}
export default TeamSelector
