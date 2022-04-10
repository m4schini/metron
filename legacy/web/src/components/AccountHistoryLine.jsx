import React, { useState, useEffect } from 'react'
import { ResponsiveLine } from '@nivo/line'

const exampleData = [
    {
        id: "following",
        data: [{ x: 0, y: 1 }, { x: 1, y: 2 }, { x: 2, y: 2 }, { x: 3, y: 4 }, { x: 4, y: 3 }, { x: 5, y: 6 }, { x: 6, y: 7 }, { x: 7, y: 6 }, { x: 8, y: 30 }, { x: 9, y: 33 }]
    }]

const CardLine = ({ data, lineColor }) => (
    <ResponsiveLine
        data={data}
        margin={{ top: 16, right: 0, bottom: 16, left: 0 }}
        xScale={{ type: 'point' }}
        yScale={{
            type: 'linear',
            min: 'auto',
            max: 'auto',
            stacked: true,
            reverse: false
        }}
        yFormat=" >-.2f"
        axisTop={null}
        axisRight={null}
        axisBottom={null}
        axisLeft={null}
        enableGridX={false}
        enableGridY={false}
        colors={lineColor || "#000"}
        lineWidth={2}
        enablePoints={false}
        pointSize={10}
        pointColor={{ theme: 'background' }}
        pointBorderWidth={2}
        pointBorderColor={{ from: 'serieColor' }}
        pointLabelYOffset={-12}
        enableCrosshair={false}
        useMesh={true}
        legends={[]}
    />
)

const AccountFollowingLine = ({ data }) => {

    return data ? <CardLine data={[data]} /> : <CardLine data={exampleData} />
}

export default AccountFollowingLine