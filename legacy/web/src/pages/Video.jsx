import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'

import Card from '../partials/Card'
import { ResponsiveLine } from '@nivo/line'
import { API_HOST, clean } from '../App'

const LineChart = ({ data /* see data tab */ }) => (
    <ResponsiveLine
        data={data}
        margin={{ top: 0, right: 0, bottom: 40, left: 60 }}
        xScale={{ type: 'point' }}
        yScale={{
            type: 'linear',
            min: '0',
            max: 'auto',
            stacked: false,
            reverse: false
        }}
        yFormat=" >-.2f"
        axisTop={null}
        axisRight={null}
        axisBottom={null}
        axisLeft={{
            orient: 'left',
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: 'views',
            legendOffset: -56,
            legendPosition: 'middle'
        }}
        colors={["#1C1C1C", "#596aff", "#b5b5b5"]}
        enablePoints={false}
        areaOpacity={1}
        useMesh={true}
        legends={[]}
    />
)

const Video = () => {
    const params = useParams()
    const account = params.account
    const video = params.video

    const [data, setData] = useState(null)

    useEffect(() => {
        try {
            fetch(API_HOST + "/" + account + "/" + video)
                .then((res) => res.json())
                .then((json) => {
                    setData(json)
                })

        } catch (e) {
            console.log(e)
        }

    }, [])

    return (
        <main className="bg-gray-100 dark:bg-gray-800 h-full min-h-screen relative">
            <div className="flex items-center justify-center">
                <div className="flex flex-col w-full lg:w-3/4 xl:w-3/5 md:space-y-4 mt-4">
                    <div className="pb-24 px-4 md:px-6">
                        <header className="flex md:mt-8 md:mb-8 h-auto">
                            <div className="mr-4">
                                <img src="\/assets/video.jpg" className="rounded-md w-auto max-h-56" />
                            </div>
                            <div className="flex flex-col h-full">
                                <h1 className="text-4xl font-semibold text-gray-800 dark:text-white">
                                    {params.video}
                                </h1>
                                <h2 className="text-md text-gray-400">
                                    Last update: <strong>11.02.22 12:11</strong>
                                </h2>
                                <p className="h-full min-h-full">
                                    {clean(data?.description)}
                                </p>
                                <p className="justify-items-end">
                                    Audio
                                </p>
                            </div>

                        </header>
                        <div className="grid grid-cols-1 w-full  xl:grid-cols-2 gap-4 my-4">
                            <Card big={data && data.views} small={"Views"}>
                                <LineChart data={[{ id: "Views", data: [{ x: 0, y: 0 }, { x: 1, y: 200 }, { x: 2, y: 700 }, { x: 3, y: 1600 }, { x: 4, y: 3000 }, { x: 5, y: 10000 }, { x: 6, y: 11000 },] }]} />
                            </Card>

                            <Card big={data && data.likes} small={"Likes"}>
                                <LineChart data={[{ id: "Likes", data: [{ x: 0, y: 0 }, { x: 1, y: 200 }, { x: 2, y: 700 }, { x: 3, y: 1600 }, { x: 4, y: 3000 }, { x: 5, y: null }, { x: 6, y: 11000 },] }]} />
                            </Card>

                            <Card big={data && data.comments} small={"Comments"} >
                                <LineChart data={[
                                    { id: "Comments", data: [{ x: 0, y: 0 }, { x: 1, y: null }, { x: 2, y: null }, { x: 3, y: 1600 }, { x: 4, y: 3000 }, { x: 5, y: 10000 }, { x: 6, y: 11000 }] },
                                    { id: "Predicted", data: [{ x: 6, y: 11000 }, { x: 7, y: 12000 }, { x: 8, y: 13000 }, { x: 9, y: 14000 }] },
                                    { id: "Average", color: "#FFF", data: [{ x: 0, y: 0 }, { x: 1, y: 100 }, { x: 2, y: 300 }, { x: 3, y: 400 }, { x: 4, y: 700 }, { x: 5, y: 1100 }, { x: 6, y: 1800 }] },
                                ]} />
                            </Card>

                            <Card big={data && data.shares} small={"Shares"} >
                                <LineChart data={[{ id: "Shares", data: [{ x: 0, y: 0 }, { x: 1, y: 200 }, { x: 2, y: 700 }, { x: 3, y: 1600 }, { x: 4, y: null }, { x: 5, y: 2600 }, { x: 6, y: 3000 },] }]} />
                            </Card>


                        </div>


                    </div>
                </div>
            </div>
        </main>
    )
}

export default Video