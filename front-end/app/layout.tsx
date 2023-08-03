import './globals.css'

export const metadata = {
  title: 'Travel',
  description: 'Travel assistance',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
