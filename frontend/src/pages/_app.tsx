import type { AppProps } from "next/app"
import 'modern-css-reset/dist/reset.min.css'

export default function App({ Component, pageProps }: AppProps) {
	return <Component {...pageProps} />
}
