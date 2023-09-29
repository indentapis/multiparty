// @ts-ignore
import { ExecEngine, SplitPath } from './engine';
import { loadPyodide } from 'pyodide';
import {
    SetupRequest, SetupResponse,
    Credential, CredentialResponse,
    Action, Response,
    CloseRequest, CloseResponse,
    OpType, opTypeToJSON,
    Query, QueryResponse
} from './protobuf-ts/indent/exec/v1/exec';

// the executor script defines the class interface for query and update scripts
const executorScript = `
class Executor:
    supported_types = {}
    query_schema = {}
    response_schema = {}

    def run_query(self, apply_request):
        raise NotImplementedError

    def apply_update(self, apply_request):
        raise NotImplementedError

    def pull_update(self, pull_request):
        raise NotImplementedError
`

// the setup script sets up protobuf and requests libraries so that exec scripts Just Work
const setupScript = `
    import micropip
    await micropip.install(["pyodide-http==0.2.1", "google", "google-api-python-client", "protobuf"])
    from google.protobuf.json_format import Parse, MessageToJson

    # patch requests library so 3rd party libraries can work
    await micropip.install(["requests"])
    import pyodide_http
    pyodide_http.patch_all()
`;

const pythonTypes = [
    'Parse',
    'MessageToJson',
    'Credential',
    'JsonPatch',
    'Event',
    'SetupRequest',
    'SetupResponse',
    'Query',
    'QueryResponse'
]

const typescriptTypes = {
    'Credential': Credential,
    'CredentialResponse': CredentialResponse,
    'Query': Query,
    'QueryResponse': QueryResponse,
    // TODO: can this be generated?
    // SetupRequest, SetupResponse,
    // CredentialResponse,
    // Action, Response,
    // CloseRequest, CloseResponse,
    // OpType, opTypeToJSON,
    // Query, QueryResponse
}

// Pyodide returns results of bare statements; this statement contains all class definitions
// that are needed to pass Python types directly in
const fetchPythonTypes = "{" + pythonTypes.reduce((acc, type) => `${acc}'${type}': ${type}, `, '') + "}";

export class PythonExecEngine implements ExecEngine {
    // TODO: how does Pyodide handle async execution? does this need to handle blocking?
    #pythonProtobuf; #executorClass; #executor; #supportedTypes; #pyodide

    private toPython(typeName: string, typeInstance: any) {
        let pyClass = this.#pythonProtobuf.get(typeName);
        let Parse = this.#pythonProtobuf.get("Parse");

        let tsType = typescriptTypes[typeName];
        let typeJSON = JSON.stringify(tsType.toJSON(typeInstance));
        let message = Parse(typeJSON, pyClass());
        return message;
    }

    private async fetchScript(scriptName) {
        const response = await fetch('/' + scriptName);
        return await response.text();
    }
 
    private async setupRuntime(exec) {
        this.#pyodide = await loadPyodide({
            indexURL: "pyodide-0.23.4/pyodide"
        });
        await this.#pyodide.loadPackage(['micropip', 'ssl']);
        await this.#pyodide.runPythonAsync(executorScript);
        await this.#pyodide.runPythonAsync(setupScript);

        await this.#pyodide.runPythonAsync(await this.fetchScript("exec_pb2.py"));
        let pythonProtobuf = (await this.#pyodide.runPythonAsync(fetchPythonTypes)).toJs();

        let script = exec.imageBody;
        if (exec.imagePath) {
            script = await this.fetchScript(exec.imagePath);
        }
        
        let results = await this.#pyodide.runPythonAsync(script);
        let pythonAPI = results.toJs();

        // the Python API returns an executor of type Executor
        return [pythonProtobuf, pythonAPI.get('types'), pythonAPI.get('executor')];
    }

    async Init(initRequest: SetupRequest) {
        [
            this.#pythonProtobuf,
            this.#supportedTypes,
            this.#executorClass
        ] = await this.setupRuntime(initRequest.exec);
        return SetupResponse.fromJSON({});
    }


    async InstantiateSecret(credentialRequest: Credential | null) {
        if (credentialRequest != null) {
            let pyCredential = this.toPython('Credential', credentialRequest);
            this.#executor = this.#executorClass(pyCredential);
        } else {
            this.#executor = this.#executorClass();
        }
        return CredentialResponse.fromJSON({});
    }

    async RunQuery(query: Query) {
        let pyQuery = this.toPython('Query', query);
        let response = (await this.#executor.run_query(pyQuery)).toJs();
        let MessageToJson = this.#pythonProtobuf.get("MessageToJson");
        return QueryResponse.fromJSON(JSON.parse(MessageToJson(response)));
    }

    async RunOperation(action: Action) {
        let [handler, kind, targetId] = SplitPath(action.patch);
        if (!this.#supportedTypes.get(handler) || !this.#supportedTypes.get(handler).get(kind)) {
            return Response.fromJSON({
                actionId: action.id,
                status: {code: -1, message: "unsupported handler or kind"}
            })
        }

        if (!this.#supportedTypes.get(handler).get(kind).includes(opTypeToJSON(action.patch.op))) {
            return Response.fromJSON({
                actionId: action.id,
                status: {code: -1, message: "unsupported operation type for this handler"}
            })
        }

        switch(action.patch.op) {
            case OpType.copy: {
                // TODO: need to translate this into a `Response`
                let response = (await this.#executor.pull_update(kind)).toJs();
                let MessageToJson = this.#pythonProtobuf.get("MessageToJson");
                return Response.fromJSON(JSON.parse(MessageToJson(response)));
            }

            case OpType.add: {
                // TODO: once this is solid, add it to types and pass in as Python object
                const event = {
                    "target": targetId,
                    "op": "add",
                    "subject": action.patch.value.name
                }
                // TODO: need to translate this into a `Response`
                let results = (await this.#executor.apply_update(event)).toJs();
                return Response.fromJSON({});
            }

            case OpType.remove: {
                const event = {
                    "target": targetId,
                    "op": "remove"
                }
                // TODO: need to translate this into a `Response`
                let results = (await this.#executor.apply_update(event)).toJs();
                return Response.fromJSON({});
            }
        }
        return Response.fromJSON({
            actionId: action.id,
            status: {code: -1, message: "optype does not exist"}
        })
    }

    async Close(closeRequest: CloseRequest) {
        return CloseResponse.fromJSON({});
    }
}
