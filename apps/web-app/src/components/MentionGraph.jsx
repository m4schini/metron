import React from 'react'
import { ResponsiveNetwork } from '@nivo/network'
import randomColor from 'randomcolor'

const colors = ["#fde725", "#7ad151", "#22a884", "#2a788e", "#414487", "#440154"]

const exampleData = {
    nodes: [
        { id: "Krawallklara", size: 5, color: "#000", height: 1 },
        { id: "Acc2", size: 10 * 2, color: colors[0], height: 1 },
        { id: "Acc3", size: 10 * 2, color: colors[1], height: 1 },
        { id: "Acc4", size: 10 * 2, color: colors[2], height: 1 },
        { id: "Acc5", size: 20 * 2, color: colors[3], height: 2 },
        { id: "Acc6", size: 20 * 2, color: colors[4], height: 2 },
        { id: "Acc7", size: 30 * 2, color: colors[5], height: 3 },
    ],
    links: [
        { source: "Krawallklara", target: "Acc2", distance: 30 * 2 },
        { source: "Krawallklara", target: "Acc3", distance: 30 * 2 },
        { source: "Krawallklara", target: "Acc4", distance: 30 * 2 },
        { source: "Krawallklara", target: "Acc5", distance: 20 * 2 },
        { source: "Krawallklara", target: "Acc6", distance: 20 * 2 },
        { source: "Krawallklara", target: "Acc7", distance: 10 * 2 },
    ]
}

const NetworkChart = ({ data }) => (
    <ResponsiveNetwork
        data={data}
        margin={{ top: 0, right: 0, bottom: 0, left: 0 }}
        linkDistance={function (e) { return e.distance }}
        centeringStrength={0.3}
        repulsivity={120}
        nodeSize={function (n) { return n.size }}
        activeNodeSize={function (n) { return 1.5 * n.size }}
        nodeColor={(e) => e.color}
        nodeBorderWidth={1}
        nodeBorderColor={{
            from: 'color',
            modifiers: [
                [
                    'darker',
                    0.2
                ]
            ]
        }}
        linkColor={{ from: 'target.color', modifiers: [] }}
        linkThickness={function (n) { return 2 + 1.2 * n.target.data.height }}
        linkBlendMode="multiply"
        motionConfig="wobbly"
    />
)

const MentionGraph = ({ data }) => {

    return data ? <NetworkChart data={data} /> : <NetworkChart data={exampleData} />
}

export default MentionGraph