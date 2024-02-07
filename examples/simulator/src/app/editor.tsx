'use client'

import AceEditor, { IAceOptions } from 'react-ace'
import 'ace-builds/src-min-noconflict/mode-javascript'
import 'ace-builds/src-min-noconflict/mode-yaml'
import { useTheme } from 'next-themes'
import { useCallback, useRef } from 'react'

require('ace-builds/src-min-noconflict/theme-textmate')
require('ace-builds/src-min-noconflict/theme-chaos')

export default function Editor({
  mode = 'javascript',
  value = '',
  className = '',
  onRun = () => {},
  onChange = (_v: string) => {},
  options = {} as IAceOptions,
}) {
  const { theme } = useTheme()
  const ref = useRef(() => {
    onRun()
  })
  ref.current = onRun
  return (
    <AceEditor
      className={className + ' ' + (options.readOnly ? 'ace-readonly' : '')}
      width="100%"
      height="100%"
      theme={theme === 'dark' ? 'chaos' : 'textmate'}
      setOptions={{
        useWorker: false,
        showGutter: true,
        showInvisibles: false,
        showPrintMargin: false,
        highlightActiveLine: false,
        wrap: true,
        tabSize: 2,
        ...options,
      }}
      mode={mode}
      value={value}
      commands={[{
        name: 'run',
        bindKey: { win: 'ctrl-enter', mac: 'cmd-enter' },
        exec: () => ref.current(),
      }]}
      onChange={onChange}
    />
  )
}
