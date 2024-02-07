const engine = require('..')
const fs = require('node:fs/promises')
const { resolve } = require('node:path')

import { expect, test } from 'bun:test'

test('[NODE] multiparty: init() + eval()', async () => {
  const result = await engine.init()
  expect(result).toBeDefined()

  const buf = await fs.readFile(
    resolve(__dirname, '../../../pkg/engine/testdata/basic.yaml')
  )
  const { output } = await result.eval(buf.toString())

  expect(output).toBeDefined()
})
