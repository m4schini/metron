import React from 'react'


const App = () => {
  return (

    <div>app</div>
  )
}

export const clean = (str) => {
  if (str) {
    return str.replace(/\.[-a-zA-Z0-9]+{[-a-z:;%0-9(0,. )]+}/g, '')
  } else {
    return ""
  }
}

export const API_HOST = "http://localhost:8080"
export default App