"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.init = void 0;
require("./wasm_exec_polyfill");
const pako_1 = __importDefault(require("pako"));
const node_path_1 = require("node:path");
const fs_1 = require("fs");
function init() {
    return __awaiter(this, void 0, void 0, function* () {
        //@ts-ignore
        const go = new Go();
        const gz = (0, fs_1.readFileSync)((0, node_path_1.resolve)(__dirname, '../wasm/multiparty.wasm.gz'));
        let bytes = pako_1.default.ungzip(gz);
        // HACK: sometimes the gzip is double-gzipped
        if (bytes[0] === 0x1f && bytes[1] === 0x8b) {
            bytes = pako_1.default.ungzip(bytes);
        }
        const module = new WebAssembly.Module(bytes);
        const instance = new WebAssembly.Instance(module, go.importObject);
        const promise = go.run(instance);
        const core = { go, instance, module, promise };
        const _engine = globalThis['MP'];
        return {
            core,
            eval: (code) => __awaiter(this, void 0, void 0, function* () { return _engine(code); }),
        };
    });
}
exports.init = init;
