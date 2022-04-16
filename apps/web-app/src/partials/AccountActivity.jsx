import React, { useEffect, useState } from 'react'
import { ResponsiveTimeRange } from '@nivo/calendar'
import { API_HOST } from '../App'

const exampleData = [{ day: "2022-01-01", value: 30 }, { day: "2022-02-10", value: 1 }, { day: "2022-02-01", value: 15 }]

const TimeRange = ({ data /* see data tab */ }) => (
    <ResponsiveTimeRange
        data={data}
        from="2021-07-22"
        to="2022-02-22"
        emptyColor="#eeeeee"
        colors={["#7ad151", "#22a884", "#2a788e", "#414487", "#440154"]}
        margin={{ top: 0, right: 0, bottom: 0, left: 0 }}
        weekdayLegendOffset={0}
        weekdayTicks={[]}
        dayRadius={0}
        dayBorderWidth={0}
        dayBorderColor="#ffffff"
        daySpacing={2}
        tooltip={(t) => <p className="bg-slate-800 text-white">
            {t.day} <br />
            Videos: {t.value}
        </p>}
        legends={[]}
    />

)

const AccountActivity = ({ account }) => {
    const [data, setData] = useState(null)

    useEffect(() => {
        try {
            fetch(API_HOST + "/" + account + "/activity")
                .then((res) => res.json())
                .then((json) => {
                    setData(json)
                })
        } catch (e) {
            console.log(e)
        }

    }, [])

    return data ? <TimeRange data={data} /> : <TimeRange data={exampleData} />
}

export default AccountActivity