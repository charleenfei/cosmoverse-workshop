/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../../eightball/v1/params";
import { Fortune } from "../../eightball/v1/fortunes";

export const protobufPackage = "eightball.v1";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryFortuneRequest {
  owner: string;
}

export interface QueryFortuneResponse {
  fortune: Fortune | undefined;
}

export interface QueryOwnedFortunesRequest {}

export interface QueryOwnedFortunesResponse {
  owned_fortunes: Fortune[];
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryFortuneRequest: object = { owner: "" };

export const QueryFortuneRequest = {
  encode(
    message: QueryFortuneRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFortuneRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryFortuneRequest } as QueryFortuneRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFortuneRequest {
    const message = { ...baseQueryFortuneRequest } as QueryFortuneRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: QueryFortuneRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryFortuneRequest>): QueryFortuneRequest {
    const message = { ...baseQueryFortuneRequest } as QueryFortuneRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseQueryFortuneResponse: object = {};

export const QueryFortuneResponse = {
  encode(
    message: QueryFortuneResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fortune !== undefined) {
      Fortune.encode(message.fortune, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFortuneResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryFortuneResponse } as QueryFortuneResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fortune = Fortune.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFortuneResponse {
    const message = { ...baseQueryFortuneResponse } as QueryFortuneResponse;
    if (object.fortune !== undefined && object.fortune !== null) {
      message.fortune = Fortune.fromJSON(object.fortune);
    } else {
      message.fortune = undefined;
    }
    return message;
  },

  toJSON(message: QueryFortuneResponse): unknown {
    const obj: any = {};
    message.fortune !== undefined &&
      (obj.fortune = message.fortune
        ? Fortune.toJSON(message.fortune)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryFortuneResponse>): QueryFortuneResponse {
    const message = { ...baseQueryFortuneResponse } as QueryFortuneResponse;
    if (object.fortune !== undefined && object.fortune !== null) {
      message.fortune = Fortune.fromPartial(object.fortune);
    } else {
      message.fortune = undefined;
    }
    return message;
  },
};

const baseQueryOwnedFortunesRequest: object = {};

export const QueryOwnedFortunesRequest = {
  encode(
    _: QueryOwnedFortunesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOwnedFortunesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOwnedFortunesRequest,
    } as QueryOwnedFortunesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryOwnedFortunesRequest {
    const message = {
      ...baseQueryOwnedFortunesRequest,
    } as QueryOwnedFortunesRequest;
    return message;
  },

  toJSON(_: QueryOwnedFortunesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryOwnedFortunesRequest>
  ): QueryOwnedFortunesRequest {
    const message = {
      ...baseQueryOwnedFortunesRequest,
    } as QueryOwnedFortunesRequest;
    return message;
  },
};

const baseQueryOwnedFortunesResponse: object = {};

export const QueryOwnedFortunesResponse = {
  encode(
    message: QueryOwnedFortunesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.owned_fortunes) {
      Fortune.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOwnedFortunesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOwnedFortunesResponse,
    } as QueryOwnedFortunesResponse;
    message.owned_fortunes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owned_fortunes.push(Fortune.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOwnedFortunesResponse {
    const message = {
      ...baseQueryOwnedFortunesResponse,
    } as QueryOwnedFortunesResponse;
    message.owned_fortunes = [];
    if (object.owned_fortunes !== undefined && object.owned_fortunes !== null) {
      for (const e of object.owned_fortunes) {
        message.owned_fortunes.push(Fortune.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryOwnedFortunesResponse): unknown {
    const obj: any = {};
    if (message.owned_fortunes) {
      obj.owned_fortunes = message.owned_fortunes.map((e) =>
        e ? Fortune.toJSON(e) : undefined
      );
    } else {
      obj.owned_fortunes = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryOwnedFortunesResponse>
  ): QueryOwnedFortunesResponse {
    const message = {
      ...baseQueryOwnedFortunesResponse,
    } as QueryOwnedFortunesResponse;
    message.owned_fortunes = [];
    if (object.owned_fortunes !== undefined && object.owned_fortunes !== null) {
      for (const e of object.owned_fortunes) {
        message.owned_fortunes.push(Fortune.fromPartial(e));
      }
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a Fortune by owner. */
  Fortune(request: QueryFortuneRequest): Promise<QueryFortuneResponse>;
  /** Queries a list of owned fortunes. */
  Fortunes(
    request: QueryOwnedFortunesRequest
  ): Promise<QueryOwnedFortunesResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Fortune(request: QueryFortuneRequest): Promise<QueryFortuneResponse> {
    const data = QueryFortuneRequest.encode(request).finish();
    const promise = this.rpc.request("eightball.v1.Query", "Fortune", data);
    return promise.then((data) =>
      QueryFortuneResponse.decode(new Reader(data))
    );
  }

  Fortunes(
    request: QueryOwnedFortunesRequest
  ): Promise<QueryOwnedFortunesResponse> {
    const data = QueryOwnedFortunesRequest.encode(request).finish();
    const promise = this.rpc.request("eightball.v1.Query", "Fortunes", data);
    return promise.then((data) =>
      QueryOwnedFortunesResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
