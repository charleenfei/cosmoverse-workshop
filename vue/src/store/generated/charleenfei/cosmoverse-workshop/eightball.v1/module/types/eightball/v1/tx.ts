/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "eightball.v1";

export interface MsgFeelingLucky {
  sender: string;
  /** TODO: non-nullable gogoproto */
  offering: Coin | undefined;
}

export interface MsgFeelingLuckyResponse {}

export interface MsgConnectToDex {
  creator: string;
  connection_id: string;
}

export interface MsgConnectToDexResponse {}

const baseMsgFeelingLucky: object = { sender: "" };

export const MsgFeelingLucky = {
  encode(message: MsgFeelingLucky, writer: Writer = Writer.create()): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.offering !== undefined) {
      Coin.encode(message.offering, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFeelingLucky {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFeelingLucky } as MsgFeelingLucky;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.offering = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFeelingLucky {
    const message = { ...baseMsgFeelingLucky } as MsgFeelingLucky;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.offering !== undefined && object.offering !== null) {
      message.offering = Coin.fromJSON(object.offering);
    } else {
      message.offering = undefined;
    }
    return message;
  },

  toJSON(message: MsgFeelingLucky): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.offering !== undefined &&
      (obj.offering = message.offering
        ? Coin.toJSON(message.offering)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgFeelingLucky>): MsgFeelingLucky {
    const message = { ...baseMsgFeelingLucky } as MsgFeelingLucky;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.offering !== undefined && object.offering !== null) {
      message.offering = Coin.fromPartial(object.offering);
    } else {
      message.offering = undefined;
    }
    return message;
  },
};

const baseMsgFeelingLuckyResponse: object = {};

export const MsgFeelingLuckyResponse = {
  encode(_: MsgFeelingLuckyResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFeelingLuckyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgFeelingLuckyResponse,
    } as MsgFeelingLuckyResponse;
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

  fromJSON(_: any): MsgFeelingLuckyResponse {
    const message = {
      ...baseMsgFeelingLuckyResponse,
    } as MsgFeelingLuckyResponse;
    return message;
  },

  toJSON(_: MsgFeelingLuckyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgFeelingLuckyResponse>
  ): MsgFeelingLuckyResponse {
    const message = {
      ...baseMsgFeelingLuckyResponse,
    } as MsgFeelingLuckyResponse;
    return message;
  },
};

const baseMsgConnectToDex: object = { creator: "", connection_id: "" };

export const MsgConnectToDex = {
  encode(message: MsgConnectToDex, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.connection_id !== "") {
      writer.uint32(18).string(message.connection_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConnectToDex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgConnectToDex } as MsgConnectToDex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.connection_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConnectToDex {
    const message = { ...baseMsgConnectToDex } as MsgConnectToDex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = String(object.connection_id);
    } else {
      message.connection_id = "";
    }
    return message;
  },

  toJSON(message: MsgConnectToDex): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.connection_id !== undefined &&
      (obj.connection_id = message.connection_id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgConnectToDex>): MsgConnectToDex {
    const message = { ...baseMsgConnectToDex } as MsgConnectToDex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.connection_id !== undefined && object.connection_id !== null) {
      message.connection_id = object.connection_id;
    } else {
      message.connection_id = "";
    }
    return message;
  },
};

const baseMsgConnectToDexResponse: object = {};

export const MsgConnectToDexResponse = {
  encode(_: MsgConnectToDexResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConnectToDexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgConnectToDexResponse,
    } as MsgConnectToDexResponse;
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

  fromJSON(_: any): MsgConnectToDexResponse {
    const message = {
      ...baseMsgConnectToDexResponse,
    } as MsgConnectToDexResponse;
    return message;
  },

  toJSON(_: MsgConnectToDexResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgConnectToDexResponse>
  ): MsgConnectToDexResponse {
    const message = {
      ...baseMsgConnectToDexResponse,
    } as MsgConnectToDexResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  FeelingLucky(request: MsgFeelingLucky): Promise<MsgFeelingLuckyResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ConnectToDex(request: MsgConnectToDex): Promise<MsgConnectToDexResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  FeelingLucky(request: MsgFeelingLucky): Promise<MsgFeelingLuckyResponse> {
    const data = MsgFeelingLucky.encode(request).finish();
    const promise = this.rpc.request("eightball.v1.Msg", "FeelingLucky", data);
    return promise.then((data) =>
      MsgFeelingLuckyResponse.decode(new Reader(data))
    );
  }

  ConnectToDex(request: MsgConnectToDex): Promise<MsgConnectToDexResponse> {
    const data = MsgConnectToDex.encode(request).finish();
    const promise = this.rpc.request("eightball.v1.Msg", "ConnectToDex", data);
    return promise.then((data) =>
      MsgConnectToDexResponse.decode(new Reader(data))
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
