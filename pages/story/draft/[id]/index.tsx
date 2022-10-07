import * as React from 'react'
import { Story } from '../../../../components/stories/story'
import { useStoryQuery } from '../../../../lib/schema/graphql'
import Head from 'next/head'
import { useRouter } from 'next/router'
import { DraftToolbar } from '../../../../components/stories/draftToolbar'
import ErrorMessage from '../../../../components/lib/error'
import LoaderSpinner from '../../../../components/lib/spinner'
import InnerContainer from '../../../../components/lib/innerContainer'

const StoryDraft = () => {
  const router = useRouter()
  const id = router.query.id as string
  const { data, error, loading } = useStoryQuery({
    variables: { id, draft: true },
  })

  if (error) return <ErrorMessage error={error} />
  if (loading || !data) return <LoaderSpinner />

  const story = data.story

  return (
    <InnerContainer>
      <Head>
        <title>Kladd - {story.name}</title>
      </Head>
      <DraftToolbar
        onSave={() => router.push(`/story/draft/${story.id}/save`)}
      />
      <div className="mt-12 flex gap-5 flex-col">
        <Story story={story} draft={true} />
      </div>
    </InnerContainer>
  )
}

export default StoryDraft
