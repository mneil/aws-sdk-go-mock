import {AwsSdkGoMock} from "./wasm_exec.js"

const mock = AwsSdkGoMock();
const result = await WebAssembly.instantiateStreaming(fetch(`main.wasm`), mock.importObject)
mock.run(result.instance)
