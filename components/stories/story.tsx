import { InView } from 'react-intersection-observer'
import { StoryQuery } from '../../lib/schema/graphql'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import Header from './header'
import LoaderSpinner from '../lib/spinner'
import TopBar from '../lib/topBar'
import * as React from 'react'
import Link from 'next/link'
import { Dispatch, SetStateAction } from 'react'
import VegaView from './vegaView'
import { MetadataTable } from './metadataTable'
import dynamic from 'next/dynamic'

interface StoryProps {
  isOwner?: boolean
  story: StoryQuery['story']
  draft?: boolean
  setShowDelete?: Dispatch<SetStateAction<boolean>>
  setShowToken?: Dispatch<SetStateAction<boolean>>
}

export function Story({
  story,
  draft,
  isOwner,
  setShowDelete,
  setShowToken,
}: StoryProps) {
  const views = story.views as StoryQuery['story']['views']
  const Plotly = dynamic(() => import("./plotly"));

  return (
    <>
      <TopBar type={'Story'} name={story.name}>
        {!draft && isOwner && (
          <div>
            <Link 
              href={`/story/${story.id}/edit`}>
              <a className="pr-2">Endre</a>
            </Link>
            <a className="border-l-[1px] border-border px-2" onClick={() => setShowToken && setShowToken(true)}>Vis token</a>
            <a
              className="border-l-[1px] border-border px-2 text-nav-red"
              onClick={() => setShowDelete && setShowDelete(true)}
            >
              Slett
            </a>
          </div>
        )}
      </TopBar>

      <MetadataTable
        created={story.created}
        lastModified={story.lastModified}
        owner={story.owner}
        keywords={story.keywords}
      />
      <div className="children-fullwidth flex flex-wrap flex-row justify-between gap-5 py-4">
        {views.map((view, id) => {
          if (view.__typename === 'StoryViewHeader') {
            return <Header key={id} text={view.content} size={view.level} />
          }

          if (view.__typename === 'StoryViewMarkdown') {
            return (
              <ReactMarkdown key={id} remarkPlugins={[remarkGfm]}>
                {view.content}
              </ReactMarkdown>
            )
          }
          if (view.__typename === 'StoryViewPlotly') {
            return (
              <InView key={id} triggerOnce={true}>
                {({ inView, ref }) => {
                  return inView ? (
                    <div ref={ref}>
                      <Plotly id={view.id} draft={draft} />
                    </div>
                  ) : (
                    <div ref={ref}>
                      <LoaderSpinner />
                    </div>
                  )
                }}
              </InView>
            )
          }
          if (view.__typename === 'StoryViewVega') {
            return (
              <InView key={id} triggerOnce={true}>
                {({ inView, ref }) => {
                  return inView ? (
                    <div ref={ref}>
                      <VegaView id={view.id} draft={draft!!} />
                    </div>
                  ) : (
                    <div ref={ref}>
                      <LoaderSpinner />
                    </div>
                  )
                }}
              </InView>
            )
          }
        })}
      </div>
    </>
  )
}

export default Story
