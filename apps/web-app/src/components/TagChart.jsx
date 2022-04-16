import React, { useEffect, useState } from 'react'
import { ResponsivePieCanvas } from '@nivo/pie'
import { API_HOST } from '../App'

const exampleData = [{ id: "tag", value: 2 }, { id: "tag2", value: 4 }, { id: "ta3", value: 6 }, { id: "tag33", value: 5 }, { id: "tag4", value: 13 }, { id: "1312", value: 10 },]
const colors = ["#7ad151", "#22a884", "#2a788e", "#414487", "#440154"]

const PieChart = ({ data /* see data tab */ }) => (
    <ResponsivePieCanvas
        data={data}
        margin={{ top: 20, right: 40, bottom: 20, left: 40 }}
        endAngle={-360}
        sortByValue={true}
        innerRadius={0.5}
        padAngle={0.7}
        activeOuterRadiusOffset={8}
        colors={{ scheme: 'category10' }}
        arcLinkLabelsSkipAngle={14}
        arcLinkLabelsTextColor="#333333"
        arcLinkLabelsThickness={2}
        arcLinkLabelsColor={{ from: 'color' }}
        arcLabelsSkipAngle={10}
        arcLabelsTextColor="#FFF"
        tooltip={(v) => <div className="bg-slate-900 text-white">{v}</div>}
        defs={[
            {
                id: 'dots',
                type: 'patternDots',
                background: 'inherit',
                color: 'rgba(255, 255, 255, 0.3)',
                size: 4,
                padding: 1,
                stagger: true
            },
            {
                id: 'lines',
                type: 'patternLines',
                background: 'inherit',
                color: 'rgba(255, 255, 255, 0.3)',
                rotation: -45,
                lineWidth: 6,
                spacing: 10
            }
        ]}
        legends={[]}
    />
)

const TagChart = ({ data }) => {
    return data
        ? <PieChart data={data} />
        : <div>test</div>
}

export default TagChart