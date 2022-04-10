import React from 'react'

const AreaLayerTest = (props) => {
    console.log("AREA LAYER", props)
    return (
        <div>{JSON.stringify(props)}</div>
    )
}

export default AreaLayerTest