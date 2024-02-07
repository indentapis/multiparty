import '../go/wasm_exec'
import pako from 'pako'

type Engine = {
  core: any
  eval: (code: string) => Promise<any>
}

const CDN_HOST = 'https://cdn.jsdelivr.net/npm/@indent/engine@latest'

export async function init(): Promise<Engine> {
  //@ts-ignore
  const go = new Go()
  const gz = await fetch(`${CDN_HOST}/wasm/multiparty.wasm.gz`).then((r) =>
    r.arrayBuffer()
  )
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
