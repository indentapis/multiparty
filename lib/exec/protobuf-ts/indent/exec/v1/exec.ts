/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Value } from "../../../google/protobuf/struct";
import { Status } from "../../../google/rpc/status";

export const protobufPackage = "indent.jsonschema.draft07";

export enum OpType {
  add = 0,
  remove = 1,
  replace = 2,
  copy = 3,
  test = 4,
  UNRECOGNIZED = -1,
}

export function opTypeFromJSON(object: any): OpType {
  switch (object) {
    case 0:
    case "add":
      return OpType.add;
    case 1:
    case "remove":
      return OpType.remove;
    case 2:
    case "replace":
      return OpType.replace;
    case 3:
    case "copy":
      return OpType.copy;
    case 4:
    case "test":
      return OpType.test;
    case -1:
    case "UNRECOGNIZED":
    default:
      return OpType.UNRECOGNIZED;
  }
}

export function opTypeToJSON(object: OpType): string {
  switch (object) {
    case OpType.add:
      return "add";
    case OpType.remove:
      return "remove";
    case OpType.replace:
      return "replace";
    case OpType.copy:
      return "copy";
    case OpType.test:
      return "test";
    case OpType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Executable {
  imagePath?: string | undefined;
  imageBody?:
    | string
    | undefined;
  /** TODO: complex semver type? */
  version: string;
}

export interface JsonPatch {
  op: OpType;
  path: string;
  value: any | undefined;
  from: string;
}

export interface Event {
  patchId: string;
  /**
   * oneof EventBody {
   * TODO: other event types that aren't patch directives will defined here
   */
  execMessage: string;
}

/** send setup call */
export interface SetupRequest {
  exec: Executable | undefined;
}

/** executable is ready */
export interface SetupResponse {
  /**
   * repeated Type supported_types = 2;
   * TODO: return reported compatible types
   */
  status: Status | undefined;
}

/** TODO: pass in any needed auth (use Workload identity to resolve a secret_name reference) */
export interface Credential {
  apiKey: string;
}

/** executable has been primed */
export interface CredentialResponse {
}

/** for Decisions */
export interface Query {
  id: string;
  input: string;
}

export interface QueryResponse {
  status: Status | undefined;
  id: string;
  results: string;
}

/**
 * for Actions
 * repeat as necessary
 */
export interface Action {
  id: string;
  patch: JsonPatch | undefined;
}

/** returned as many times as necessary FROM the executable */
export interface Response {
  status: Status | undefined;
  actionId: string;
  results: any[];
}

/** close the executor */
export interface CloseRequest {
}

export interface CloseResponse {
}

function createBaseExecutable(): Executable {
  return { imagePath: undefined, imageBody: undefined, version: "" };
}

export const Executable = {
  encode(message: Executable, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.imagePath !== undefined) {
      writer.uint32(10).string(message.imagePath);
    }
    if (message.imageBody !== undefined) {
      writer.uint32(18).string(message.imageBody);
    }
    if (message.version !== "") {
      writer.uint32(26).string(message.version);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Executable {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExecutable();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.imagePath = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.imageBody = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.version = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Executable {
    return {
      imagePath: isSet(object.imagePath) ? String(object.imagePath) : undefined,
      imageBody: isSet(object.imageBody) ? String(object.imageBody) : undefined,
      version: isSet(object.version) ? String(object.version) : "",
    };
  },

  toJSON(message: Executable): unknown {
    const obj: any = {};
    if (message.imagePath !== undefined) {
      obj.imagePath = message.imagePath;
    }
    if (message.imageBody !== undefined) {
      obj.imageBody = message.imageBody;
    }
    if (message.version !== "") {
      obj.version = message.version;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Executable>, I>>(base?: I): Executable {
    return Executable.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Executable>, I>>(object: I): Executable {
    const message = createBaseExecutable();
    message.imagePath = object.imagePath ?? undefined;
    message.imageBody = object.imageBody ?? undefined;
    message.version = object.version ?? "";
    return message;
  },
};

function createBaseJsonPatch(): JsonPatch {
  return { op: 0, path: "", value: undefined, from: "" };
}

export const JsonPatch = {
  encode(message: JsonPatch, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.op !== 0) {
      writer.uint32(8).int32(message.op);
    }
    if (message.path !== "") {
      writer.uint32(18).string(message.path);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(26).fork()).ldelim();
    }
    if (message.from !== "") {
      writer.uint32(34).string(message.from);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JsonPatch {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJsonPatch();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.op = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.path = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.from = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): JsonPatch {
    return {
      op: isSet(object.op) ? opTypeFromJSON(object.op) : 0,
      path: isSet(object.path) ? String(object.path) : "",
      value: isSet(object?.value) ? object.value : undefined,
      from: isSet(object.from) ? String(object.from) : "",
    };
  },

  toJSON(message: JsonPatch): unknown {
    const obj: any = {};
    if (message.op !== 0) {
      obj.op = opTypeToJSON(message.op);
    }
    if (message.path !== "") {
      obj.path = message.path;
    }
    if (message.value !== undefined) {
      obj.value = message.value;
    }
    if (message.from !== "") {
      obj.from = message.from;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<JsonPatch>, I>>(base?: I): JsonPatch {
    return JsonPatch.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<JsonPatch>, I>>(object: I): JsonPatch {
    const message = createBaseJsonPatch();
    message.op = object.op ?? 0;
    message.path = object.path ?? "";
    message.value = object.value ?? undefined;
    message.from = object.from ?? "";
    return message;
  },
};

function createBaseEvent(): Event {
  return { patchId: "", execMessage: "" };
}

export const Event = {
  encode(message: Event, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.patchId !== "") {
      writer.uint32(10).string(message.patchId);
    }
    if (message.execMessage !== "") {
      writer.uint32(18).string(message.execMessage);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Event {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.patchId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.execMessage = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Event {
    return {
      patchId: isSet(object.patchId) ? String(object.patchId) : "",
      execMessage: isSet(object.execMessage) ? String(object.execMessage) : "",
    };
  },

  toJSON(message: Event): unknown {
    const obj: any = {};
    if (message.patchId !== "") {
      obj.patchId = message.patchId;
    }
    if (message.execMessage !== "") {
      obj.execMessage = message.execMessage;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Event>, I>>(base?: I): Event {
    return Event.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Event>, I>>(object: I): Event {
    const message = createBaseEvent();
    message.patchId = object.patchId ?? "";
    message.execMessage = object.execMessage ?? "";
    return message;
  },
};

function createBaseSetupRequest(): SetupRequest {
  return { exec: undefined };
}

export const SetupRequest = {
  encode(message: SetupRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.exec !== undefined) {
      Executable.encode(message.exec, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetupRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetupRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.exec = Executable.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetupRequest {
    return { exec: isSet(object.exec) ? Executable.fromJSON(object.exec) : undefined };
  },

  toJSON(message: SetupRequest): unknown {
    const obj: any = {};
    if (message.exec !== undefined) {
      obj.exec = Executable.toJSON(message.exec);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetupRequest>, I>>(base?: I): SetupRequest {
    return SetupRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SetupRequest>, I>>(object: I): SetupRequest {
    const message = createBaseSetupRequest();
    message.exec = (object.exec !== undefined && object.exec !== null)
      ? Executable.fromPartial(object.exec)
      : undefined;
    return message;
  },
};

function createBaseSetupResponse(): SetupResponse {
  return { status: undefined };
}

export const SetupResponse = {
  encode(message: SetupResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== undefined) {
      Status.encode(message.status, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetupResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetupResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.status = Status.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetupResponse {
    return { status: isSet(object.status) ? Status.fromJSON(object.status) : undefined };
  },

  toJSON(message: SetupResponse): unknown {
    const obj: any = {};
    if (message.status !== undefined) {
      obj.status = Status.toJSON(message.status);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetupResponse>, I>>(base?: I): SetupResponse {
    return SetupResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SetupResponse>, I>>(object: I): SetupResponse {
    const message = createBaseSetupResponse();
    message.status = (object.status !== undefined && object.status !== null)
      ? Status.fromPartial(object.status)
      : undefined;
    return message;
  },
};

function createBaseCredential(): Credential {
  return { apiKey: "" };
}

export const Credential = {
  encode(message: Credential, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.apiKey !== "") {
      writer.uint32(10).string(message.apiKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Credential {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCredential();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.apiKey = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Credential {
    return { apiKey: isSet(object.apiKey) ? String(object.apiKey) : "" };
  },

  toJSON(message: Credential): unknown {
    const obj: any = {};
    if (message.apiKey !== "") {
      obj.apiKey = message.apiKey;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Credential>, I>>(base?: I): Credential {
    return Credential.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Credential>, I>>(object: I): Credential {
    const message = createBaseCredential();
    message.apiKey = object.apiKey ?? "";
    return message;
  },
};

function createBaseCredentialResponse(): CredentialResponse {
  return {};
}

export const CredentialResponse = {
  encode(_: CredentialResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CredentialResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCredentialResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): CredentialResponse {
    return {};
  },

  toJSON(_: CredentialResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CredentialResponse>, I>>(base?: I): CredentialResponse {
    return CredentialResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CredentialResponse>, I>>(_: I): CredentialResponse {
    const message = createBaseCredentialResponse();
    return message;
  },
};

function createBaseQuery(): Query {
  return { id: "", input: "" };
}

export const Query = {
  encode(message: Query, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.input !== "") {
      writer.uint32(18).string(message.input);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Query {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.input = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Query {
    return { id: isSet(object.id) ? String(object.id) : "", input: isSet(object.input) ? String(object.input) : "" };
  },

  toJSON(message: Query): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.input !== "") {
      obj.input = message.input;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Query>, I>>(base?: I): Query {
    return Query.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Query>, I>>(object: I): Query {
    const message = createBaseQuery();
    message.id = object.id ?? "";
    message.input = object.input ?? "";
    return message;
  },
};

function createBaseQueryResponse(): QueryResponse {
  return { status: undefined, id: "", results: "" };
}

export const QueryResponse = {
  encode(message: QueryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== undefined) {
      Status.encode(message.status, writer.uint32(10).fork()).ldelim();
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.results !== "") {
      writer.uint32(26).string(message.results);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.status = Status.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.id = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.results = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryResponse {
    return {
      status: isSet(object.status) ? Status.fromJSON(object.status) : undefined,
      id: isSet(object.id) ? String(object.id) : "",
      results: isSet(object.results) ? String(object.results) : "",
    };
  },

  toJSON(message: QueryResponse): unknown {
    const obj: any = {};
    if (message.status !== undefined) {
      obj.status = Status.toJSON(message.status);
    }
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.results !== "") {
      obj.results = message.results;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryResponse>, I>>(base?: I): QueryResponse {
    return QueryResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryResponse>, I>>(object: I): QueryResponse {
    const message = createBaseQueryResponse();
    message.status = (object.status !== undefined && object.status !== null)
      ? Status.fromPartial(object.status)
      : undefined;
    message.id = object.id ?? "";
    message.results = object.results ?? "";
    return message;
  },
};

function createBaseAction(): Action {
  return { id: "", patch: undefined };
}

export const Action = {
  encode(message: Action, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.patch !== undefined) {
      JsonPatch.encode(message.patch, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Action {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.patch = JsonPatch.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Action {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      patch: isSet(object.patch) ? JsonPatch.fromJSON(object.patch) : undefined,
    };
  },

  toJSON(message: Action): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.patch !== undefined) {
      obj.patch = JsonPatch.toJSON(message.patch);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Action>, I>>(base?: I): Action {
    return Action.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Action>, I>>(object: I): Action {
    const message = createBaseAction();
    message.id = object.id ?? "";
    message.patch = (object.patch !== undefined && object.patch !== null)
      ? JsonPatch.fromPartial(object.patch)
      : undefined;
    return message;
  },
};

function createBaseResponse(): Response {
  return { status: undefined, actionId: "", results: [] };
}

export const Response = {
  encode(message: Response, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== undefined) {
      Status.encode(message.status, writer.uint32(10).fork()).ldelim();
    }
    if (message.actionId !== "") {
      writer.uint32(18).string(message.actionId);
    }
    for (const v of message.results) {
      Value.encode(Value.wrap(v!), writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Response {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.status = Status.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.actionId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.results.push(Value.unwrap(Value.decode(reader, reader.uint32())));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Response {
    return {
      status: isSet(object.status) ? Status.fromJSON(object.status) : undefined,
      actionId: isSet(object.actionId) ? String(object.actionId) : "",
      results: Array.isArray(object?.results) ? [...object.results] : [],
    };
  },

  toJSON(message: Response): unknown {
    const obj: any = {};
    if (message.status !== undefined) {
      obj.status = Status.toJSON(message.status);
    }
    if (message.actionId !== "") {
      obj.actionId = message.actionId;
    }
    if (message.results?.length) {
      obj.results = message.results;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Response>, I>>(base?: I): Response {
    return Response.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Response>, I>>(object: I): Response {
    const message = createBaseResponse();
    message.status = (object.status !== undefined && object.status !== null)
      ? Status.fromPartial(object.status)
      : undefined;
    message.actionId = object.actionId ?? "";
    message.results = object.results?.map((e) => e) || [];
    return message;
  },
};

function createBaseCloseRequest(): CloseRequest {
  return {};
}

export const CloseRequest = {
  encode(_: CloseRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CloseRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCloseRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): CloseRequest {
    return {};
  },

  toJSON(_: CloseRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CloseRequest>, I>>(base?: I): CloseRequest {
    return CloseRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CloseRequest>, I>>(_: I): CloseRequest {
    const message = createBaseCloseRequest();
    return message;
  },
};

function createBaseCloseResponse(): CloseResponse {
  return {};
}

export const CloseResponse = {
  encode(_: CloseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CloseResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCloseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): CloseResponse {
    return {};
  },

  toJSON(_: CloseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CloseResponse>, I>>(base?: I): CloseResponse {
    return CloseResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CloseResponse>, I>>(_: I): CloseResponse {
    const message = createBaseCloseResponse();
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
