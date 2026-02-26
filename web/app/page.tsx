import Header from "./_components/header"
import Content from "./_components/content"

export default function Page() {
	return (
		<div>
			<div className="h-[10vh]">
				<Header />
			</div>
			<div className="h-[90vh]">
				<Content />
			</div>
		</div>
	)
}
