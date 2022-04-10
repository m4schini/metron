import React, { useState, useEffect } from 'react'
import HistoryLine from '../components/AccountHistoryLine'
import TagChart from '../components/TagChart'
import MentionGraph from '../components/MentionGraph'

import randomColor from 'randomcolor'
import aveta from 'aveta';

import { API_HOST } from '../App'



const Card = ({ small, big, children, height, className }) => {
    return (
        <article className={"shadow-lg rounded-lg px-4 py-6 bg-white dark:bg-gray-700 relative" + " " + className}>
            <div className={"text-black dark:text-white h-full " + (height || "max-h-56")}>
                {children}

            </div>
            <p className="text-2xl text-black dark:text-white font-bold">
                {big ? aveta(big) : big}
            </p>
            <p className="text-gray-400 text-sm">
                {small}
            </p>
        </article>
    )
}

export const AccountFollowing = ({ account }) => {
    return (
        <Card big={1095} small={"Following"} height="max-h-36">
            <HistoryLine data={account} />
        </Card>
    )
}

export const AccountFollowers = ({ account, limit }) => {
    const [accData, setAccData] = useState(null)

    useEffect(() => {
        try {
            fetch(API_HOST + "/" + account + "/followers?limit=" + (limit ? limit : ""))
                .then((res) => res.json())
                .then((json) => {
                    json.data.reverse()
                    setAccData(json)
                })
        } catch (e) {
            console.log(e)
        }

    }, [])

    return (
        <Card big={accData && (accData.data[0].y)} small={"Followers"} height="max-h-36">
            {accData && <HistoryLine data={accData} />}
        </Card>
    )
}

export const AccountLikes = ({ account, limit }) => {
    const [accData, setAccData] = useState(null)

    useEffect(() => {
        try {
            fetch(API_HOST + "/" + account + "/likes?limit=" + (limit ? limit : ""))
                .then((res) => res.json())
                .then((json) => {
                    json.data.reverse()
                    setAccData(json)
                })
        } catch (e) {
            console.log(e)
        }

    }, [])


    return (
        <Card big={accData && accData.data[0].y} small={"Likes"} height="max-h-36">
            {accData && <HistoryLine data={accData} />}
        </Card>
    )
}

export const AccountTags = ({ account }) => {
    const [data, setData] = useState(null)
    const [tagCount, setTagCount] = useState(0)

    useEffect(() => {
        try {
            fetch(API_HOST + "/" + account + "/tags")
                .then((res) => res.json())
                .then((json) => {
                    setTagCount(json.length)
                    json = json.filter((v) => v.value > 1)
                    setData(json)
                })
        } catch (e) {
            console.log(e)
        }

    }, [])

    return (
        <Card big={tagCount} small={"Tags"} height="max-h-36">
            {data ? <TagChart data={data} /> : <div>loading...</div>}
        </Card>
    )
}

export const AccountMentions = ({ account }) => {
    const [data, setData] = useState(null)
    const [mentionCount, setMentionCount] = useState(0)

    useEffect(() => {
        try {
            fetch(API_HOST + "/" + account + "/mentions")
                .then((res) => res.json())
                .then((json) => {

                    json.nodes = json.nodes.map((n) => {
                        return { color: (n?.id != account ? randomColor() : "#FFF"), ...n }
                    })
                    setMentionCount(json.links.length)

                    setData(json)
                })
        } catch (e) {
            console.log(e)
        }

    }, [])

    return (
        <Card big={mentionCount} small={"Mentions"} height="max-h-36">
            {data ? <MentionGraph data={data} /> : <div>loading...</div>}
        </Card>
    )
}


export default Card