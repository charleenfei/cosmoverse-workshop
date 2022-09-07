/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "eightball.v1";

export interface Fortune {
  owner: string;
  price: string;
  fortune: string;
}

export interface FortuneList {
  fortunes: Fortune[];
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

const baseFortuneList: object = {};

export const FortuneList = {
  encode(message: FortuneList, writer: Writer = Writer.create()): Writer {
    for (const v of message.fortunes) {
      Fortune.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): FortuneList {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFortuneList } as FortuneList;
    message.fortunes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fortunes.push(Fortune.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FortuneList {
    const message = { ...baseFortuneList } as FortuneList;
    message.fortunes = [];
    if (object.fortunes !== undefined && object.fortunes !== null) {
      for (const e of object.fortunes) {
        message.fortunes.push(Fortune.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: FortuneList): unknown {
    const obj: any = {};
    if (message.fortunes) {
      obj.fortunes = message.fortunes.map((e) =>
        e ? Fortune.toJSON(e) : undefined
      );
    } else {
      obj.fortunes = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<FortuneList>): FortuneList {
    const message = { ...baseFortuneList } as FortuneList;
    message.fortunes = [];
    if (object.fortunes !== undefined && object.fortunes !== null) {
      for (const e of object.fortunes) {
        message.fortunes.push(Fortune.fromPartial(e));
      }
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
