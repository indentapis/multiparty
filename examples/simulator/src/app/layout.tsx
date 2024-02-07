import type { Metadata } from 'next'
import './globals.css'

import { ThemeProvider } from './theme-provider'

export const metadata: Metadata = {
  title: 'Multiparty Simulator',
  description: 'Basic simulator for writing and testing multiparty rules.',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="en">
      <body>
        <ThemeProvider
          attribute="class"
          defaultTheme="light"
          disableTransitionOnChange
        >
          {children}
        </ThemeProvider>
      </body>
    </html>
  )
}
