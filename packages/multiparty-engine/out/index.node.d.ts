import './wasm_exec_polyfill';
type Engine = {
    core: any;
    eval: (code: string) => Promise<any>;
};
export declare function init(): Promise<Engine>;
export {};
