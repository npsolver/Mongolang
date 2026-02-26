import Pad from "./pad";

export default function Content() {
	return (
		<div className="h-full flex justify-evenly items-center">
			<Pad
				background="bg-green-100"
				title="MongoDB Query"
			/>
			<Pad
				background="bg-blue-100"
				title="Golang Driver Code"
			/>
		</div>
	)
}