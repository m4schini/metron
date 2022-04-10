import React, { useState } from 'react'

const Start = () => {
  const [query, setQuery] = useState('')

  return (
    <main className="flex justify-center items-center h-screen">
      <div className="flex flex-col items-center">
        <h1 className="text-4xl md:text-8xl font-bold">Metronom</h1>
        <div className="mt-8 w-full">
          <form class="flex flex-col md:flex-row md:w-full max-w-4xl md:space-x-3 space-y-3 md:space-y-0 justify-center"
            onSubmit={(e) => {
              e.preventDefault()
              console.log("submit", query)
              document.location.href = "/" + query

            }}>
            <div class=" relative w-full">
              <input type="text" id="searchAccountFIeld" class=" rounded-lg border-transparent flex-1 appearance-none border border-gray-300 w-full py-2 px-4 bg-white text-gray-700 placeholder-gray-400 shadow-sm text-base focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent"
                placeholder="Account"
                onChange={(e) => {
                  setQuery(e.target.value)
                }}
                value={query} />
            </div>
            <button class="flex-shrink-0 px-4 py-2 text-base font-semibold text-white bg-blue-600 rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:ring-offset-blue-200" type="submit"
            >
              Search
            </button>
          </form>
        </div>
      </div>
    </main >
  )
}

export default Start