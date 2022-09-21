/* eslint-disable */
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "eightball.v1";

export interface Request {
  offerer: string;
  swapped_coin: Coin | undefined;
}

const baseRequest: object = { offerer: "" };

export const Request = {
  encode(message: Request, writer: Writer = Writer.create()): Writer {
    if (message.offerer !== "") {
      writer.uint32(10).string(message.offerer);
    }
    if (message.swapped_coin !== undefined) {
      Coin.encode(message.swapped_coin, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Request {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRequest } as Request;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.offerer = reader.string();
          break;
        case 2:
          message.swapped_coin = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Request {
    const message = { ...baseRequest } as Request;
    if (object.offerer !== undefined && object.offerer !== null) {
      message.offerer = String(object.offerer);
    } else {
      message.offerer = "";
    }
    if (object.swapped_coin !== undefined && object.swapped_coin !== null) {
      message.swapped_coin = Coin.fromJSON(object.swapped_coin);
    } else {
      message.swapped_coin = undefined;
    }
    return message;
  },

  toJSON(message: Request): unknown {
    const obj: any = {};
    message.offerer !== undefined && (obj.offerer = message.offerer);
    message.swapped_coin !== undefined &&
      (obj.swapped_coin = message.swapped_coin
        ? Coin.toJSON(message.swapped_coin)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Request>): Request {
    const message = { ...baseRequest } as Request;
    if (object.offerer !== undefined && object.offerer !== null) {
      message.offerer = object.offerer;
    } else {
      message.offerer = "";
    }
    if (object.swapped_coin !== undefined && object.swapped_coin !== null) {
      message.swapped_coin = Coin.fromPartial(object.swapped_coin);
    } else {
      message.swapped_coin = undefined;
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
