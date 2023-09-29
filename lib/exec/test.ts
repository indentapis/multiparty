// these need to be moved into a proper test framework.
// for the time being: unzip pyodide-0.23.4 into `/public/`
// and serve it such that pyodide.js is available at
// `/public/pyodide-0.23.4/pyodide/pyodide.js`.

import { PythonExecEngine } from "./python";
import {
    SetupRequest, SetupResponse,
    Credential, CredentialResponse,
    Action, Response,
    CloseRequest, CloseResponse,
    JsonPatch,
    Query
} from './protobuf-ts/indent/exec/v1/exec';

const mailExec = new PythonExecEngine();

await mailExec.Init(SetupRequest.fromJSON({exec: {imagePath: "mailcheck.py"}}));
await mailExec.InstantiateSecret(null);

// mailcheck.ai
// this should reply with a disposable = true, alias = false
let queryResponse = await mailExec.RunQuery(Query.fromJSON({
    input: '{"user.email": "fakeemail@mailinator.com"}'
}));
console.log(queryResponse);

// mailcheck.ai
// this should reply with a disposable = false, alias = false
queryResponse = await mailExec.RunQuery(Query.fromJSON({
    input: '{"user.email": "steve@apple.com"}'
}));
console.log(queryResponse);

// mailcheck.ai
// this should reply with a disposable = false, alias = true
queryResponse = await mailExec.RunQuery(Query.fromJSON({
    input: '{"user.email": "page+alias@gmail.com"}'
}));
console.log(queryResponse);

/*

const stripe_api_secret = "s####"

const stripeExec = new PythonExecEngine();

await stripeExec.Init(SetupRequest.fromJSON({exec: {imagePath: "stripe_api.py"}}));
await stripeExec.InstantiateSecret(Credential.fromJSON({apiKey: stripe_api_secret}));


// this should succeed, and place the Customers in the $out results as an array
let response = await stripeExec.RunOperation(Action.fromJSON({
    id: "abc123",
    patch: {
        "op": "copy",
        "from": "/stripe/customers",
        "path": "$out" // equivalent to empty string, which is JsonPatch for "root of document"
    },
}));
console.log(response);

// this should error, as this instantiation does not support Okta
response = await stripeExec.RunOperation(Action.fromJSON({
    id: "xyz987",
    patch: {
        "op": "copy",
        "from": "/okta/groups",
        "path": "$out"
    },
}));
console.log(response);
*/