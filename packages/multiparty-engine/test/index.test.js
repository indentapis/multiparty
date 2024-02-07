const engine = require('../out/index.node')
const fs = require('node:fs/promises')
const { resolve } = require('node:path')

import { expect, test } from 'bun:test'

test('mp', async () => {
  const result = await engine.init()
  if (!result) {
    throw new Error('Missing: @indent/engine')
  }

  const buf = await fs.readFile(
    resolve(__dirname, '../../../pkg/engine/testdata/basic.yaml')
  )
  const { output } = await result.eval(buf.toString())

  expect(output).toBeDefined()
})
