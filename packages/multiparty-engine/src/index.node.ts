import './wasm_exec_polyfill'
import pako from 'pako'
import { resolve } from 'node:path'
import { readFileSync } from 'fs'

type Engine = {
  core: any
  eval: (code: string) => Promise<any>
}

export async function init(): Promise<Engine> {
  //@ts-ignore
  const go = new Go()
  const gz = readFileSync(resolve(__dirname, '../wasm/multiparty.wasm.gz'))
  let bytes = pako.ungzip(gz)
  // HACK: sometimes the gzip is double-gzipped
  if (bytes[0] === 0x1f && bytes[1] === 0x8b) {
    bytes = pako.ungzip(bytes)
  }
  const module = new WebAssembly.Module(bytes)
  const instance = new WebAssembly.Instance(module, go.importObject)
  const promise = go.run(instance)
  const core = { go, instance, module, promise }
  const _engine: any = (globalThis as any)['MP']
  return {
    core,
    eval: async (code: string) => _engine(code),
  }
}
