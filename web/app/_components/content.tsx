"use client"

import Pad from "./pad";
import { useState } from "react";

export default function Content() {

	const [query, setQuery] = useState(`db.users.find({
    $or: [
        { age: 
            { $lt: 18 } 
        }, 
        { age: 
            { $gt: 65 } }
    ]
})`)
	const [code, setCode] = useState(`bson.M{
    "$or": bson.A{
        bson.M{
            "age": bson.M{
                "$lt": 18
            }
        },
        bson.M{
            "age": bson.M{
                "$gt": 65
            }
        }
    }
}`)

	return (
		<div className="h-full flex justify-evenly items-center">
			<Pad
				background="bg-green-100"
				title="MongoDB Query"
				readOnly={false}
				value={query}
				setValue={setQuery}
			/>
			<Pad
				background="bg-blue-100"
				title="Golang Driver Code"
				readOnly={true}
				value={code}
				setValue={setCode}
			/>
		</div>
	)
}