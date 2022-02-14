/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "mattverse.cardhaus.namecard";
const baseHistory = { address: "" };
export const History = {
    encode(message, writer = Writer.create()) {
        if (message.sentTime !== undefined) {
            Timestamp.encode(toTimestamp(message.sentTime), writer.uint32(10).fork()).ldelim();
        }
        if (message.address !== "") {
            writer.uint32(18).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseHistory };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.sentTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseHistory };
        if (object.sentTime !== undefined && object.sentTime !== null) {
            message.sentTime = fromJsonTimestamp(object.sentTime);
        }
        else {
            message.sentTime = undefined;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.sentTime !== undefined &&
            (obj.sentTime =
                message.sentTime !== undefined ? message.sentTime.toISOString() : null);
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseHistory };
        if (object.sentTime !== undefined && object.sentTime !== null) {
            message.sentTime = object.sentTime;
        }
        else {
            message.sentTime = undefined;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        return message;
    },
};
const baseNameCard = {
    address: "",
    name: "",
    imageUrl: "",
    selfIntro: "",
};
export const NameCard = {
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
        for (const v of message.sentInfo) {
            History.encode(v, writer.uint32(42).fork()).ldelim();
        }
        for (const v of message.receivedInfo) {
            History.encode(v, writer.uint32(50).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNameCard };
        message.sentInfo = [];
        message.receivedInfo = [];
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
                case 5:
                    message.sentInfo.push(History.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.receivedInfo.push(History.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseNameCard };
        message.sentInfo = [];
        message.receivedInfo = [];
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
        if (object.sentInfo !== undefined && object.sentInfo !== null) {
            for (const e of object.sentInfo) {
                message.sentInfo.push(History.fromJSON(e));
            }
        }
        if (object.receivedInfo !== undefined && object.receivedInfo !== null) {
            for (const e of object.receivedInfo) {
                message.receivedInfo.push(History.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.name !== undefined && (obj.name = message.name);
        message.imageUrl !== undefined && (obj.imageUrl = message.imageUrl);
        message.selfIntro !== undefined && (obj.selfIntro = message.selfIntro);
        if (message.sentInfo) {
            obj.sentInfo = message.sentInfo.map((e) => e ? History.toJSON(e) : undefined);
        }
        else {
            obj.sentInfo = [];
        }
        if (message.receivedInfo) {
            obj.receivedInfo = message.receivedInfo.map((e) => e ? History.toJSON(e) : undefined);
        }
        else {
            obj.receivedInfo = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseNameCard };
        message.sentInfo = [];
        message.receivedInfo = [];
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
        if (object.sentInfo !== undefined && object.sentInfo !== null) {
            for (const e of object.sentInfo) {
                message.sentInfo.push(History.fromPartial(e));
            }
        }
        if (object.receivedInfo !== undefined && object.receivedInfo !== null) {
            for (const e of object.receivedInfo) {
                message.receivedInfo.push(History.fromPartial(e));
            }
        }
        return message;
    },
};
function toTimestamp(date) {
    const seconds = date.getTime() / 1000;
    const nanos = (date.getTime() % 1000) * 1000000;
    return { seconds, nanos };
}
function fromTimestamp(t) {
    let millis = t.seconds * 1000;
    millis += t.nanos / 1000000;
    return new Date(millis);
}
function fromJsonTimestamp(o) {
    if (o instanceof Date) {
        return o;
    }
    else if (typeof o === "string") {
        return new Date(o);
    }
    else {
        return fromTimestamp(Timestamp.fromJSON(o));
    }
}
