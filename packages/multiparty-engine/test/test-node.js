const engine = require('..')
const fs = require('node:fs/promises')
const { resolve } = require('node:path')

async function run() {
  const result = await engine.init()
  if (!result) {
    throw new Error('Missing: @indent/engine')
  }

  const buf = await fs.readFile(
    resolve(__dirname, '../../../pkg/engine/testdata/basic.yaml')
  )
  const {output } = await result.eval(buf.toString())

  console.log(output)
  return
}

run()
  .then(() => {})
  .catch((err) => {
    if (err) {
      console.error(err)
      process.kill(process.pid, 'SIGTERM')
    }
  })
