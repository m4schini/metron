import React, { useEffect, useState } from 'react'
import { API_HOST, clean } from '../App'
import Video from './Video'

const Videos = ({ account }) => {
    const [data, setData] = useState(null)

    useEffect(() => {
        try {
            fetch(API_HOST + "/" + account + "/videos")
                .then((res) => res.json())
                .then((json) => {

                    const d = json.map((v) => {
                        const desc = clean(v.description)

                        console.log(desc)
                        v.description = desc
                        return v
                    })
                    console.log(d)

                    setData(d.reverse())
                })

        } catch (e) {
            console.log(e)
        }

    }, [])

    if (!data) {
        return <div></div>
    }

    return <div className="grid grid-cols-1 gap-4 my-4">
        {data && data.map((v) => {
            return <Video href={"/" + account + "/" + v.id} data={v} />
        })}
    </div>
}

export default Videos