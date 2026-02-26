import Image from "next/image"
import Logo from "../../public/logo-web.png"

export default function Header() {
	return (
		<Image
			src={Logo}
			alt="logo.png"
			width={0}
			height={0}
			sizes="100vw"
			style={{ width: 'auto', height: '100%' }}
		/>
	)
}