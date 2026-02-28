import Editor from "@monaco-editor/react";

type padProps = {
	background: string,
	title: string,
	readOnly: boolean,
	value: string,
	setValue: (value: string) => void
}

export default function Pad({
	background,
	title,
	readOnly,
	value,
	setValue
}: padProps) {

	return (
		<div className={`h-[85vh] w-[45vw] ${background}`}>
			<div className="flex items-center justify-center h-[5vh]">
				{title}
			</div>
			<div className="h-[80vh] p-4">
				<Editor
					defaultLanguage="python"
					value={value}
					onChange={(value, _) => setValue(value ?? "")}
					theme="light"
					options={{
						fontSize: 14,
						readOnly: readOnly,
						minimap: { enabled: false },
						lineNumbers: "on",
						wordWrap: "on",
					}}
				/>
			</div>
		</div>
	)
}