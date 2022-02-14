/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "mattverse.cardhaus.namecard";
const baseMsgCreateNamecard = {
    address: "",
    name: "",
    imageUrl: "",
    selfIntro: "",
};
export const MsgCreateNamecard = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateNamecard };
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
    fromJSON(object) {
        const message = { ...baseMsgCreateNamecard };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.imageUrl !== undefined && object.imageUrl !== null) {
            message.imageUrl = String(object.imageUrl);
        }
        else {
            message.imageUrl = "";
        }
        if (object.selfIntro !== undefined && object.selfIntro !== null) {
            message.selfIntro = String(object.selfIntro);
        }
        else {
            message.selfIntro = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.name !== undefined && (obj.name = message.name);
        message.imageUrl !== undefined && (obj.imageUrl = message.imageUrl);
        message.selfIntro !== undefined && (obj.selfIntro = message.selfIntro);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateNamecard };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.imageUrl !== undefined && object.imageUrl !== null) {
            message.imageUrl = object.imageUrl;
        }
        else {
            message.imageUrl = "";
        }
        if (object.selfIntro !== undefined && object.selfIntro !== null) {
            message.selfIntro = object.selfIntro;
        }
        else {
            message.selfIntro = "";
        }
        return message;
    },
};
const baseMsgCreateNamecardResponse = {};
export const MsgCreateNamecardResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateNamecardResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgCreateNamecardResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateNamecardResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateNamecard(request) {
        const data = MsgCreateNamecard.encode(request).finish();
        const promise = this.rpc.request("mattverse.cardhaus.namecard.Msg", "CreateNamecard", data);
        return promise.then((data) => MsgCreateNamecardResponse.decode(new Reader(data)));
    }
}
