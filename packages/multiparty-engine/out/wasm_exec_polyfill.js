"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_crypto_1 = __importDefault(require("node:crypto"));
const node_fs_1 = __importDefault(require("node:fs"));
globalThis.TextEncoder = require('util').TextEncoder;
globalThis.TextDecoder = require('util').TextDecoder;
//@ts-ignore
if (!globalThis.fs) {
    Object.defineProperty(globalThis, 'fs', {
        value: node_fs_1.default,
    });
}
if (!globalThis.process) {
    Object.defineProperty(globalThis, 'process', {
        value: process,
    });
}
if (!globalThis.crypto) {
    Object.defineProperty(globalThis, 'crypto', {
        value: node_crypto_1.default.webcrypto
            ? node_crypto_1.default.webcrypto
            : {
                getRandomValues(b) {
                    return node_crypto_1.default.randomFillSync(b);
                },
            },
    });
}
if (!globalThis.performance) {
    Object.defineProperty(globalThis, 'performance', {
        value: {
            now() {
                const [sec, nsec] = process.hrtime();
                return sec * 1000 + nsec / 1000000;
            },
        },
    });
}
require("../go/wasm_exec");
