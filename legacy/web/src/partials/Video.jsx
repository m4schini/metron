import React from 'react'
import TinyCard from './TinyCard'
import { EyeIcon, ThumbUpIcon, ChatAltIcon, ShareIcon } from '@heroicons/react/outline'
import Start from '../pages/Start'
import aveta from 'aveta'

const Stats = ({ load, views, comments, shares, likes }) => {
    if (load) {
        return <div className="grid grid-cols-2 sm:flex gap-2 text-gray-500 dark:text-gray-400 text-xs my-2 ">
            <div className="border-4 rounded-2xl w-36 h-24 p-4 bg-gray-200 animate-pulse"></div>
            <div className="border-4 rounded-2xl w-36 h-24 p-4 bg-gray-200 animate-pulse"></div>
            <div className="border-4 rounded-2xl w-36 h-24 p-4 bg-gray-200 animate-pulse"></div>
            <div className="border-4 rounded-2xl w-36 h-24 p-4 bg-gray-200 animate-pulse"></div>
        </div>
    }

    return <div className="grid grid-cols-2 sm:flex gap-2 text-gray-500 dark:text-gray-400 text-xs my-2 ">
        <TinyCard icon={<EyeIcon className="block h-4 w-4" />} title={"Views"} content={aveta(views || "0")} />
        <TinyCard icon={<ThumbUpIcon className="block h-4 w-4" />} title={"Likes"} content={aveta(likes || "0")} />
        <TinyCard icon={<ChatAltIcon className="block h-4 w-4" />} title={"Comments"} content={aveta(comments || "0")} />
        <TinyCard icon={<ShareIcon className="block h-4 w-4" />} title={"Shares"} content={aveta(shares || "0")} />
    </div>
}

const Description = ({ data, load }) => {
    if (load) {
        return <div className="flex flex-1 flex-col gap-3 mb-2">
            <div className="bg-gray-200 w-full animate-pulse h-3 rounded-2xl">
            </div>
            <div className="bg-gray-200 w-full animate-pulse h-3 rounded-2xl">
            </div>
        </div>
    }

    return <div className="flex items-start justify-between text-gray-700 dark:text-white my-2 md:m-0">
        <p className="text-xl leading-5">
            {data && data.replace('\.[-a-zA-Z0-9]+{[a-zA-Z-:0-9;]+}', "")}
        </p>
    </div>
}

const Tags = ({ data, load }) => {
    if (load) {
        return <div className="flex items-center gap-2 my-2">
            <div className="bg-gray-200 w-20 h-6 animate-pulse rounded-full"></div>
            <div className="bg-gray-200 w-20 h-6 animate-pulse rounded-full"></div>
            <div className="bg-gray-200 w-20 h-6 animate-pulse rounded-full"></div>
        </div>
    }

    return <div className="flex flex-wrap items-center gap-2 my-2">
        {data.map((tagName) => {
            return <span className="px-2 py-1 text-xs font-bold rounded-full text-white  bg-indigo-500 ">
                {tagName}
            </span>
        })}
    </div>
}

const PostedAt = ({ data, load }) => {
    if (load) {
        return <div>
            <div className="bg-gray-200 w-16 animate-pulse h-3 rounded-2xl"></div>
        </div>
    }

    return <div className="flex items-start my-4 md:my-1">
        <div className="flex flex-col items-start justify-center ml-2">
            <span className="text-gray-400 text-xs">
                {data}
            </span>
        </div>
    </div>
}

const Thumbnail = ({ src, load }) => {
    if (load) {
        return <div className="hidden md:block h-52 sm:h-full w-36 rounded-xl bg-gray-200 animate-pulse"></div>
    }

    return <div className="relative">
        <img src={src} className="hidden md:block rounded-xl w-full md:w-40 md:h-auto" />
    </div>
}

const Video = ({ href, data }) => {
    return (

        <a key={data.id} href={href} className="p-4 bg-white shadow-xl rounded-xl flex justify-start dark:bg-gray-800 md:flex-row flex-col gap-4">
            <Thumbnail load={!data} src={"assets/video.jpg"} />
            <div className="flex flex-col justify-between">
                <Description data={data.description} load={!data} />
                <Tags data={[]} load={!data} />
                <Stats load={!data} views={data.views} comments={data.comments} shares={data.shares} likes={data.likes} />
                <PostedAt data={"nope"} load={!data} />

            </div>
        </a>

    )
}

export default Video