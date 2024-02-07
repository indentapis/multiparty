import crypto from 'node:crypto'
import fs from 'node:fs'
globalThis.TextEncoder = require('util').TextEncoder
globalThis.TextDecoder = require('util').TextDecoder

//@ts-ignore
if (!globalThis.fs) {
  Object.defineProperty(globalThis, 'fs', {
    value: fs,
  })
}

if (!globalThis.process) {
  Object.defineProperty(globalThis, 'process', {
    value: process,
  })
}

if (!globalThis.crypto) {
  Object.defineProperty(globalThis, 'crypto', {
    value: crypto.webcrypto
      ? crypto.webcrypto
      : {
          getRandomValues(b: any) {
            return crypto.randomFillSync(b)
          },
        },
  })
}

if (!globalThis.performance) {
  Object.defineProperty(globalThis, 'performance', {
    value: {
      now() {
        const [sec, nsec] = process.hrtime()
        return sec * 1000 + nsec / 1000000
      },
    },
  })
}

import '../go/wasm_exec'