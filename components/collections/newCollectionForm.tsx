import { Fieldset, Select, TextField } from '@navikt/ds-react'
import { useForm } from 'react-hook-form'
import { useContext, useEffect, useState } from 'react'
import { newCollectionValidation } from '../../lib/schema/yupValidations'
import { yupResolver } from '@hookform/resolvers/yup'
import RightJustifiedSubmitButton from '../widgets/formSubmit'
import { UserState } from '../../lib/context'
import KeywordsInput from '../lib/KeywordsInput'
import { useMutation } from '@apollo/client'
import { CREATE_DATAPRODUCT } from '../../lib/queries/dataproduct/createDataproduct'
import { CREATE_COLLECTION } from '../../lib/queries/collection/newCollection'

interface NewDatacollectionFormProps {
  onSubmit: (data: any) => Promise<void>
  onCancel: () => void
}

export const NewCollectionForm = ({
  onSubmit,
  onCancel,
}: NewDatacollectionFormProps) => {
  const [keywords, setKeywords] = useState<string[]>([])
  const { register, handleSubmit, formState, setValue } = useForm({
    resolver: yupResolver(newCollectionValidation),
  })
  // propagates change in keywords state to useForm state
  useEffect(() => setValue('keywords', keywords), [keywords])
  const { errors } = formState

  const groups = useContext(UserState)?.groups

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Fieldset legend="Samling" errorPropagation={false}>
        <TextField
          id="name"
          label="Navn"
          {...register('name')}
          error={errors.name?.message}
        />
        <TextField
          id="description"
          label="Beskrivelse"
          {...register('description')}
          error={errors.description?.message}
        />
        <Select
          label="Team"
          {...register('group')}
          error={errors.group?.message}
        >
          <option value="">Velg team</option>
          {groups?.map((group) => (
            <option value={group.email} key={group.name}>
              {group.name}
            </option>
          ))}
        </Select>
        <KeywordsInput
          keywords={keywords}
          setKeywords={setKeywords}
          {...register('keywords')}
          error={errors.keywords?.[0].message}
        />
      </Fieldset>
      <RightJustifiedSubmitButton onCancel={onCancel} />
    </form>
  )
}
