/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../eightball/params";
import { Fortunes } from "../eightball/fortunes";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage =
  "charleenfei.icqics20cosmoverseworkshop.eightball";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetFortunesRequest {
  owner: string;
}

export interface QueryGetFortunesResponse {
  fortunes: Fortunes | undefined;
}

export interface QueryAllFortunesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllFortunesResponse {
  fortunes: Fortunes[];
  pagination: PageResponse | undefined;
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

const baseQueryGetFortunesRequest: object = { owner: "" };

export const QueryGetFortunesRequest = {
  encode(
    message: QueryGetFortunesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFortunesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFortunesRequest,
    } as QueryGetFortunesRequest;
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

  fromJSON(object: any): QueryGetFortunesRequest {
    const message = {
      ...baseQueryGetFortunesRequest,
    } as QueryGetFortunesRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: QueryGetFortunesRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFortunesRequest>
  ): QueryGetFortunesRequest {
    const message = {
      ...baseQueryGetFortunesRequest,
    } as QueryGetFortunesRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseQueryGetFortunesResponse: object = {};

export const QueryGetFortunesResponse = {
  encode(
    message: QueryGetFortunesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fortunes !== undefined) {
      Fortunes.encode(message.fortunes, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetFortunesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFortunesResponse,
    } as QueryGetFortunesResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fortunes = Fortunes.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFortunesResponse {
    const message = {
      ...baseQueryGetFortunesResponse,
    } as QueryGetFortunesResponse;
    if (object.fortunes !== undefined && object.fortunes !== null) {
      message.fortunes = Fortunes.fromJSON(object.fortunes);
    } else {
      message.fortunes = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFortunesResponse): unknown {
    const obj: any = {};
    message.fortunes !== undefined &&
      (obj.fortunes = message.fortunes
        ? Fortunes.toJSON(message.fortunes)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFortunesResponse>
  ): QueryGetFortunesResponse {
    const message = {
      ...baseQueryGetFortunesResponse,
    } as QueryGetFortunesResponse;
    if (object.fortunes !== undefined && object.fortunes !== null) {
      message.fortunes = Fortunes.fromPartial(object.fortunes);
    } else {
      message.fortunes = undefined;
    }
    return message;
  },
};

const baseQueryAllFortunesRequest: object = {};

export const QueryAllFortunesRequest = {
  encode(
    message: QueryAllFortunesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFortunesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllFortunesRequest,
    } as QueryAllFortunesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFortunesRequest {
    const message = {
      ...baseQueryAllFortunesRequest,
    } as QueryAllFortunesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFortunesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFortunesRequest>
  ): QueryAllFortunesRequest {
    const message = {
      ...baseQueryAllFortunesRequest,
    } as QueryAllFortunesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllFortunesResponse: object = {};

export const QueryAllFortunesResponse = {
  encode(
    message: QueryAllFortunesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.fortunes) {
      Fortunes.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllFortunesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllFortunesResponse,
    } as QueryAllFortunesResponse;
    message.fortunes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fortunes.push(Fortunes.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFortunesResponse {
    const message = {
      ...baseQueryAllFortunesResponse,
    } as QueryAllFortunesResponse;
    message.fortunes = [];
    if (object.fortunes !== undefined && object.fortunes !== null) {
      for (const e of object.fortunes) {
        message.fortunes.push(Fortunes.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFortunesResponse): unknown {
    const obj: any = {};
    if (message.fortunes) {
      obj.fortunes = message.fortunes.map((e) =>
        e ? Fortunes.toJSON(e) : undefined
      );
    } else {
      obj.fortunes = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFortunesResponse>
  ): QueryAllFortunesResponse {
    const message = {
      ...baseQueryAllFortunesResponse,
    } as QueryAllFortunesResponse;
    message.fortunes = [];
    if (object.fortunes !== undefined && object.fortunes !== null) {
      for (const e of object.fortunes) {
        message.fortunes.push(Fortunes.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Fortunes by index. */
  Fortunes(request: QueryGetFortunesRequest): Promise<QueryGetFortunesResponse>;
  /** Queries a list of Fortunes items. */
  FortunesAll(
    request: QueryAllFortunesRequest
  ): Promise<QueryAllFortunesResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "charleenfei.icqics20cosmoverseworkshop.eightball.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Fortunes(
    request: QueryGetFortunesRequest
  ): Promise<QueryGetFortunesResponse> {
    const data = QueryGetFortunesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "charleenfei.icqics20cosmoverseworkshop.eightball.Query",
      "Fortunes",
      data
    );
    return promise.then((data) =>
      QueryGetFortunesResponse.decode(new Reader(data))
    );
  }

  FortunesAll(
    request: QueryAllFortunesRequest
  ): Promise<QueryAllFortunesResponse> {
    const data = QueryAllFortunesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "charleenfei.icqics20cosmoverseworkshop.eightball.Query",
      "FortunesAll",
      data
    );
    return promise.then((data) =>
      QueryAllFortunesResponse.decode(new Reader(data))
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
