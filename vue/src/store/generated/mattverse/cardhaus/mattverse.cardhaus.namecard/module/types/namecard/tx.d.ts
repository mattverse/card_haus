import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "mattverse.cardhaus.namecard";
export interface MsgCreateNamecard {
    address: string;
    name: string;
    imageUrl: string;
    selfIntro: string;
}
export interface MsgCreateNamecardResponse {
}
export declare const MsgCreateNamecard: {
    encode(message: MsgCreateNamecard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateNamecard;
    fromJSON(object: any): MsgCreateNamecard;
    toJSON(message: MsgCreateNamecard): unknown;
    fromPartial(object: DeepPartial<MsgCreateNamecard>): MsgCreateNamecard;
};
export declare const MsgCreateNamecardResponse: {
    encode(_: MsgCreateNamecardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateNamecardResponse;
    fromJSON(_: any): MsgCreateNamecardResponse;
    toJSON(_: MsgCreateNamecardResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateNamecardResponse>): MsgCreateNamecardResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateNamecard(request: MsgCreateNamecard): Promise<MsgCreateNamecardResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateNamecard(request: MsgCreateNamecard): Promise<MsgCreateNamecardResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
