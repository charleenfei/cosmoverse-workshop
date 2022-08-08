/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage =
  "charleenfei.icqics20cosmoverseworkshop.eightball";

export interface Fortuneteller {
  index: string;
  owner: string;
  fortune: string;
  expiry: string;
}

const baseFortuneteller: object = {
  index: "",
  owner: "",
  fortune: "",
  expiry: "",
};

export const Fortuneteller = {
  encode(message: Fortuneteller, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.fortune !== "") {
      writer.uint32(26).string(message.fortune);
    }
    if (message.expiry !== "") {
      writer.uint32(34).string(message.expiry);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Fortuneteller {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFortuneteller } as Fortuneteller;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.fortune = reader.string();
          break;
        case 4:
          message.expiry = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Fortuneteller {
    const message = { ...baseFortuneteller } as Fortuneteller;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.fortune !== undefined && object.fortune !== null) {
      message.fortune = String(object.fortune);
    } else {
      message.fortune = "";
    }
    if (object.expiry !== undefined && object.expiry !== null) {
      message.expiry = String(object.expiry);
    } else {
      message.expiry = "";
    }
    return message;
  },

  toJSON(message: Fortuneteller): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.owner !== undefined && (obj.owner = message.owner);
    message.fortune !== undefined && (obj.fortune = message.fortune);
    message.expiry !== undefined && (obj.expiry = message.expiry);
    return obj;
  },

  fromPartial(object: DeepPartial<Fortuneteller>): Fortuneteller {
    const message = { ...baseFortuneteller } as Fortuneteller;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.fortune !== undefined && object.fortune !== null) {
      message.fortune = object.fortune;
    } else {
      message.fortune = "";
    }
    if (object.expiry !== undefined && object.expiry !== null) {
      message.expiry = object.expiry;
    } else {
      message.expiry = "";
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
