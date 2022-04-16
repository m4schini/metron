import React from 'react'
import aveta from 'aveta';


const RedGreenText = ({ value, postfix }) => {
    let css = ""
    if (value >= 0) {
        css = "flex items-center text-sm gap-1 text-green-500"
    } else {
        css = "flex items-center text-sm gap-1 text-red-500"
    }

    return <div className={css}>
        <span>
            {value}%
        </span>
        <span className="text-gray-400">
            {postfix}
        </span>
    </div>
}

const MiniCardSkeleton = ({ icon, title }) => {
    return <div className="shadow-lg rounded-2xl p-4 bg-white dark:bg-gray-800">
        <div className="flex items-center">
            <span className={`rounded-xl relative`}>
                {icon}
            </span>
            <p className="text-md text-black dark:text-white ml-2">
                {title}
            </p>
        </div>
        <div className="flex justify-between">
            <div className="h-12 w-3/4 mt-1 rounded-xl bg-gray-200 animate-pulse"></div>
            <div>
                <div className="bg-gray-200 w-16 animate-pulse h-3 mb-1 rounded-2xl"></div>
                <div className="bg-gray-200 w-16 animate-pulse h-3 mb-1 rounded-2xl"></div>
                <div className="bg-gray-200 w-16 animate-pulse h-3 mb-1 rounded-2xl"></div>

            </div>
        </div>
    </div>
}

const MiniCard = ({ title, content, unit, icon }) => {
    if (content == null) {
        return <MiniCardSkeleton icon={icon} title={title} />
    }

    return (

        <div className="shadow-lg rounded-2xl p-4 bg-white dark:bg-gray-800">
            <div className="flex items-center">
                <span className={`rounded-xl relative`}>
                    {icon}
                </span>
                <p className="text-md text-black dark:text-white ml-2">
                    {title}
                </p>
            </div>
            <div className="flex justify-between">
                <p className="text-gray-700 dark:text-gray-100 text-4xl text-left font-bold my-4">
                    {aveta(content?.value)}
                    <span className="text-sm">
                        {unit}
                    </span>
                </p>
                <div>
                    <RedGreenText value={content?.fiveMin} postfix={"5min"} />
                    <RedGreenText value={content?.oneDay} postfix={"24h"} />
                    <RedGreenText value={content?.oneMonth} postfix={"30d"} />

                </div>
            </div>
        </div>

    )
}

export const AccountViews = () => {

}

export default MiniCard