/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "eightball.v1";

export interface Fortune {
  owner: string;
  price: string;
  fortune: string;
}

const baseFortune: object = { owner: "", price: "", fortune: "" };

export const Fortune = {
  encode(message: Fortune, writer: Writer = Writer.create()): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.price !== "") {
      writer.uint32(18).string(message.price);
    }
    if (message.fortune !== "") {
      writer.uint32(26).string(message.fortune);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Fortune {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFortune } as Fortune;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.price = reader.string();
          break;
        case 3:
          message.fortune = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Fortune {
    const message = { ...baseFortune } as Fortune;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.fortune !== undefined && object.fortune !== null) {
      message.fortune = String(object.fortune);
    } else {
      message.fortune = "";
    }
    return message;
  },

  toJSON(message: Fortune): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.price !== undefined && (obj.price = message.price);
    message.fortune !== undefined && (obj.fortune = message.fortune);
    return obj;
  },

  fromPartial(object: DeepPartial<Fortune>): Fortune {
    const message = { ...baseFortune } as Fortune;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.fortune !== undefined && object.fortune !== null) {
      message.fortune = object.fortune;
    } else {
      message.fortune = "";
    }
    return message;
  },
};

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
