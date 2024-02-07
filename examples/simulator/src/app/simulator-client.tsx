'use client'

import '../lib/gowasm'

import { Button } from '@/components/ui/button'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { CheckCircledIcon, CrossCircledIcon } from '@radix-ui/react-icons'
import pako from 'pako'
import { Fragment, useCallback, useEffect, useMemo, useState } from 'react'
import { parse, stringify } from 'yaml'
import Editor from './editor'

export function SimulatorClient({
  examplePolicies = {},
}: {
  examplePolicies: Record<string, string>
}) {
  const [state, setState] = useState({
    runTime: '',
    loadTime: '',
    running: false,
    wasmLoading: false,
    value: examplePolicies.basic,
    output: '',
    expanded: {} as Record<string, boolean>,
    results: [] as any[],
    tests: [] as { test: string; result: boolean }[],
  })
  const onRun = useCallback(() => {
    if (state.running || state.wasmLoading) {
      return
    }
    setState((s) => ({ ...s, running: true }))

    // HACK: show running indicator for at least 200ms
    setTimeout(() => {
      setState((s) => {
        const startTime = performance.now()
        // @ts-ignore
        const result = window.mp(s.value)
        console.log('value', s.value)
        console.log('result', result)
        if (!result) {
          return { ...s, running: false }
        }
        const { output, error: _isErr } = result
        const endTime = performance.now()
        const runTime = `${Math.round((endTime - startTime) * 100) / 100}ms`
        const parsedInput = parse(s.value)
        const parsedOutput = parse(output)
        const tests = parsedInput.tests.map((test: any, i: number) => ({
          test: test.name,
          result:
            !parsedOutput.results[i]?.failures ||
            parsedOutput.results[i]?.failures.length === 0,
        }))
        const results =
          parsedOutput.results?.map((r: any, i: number) => ({
            test: parsedInput?.tests?.[i]?.name,
            in: parsedInput?.tests?.[i]?.in,
            out: r.events[0],
            ...r,
          })) || []
        const newState = {
          runTime,
          output: stringify({ results }),
          tests,
          results,
          expanded: {},
          running: false,
        }
        return { ...s, ...newState }
      })
    }, 200)
  }, [setState, state.value, state.wasmLoading])
  const onSelectExample = useCallback(
    (value: string) => {
      setState((s) => ({
        ...s,
        value: examplePolicies[value].trim(),
        output: '',
        runTime: '',
      }))

      setTimeout(() => {
        onRun()
      }, 100)
    },
    [state.value, onRun]
  )
  useEffect(() => {
    setState((s) => ({ ...s, wasmLoading: true }))
    async function load() {
      const startTime = performance.now()
      // @ts-ignore
      const go = new Go()
      const res = await fetch('/static/multiparty.wasm.gz')
      const gzBuf = await res.arrayBuffer()
      let buffer = pako.ungzip(gzBuf)
      // HACK: sometimes the gzip is double-gzipped
      if (buffer[0] === 0x1f && buffer[1] === 0x8b) {
        buffer = pako.ungzip(buffer)
      }
      const endTime = performance.now()
      const loadTime = `${Math.round((endTime - startTime) * 100) / 100}ms`
      setState((s) => ({ ...s, loadTime, wasmLoading: false }))
      setTimeout(() => onRun(), 0)
      await WebAssembly.instantiate(buffer, go.importObject).then((result) =>
        go.run(result.instance)
      )
    }
    // @ts-ignore
    if (!window.mp) {
      // @ts-ignore
      window.mp = () => {}
      load()
    } else {
      setState((s) => ({ ...s, wasmLoading: false }))
    }
  }, [setState])
  return (
    <div className="h-full flex flex-col space-y-8 pt-8 sm:pt-0 sm:space-y-0 sm:grid grid-cols-2 gap-4">
      <div className="min-h-screen sm:min-h-0 relative">
        <div className="absolute w-full flex items-center justify-between -top-8 z-10">
          <span />
          <span className="text-xs text-gray-600">
            <Select value="" onValueChange={onSelectExample}>
              <SelectTrigger className="w-[180px] h-6 text-xs px-1">
                <SelectValue placeholder="Select an example" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem className="text-xs" value="basic">
                  Basic
                </SelectItem>
                <SelectItem className="text-xs" value="github">
                  GitHub
                </SelectItem>
                <SelectItem className="text-xs" value="finance">
                  Finance
                </SelectItem>
              </SelectContent>
            </Select>
          </span>
        </div>
        <div className="h-full rounded-md overflow-hidden border">
          <div className="absolute top-2 right-2 z-10">
            <Button
              onClick={onRun}
              disabled={state.running}
              className="light:bg-white h-8 px-2 leading-5"
              variant="outline"
            >
              {state.running ? 'Running' : 'Run'}
            </Button>
          </div>
          <Editor
            mode="yaml"
            onRun={onRun}
            value={state.value}
            onChange={(value) =>
              setState((s) => ({ ...s, value, running: false }))
            }
          />
        </div>
      </div>
      <div className="min-h-screen sm:min-h-0 relative">
        <div className="absolute w-full flex items-center justify-between -top-8 z-10">
          <span className="font-semibold">Output</span>
          <span className="text-xs text-gray-600">
            {state.wasmLoading ? (
              <>Loading...</>
            ) : state.runTime ? (
              <>
                {state.tests.filter((t) => t.result).length}/
                {state.tests.length} passed &middot; Ran tests in{' '}
                {state.runTime}
              </>
            ) : state.loadTime ? (
              <>Ready in {state.loadTime}</>
            ) : null}
          </span>
        </div>
        <div className="h-full rounded-md overflow-scroll border bg-[#f0f0f0] dark:bg-[#0b0b0b] flex flex-col">
          <div
            className={
              'overflow-scroll border-b ' +
              (state.output ? 'opacity-100' : 'opacity-0')
            }
          >
            <Table wrapperClassName="relative h-60" className="relative">
              <TableHeader className="h-8 bg-muted border-b">
                <TableRow>
                  <TableHead className="flex-1">Test</TableHead>
                  <TableHead className="text-right">Result</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {state.tests.map((test, i) => (
                  <Fragment key={test.test + '-' + i}>
                    <TableRow
                      className="cursor-pointer bg-white hover:bg-gray-50 dark:bg-[#0b0b0b] dark:hover:bg-[#1b1b1b]"
                      onClick={() => {
                        setState((s) => ({
                          ...s,
                          expanded: {
                            ...s.expanded,
                            [test.test + '-' + i]:
                              !s.expanded[test.test + '-' + i],
                          },
                        }))
                      }}
                    >
                      <TableCell className="font-medium">{test.test}</TableCell>
                      <TableCell className="flex justify-end">
                        {test.result ? (
                          <CheckCircledIcon
                            className="text-green-600"
                            width={20}
                            height={20}
                          />
                        ) : (
                          <CrossCircledIcon
                            className="text-red-600"
                            width={20}
                            height={20}
                          />
                        )}
                      </TableCell>
                    </TableRow>
                    {state.expanded[test.test + '-' + i] && (
                      <TableRow className="bg-gray-50 hover:bg-gray-50 dark:bg-[#1b1b1b] h-24">
                        <td>
                          <pre className="p-2.5">
                            {stringify(state.results[i])}
                          </pre>
                        </td>
                        <td />
                      </TableRow>
                    )}
                  </Fragment>
                ))}
              </TableBody>
            </Table>
          </div>
          <div
            className={'flex-1 ' + (state.output ? 'opacity-100' : 'opacity-0')}
          >
            <Editor
              mode="yaml"
              value={state.running ? 'Running...' : state.output.trim()}
              options={{ readOnly: true }}
            />
          </div>
        </div>
      </div>
    </div>
  )
}
