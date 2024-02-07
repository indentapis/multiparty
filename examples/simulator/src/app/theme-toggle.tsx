'use client'

import { Button } from '@/components/ui/button'
import { MoonIcon, SunIcon } from '@radix-ui/react-icons'
import { useTheme } from 'next-themes'

export function ThemeToggle() {
  const { theme, setTheme } = useTheme()

  return (
    <Button
      onClick={() => (theme === 'dark' ? setTheme('light') : setTheme('dark'))}
      className="w-5 h-5 transform translate-y-[-0.75px]"
      variant="outline"
      size="icon"
    >
      <SunIcon className="h-[0.5rem] w-[0.5rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
      <MoonIcon className="absolute h-[0.65rem] w-[0.65rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
      <span className="sr-only">Toggle theme</span>
    </Button>
  )
}
