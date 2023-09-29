import {
    SetupRequest, SetupResponse,
    Credential, CredentialResponse,
    Action, Response,
    CloseRequest, CloseResponse,
    JsonPatch, Query, QueryResponse
} from './protobuf-ts/indent/exec/v1/exec';

// ExecEngine is an implementation of a backend for executing an Action
export interface ExecEngine {
    // Init handles setting up the exec environmen  t (ex: instantiating a Python runtime)
    Init(initRequest: SetupRequest): Promise<SetupResponse>;
    // InstantiateSecret loads any secrets required for the given execution
    InstantiateSecret(secretRequest: Credential | null): Promise<CredentialResponse>;
    // RunOperation runs an Operation
    RunOperation(action: Action): Promise<Response>;
    // RunQuery runs a Query
    RunQuery(query: Query): Promise<QueryResponse>;
    // Close destroys the environment
    Close(closeRequest: CloseRequest): Promise<CloseResponse>;
}

// SplitPath returns the handler, kind, and targetId from a JsonPatch
// /{handler}/{kind}/{targetId?}
export function SplitPath(patch: JsonPatch | undefined): [string, string, string] {
    if (patch === undefined) {
        return ["", "", ""]
    }
    // if `from` exists, this is a `copy` request, and `.path` stores the $out
    const opPath = patch.from ? patch.from : patch.path; 
    const path = opPath.split("/");
    const handler = path[1];
    const kind = path[2];
    const targetId = path[3];
    return [handler, kind, targetId]
}
