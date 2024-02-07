import engine from '..'
import fs from 'node:fs/promises'
import { resolve } from 'node:path'
import { expect, test } from 'bun:test'

test('[ESM] multiparty: init() + eval()', async () => {
  const result = await engine.init()
  expect(result).toBeDefined()

  const buf = await fs.readFile(
    resolve(__dirname, '../../../pkg/engine/testdata/basic.yaml')
  )
  const { output } = await result.eval(buf.toString())

  expect(output).toBeDefined()
})
