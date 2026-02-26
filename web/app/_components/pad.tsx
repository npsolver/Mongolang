export default function Pad({ background, title }: { background: string, title: string }) {
	return (
		<div className={`h-[85vh] w-[45vw] ${background}`}>
			<div className="flex items-center justify-center h-[5vh]">
				{title}
			</div>
			<div className="h-[80vh] p-4">
				Bottom
			</div>
		</div>
	)
}