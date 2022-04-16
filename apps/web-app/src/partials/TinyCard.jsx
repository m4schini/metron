import React from 'react'

const TinyCard = ({ title, content, icon }) => {
    return (
        <div className="border-4 rounded-2xl w-36 p-4 bg-white dark:bg-gray-800">
            <div className="flex items-center">
                <div className="h-4">
                    {icon}
                </div>
                <p className="text-md text-gray-700 dark:text-gray-50 ml-2">
                    {title}
                </p>
            </div>
            <div className="flex flex-col justify-start">
                <p className="text-gray-800 text-4xl text-left dark:text-white font-bold my-1">
                    {content}
                </p>
            </div>
        </div>
    )
}

export default TinyCard