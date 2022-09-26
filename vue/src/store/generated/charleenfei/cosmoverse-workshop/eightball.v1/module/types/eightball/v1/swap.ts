/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "eightball.v1";

export interface MsgSwap {
  sender: string;
  offer: Coin | undefined;
  minAsk: Coin | undefined;
  /** information for how to send the tokens to intended receiver */
  port_id: string;
  channel_id: string;
  receiver: string;
}

export interface MsgSwapResponse {
  sequence: number;
}

const baseMsgSwap: object = {
  sender: "",
  port_id: "",
  channel_id: "",
  receiver: "",
};

export const MsgSwap = {
  encode(message: MsgSwap, writer: Writer = Writer.create()): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.offer !== undefined) {
      Coin.encode(message.offer, writer.uint32(18).fork()).ldelim();
    }
    if (message.minAsk !== undefined) {
      Coin.encode(message.minAsk, writer.uint32(26).fork()).ldelim();
    }
    if (message.port_id !== "") {
      writer.uint32(34).string(message.port_id);
    }
    if (message.channel_id !== "") {
      writer.uint32(42).string(message.channel_id);
    }
    if (message.receiver !== "") {
      writer.uint32(50).string(message.receiver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwap } as MsgSwap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.offer = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.minAsk = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.port_id = reader.string();
          break;
        case 5:
          message.channel_id = reader.string();
          break;
        case 6:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.offer !== undefined && object.offer !== null) {
      message.offer = Coin.fromJSON(object.offer);
    } else {
      message.offer = undefined;
    }
    if (object.minAsk !== undefined && object.minAsk !== null) {
      message.minAsk = Coin.fromJSON(object.minAsk);
    } else {
      message.minAsk = undefined;
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = String(object.port_id);
    } else {
      message.port_id = "";
    }
    if (object.channel_id !== undefined && object.channel_id !== null) {
      message.channel_id = String(object.channel_id);
    } else {
      message.channel_id = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    return message;
  },

  toJSON(message: MsgSwap): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.offer !== undefined &&
      (obj.offer = message.offer ? Coin.toJSON(message.offer) : undefined);
    message.minAsk !== undefined &&
      (obj.minAsk = message.minAsk ? Coin.toJSON(message.minAsk) : undefined);
    message.port_id !== undefined && (obj.port_id = message.port_id);
    message.channel_id !== undefined && (obj.channel_id = message.channel_id);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwap>): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.offer !== undefined && object.offer !== null) {
      message.offer = Coin.fromPartial(object.offer);
    } else {
      message.offer = undefined;
    }
    if (object.minAsk !== undefined && object.minAsk !== null) {
      message.minAsk = Coin.fromPartial(object.minAsk);
    } else {
      message.minAsk = undefined;
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = object.port_id;
    } else {
      message.port_id = "";
    }
    if (object.channel_id !== undefined && object.channel_id !== null) {
      message.channel_id = object.channel_id;
    } else {
      message.channel_id = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    return message;
  },
};

const baseMsgSwapResponse: object = { sequence: 0 };

export const MsgSwapResponse = {
  encode(message: MsgSwapResponse, writer: Writer = Writer.create()): Writer {
    if (message.sequence !== 0) {
      writer.uint32(8).uint64(message.sequence);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sequence = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = Number(object.sequence);
    } else {
      message.sequence = 0;
    }
    return message;
  },

  toJSON(message: MsgSwapResponse): unknown {
    const obj: any = {};
    message.sequence !== undefined && (obj.sequence = message.sequence);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwapResponse>): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = object.sequence;
    } else {
      message.sequence = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
