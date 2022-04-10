import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import Card, {
    AccountFollowers as Followers,
    AccountLikes as Likes,
    AccountTags as Tags,
    AccountMentions as Mentions
} from '../partials/Card'
import Videos from '../partials/Videos'
import MiniCard from '../partials/MiniCard'
import Video from '../partials/Video'
import Activity from '../partials/AccountActivity'


import { EyeIcon, ThumbUpIcon, ChatAltIcon, ShareIcon } from '@heroicons/react/outline'

const Dashboard = () => {
    const params = useParams()
    const accountName = params.account || "Krawallklara"

    const apiHost = "http://localhost:8080"

    const [accData, setAccData] = useState({})

    useEffect(() => {
        try {
            fetch(apiHost + "/" + accountName)
                .then((res) => res.json())
                .then((json) => setAccData(json))
        } catch (e) {
            console.log(e)
        }

    }, [])


    return (

        <main className="bg-gray-100 dark:bg-gray-800 h-full min-h-screen relative">
            <div className="flex items-center justify-center">
                <div className="flex flex-col w-full lg:w-3/4 xl:w-3/5 md:space-y-4 mt-4">
                    <div className="pb-24 px-4 md:px-6">
                        <header className="flex flex-col md:flex-row justify-between gap-4 md:mt-8 md:mb-8 h-48 md:h-24">

                            <div className="flex gap-4">
                                <div className="w-28">
                                    <img className="rounded-3xl" src="assets/accountAvatar.jpeg" alt="alt" />
                                </div>
                                <div>
                                    <h1 className="text-4xl font-semibold text-gray-800 dark:text-white">
                                        @{accountName.toUpperCase()}
                                    </h1>
                                    <h2 className="text-md text-gray-400">
                                        Last update: <strong>{accData?.lastUpdate}</strong>
                                    </h2>

                                </div>
                            </div>
                            <div className="h-24 w-full flex justify-end">
                                <div className="w-full max-w-md h-56">
                                    <Activity account={accountName} />
                                </div>
                            </div>
                        </header>
                        <h1 className="text-black dark:text-white text-lg font-bold">Overview</h1>
                        <div className="grid grid-cols-1 w-full  md:grid-cols-2 gap-4 my-4">
                            {/* <Following account={accountName} /> */}

                            <Followers account={accountName} limit={100} />

                            <Likes account={accountName} limit={100} />

                            <Tags account={accountName} />

                            <Mentions account={accountName} />

                        </div>

                        <h1 className="text-black dark:text-white text-lg font-bold">Summary <span className="text-sm font-light">({accData?.videos || 0} Videos)</span></h1>

                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 my-4">
                            <MiniCard
                                icon={<EyeIcon className="block h-4 w-4" />}
                                iconcolor={"blue"} title={"Views"}
                                loading={true}
                                content={accData?.summary?.views} />
                            <MiniCard
                                icon={<ThumbUpIcon className="block h-4 w-4" />}
                                iconcolor={"blue"} title={"Comments"}
                                content={accData?.summary?.comments} />
                            <MiniCard
                                icon={<ChatAltIcon className="block h-4 w-4" />}
                                iconcolor={"blue"} title={"Shares"}
                                content={accData?.summary?.shares} />
                            <MiniCard
                                icon={<ShareIcon className="block h-4 w-4" />}
                                iconcolor={"blue"} title={"Likes"}
                                content={accData?.summary?.likes} />
                        </div>

                        <div className="text-black dark:text-white text-lg font-bold">
                            <h1>Videos</h1>
                            <Videos account={accountName} />

                        </div>
                    </div>
                </div>
            </div>
        </main>

    )
}

export default Dashboard