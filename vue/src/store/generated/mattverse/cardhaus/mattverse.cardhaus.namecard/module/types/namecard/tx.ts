/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "mattverse.cardhaus.namecard";

export interface MsgCreateNamecard {
  address: string;
  name: string;
  imageUrl: string;
  selfIntro: string;
}

export interface MsgCreateNamecardResponse {}

const baseMsgCreateNamecard: object = {
  address: "",
  name: "",
  imageUrl: "",
  selfIntro: "",
};

export const MsgCreateNamecard = {
  encode(message: MsgCreateNamecard, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.imageUrl !== "") {
      writer.uint32(26).string(message.imageUrl);
    }
    if (message.selfIntro !== "") {
      writer.uint32(34).string(message.selfIntro);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateNamecard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateNamecard } as MsgCreateNamecard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.imageUrl = reader.string();
          break;
        case 4:
          message.selfIntro = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateNamecard {
    const message = { ...baseMsgCreateNamecard } as MsgCreateNamecard;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.imageUrl !== undefined && object.imageUrl !== null) {
      message.imageUrl = String(object.imageUrl);
    } else {
      message.imageUrl = "";
    }
    if (object.selfIntro !== undefined && object.selfIntro !== null) {
      message.selfIntro = String(object.selfIntro);
    } else {
      message.selfIntro = "";
    }
    return message;
  },

  toJSON(message: MsgCreateNamecard): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.name !== undefined && (obj.name = message.name);
    message.imageUrl !== undefined && (obj.imageUrl = message.imageUrl);
    message.selfIntro !== undefined && (obj.selfIntro = message.selfIntro);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateNamecard>): MsgCreateNamecard {
    const message = { ...baseMsgCreateNamecard } as MsgCreateNamecard;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.imageUrl !== undefined && object.imageUrl !== null) {
      message.imageUrl = object.imageUrl;
    } else {
      message.imageUrl = "";
    }
    if (object.selfIntro !== undefined && object.selfIntro !== null) {
      message.selfIntro = object.selfIntro;
    } else {
      message.selfIntro = "";
    }
    return message;
  },
};

const baseMsgCreateNamecardResponse: object = {};

export const MsgCreateNamecardResponse = {
  encode(
    _: MsgCreateNamecardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateNamecardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateNamecardResponse,
    } as MsgCreateNamecardResponse;
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

  fromJSON(_: any): MsgCreateNamecardResponse {
    const message = {
      ...baseMsgCreateNamecardResponse,
    } as MsgCreateNamecardResponse;
    return message;
  },

  toJSON(_: MsgCreateNamecardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateNamecardResponse>
  ): MsgCreateNamecardResponse {
    const message = {
      ...baseMsgCreateNamecardResponse,
    } as MsgCreateNamecardResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateNamecard(
    request: MsgCreateNamecard
  ): Promise<MsgCreateNamecardResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateNamecard(
    request: MsgCreateNamecard
  ): Promise<MsgCreateNamecardResponse> {
    const data = MsgCreateNamecard.encode(request).finish();
    const promise = this.rpc.request(
      "mattverse.cardhaus.namecard.Msg",
      "CreateNamecard",
      data
    );
    return promise.then((data) =>
      MsgCreateNamecardResponse.decode(new Reader(data))
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
