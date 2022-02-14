import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "mattverse.cardhaus.namecard";
export interface History {
    sentTime: Date | undefined;
    address: string;
}
export interface NameCard {
    address: string;
    name: string;
    imageUrl: string;
    selfIntro: string;
    sentInfo: History[];
    receivedInfo: History[];
}
export declare const History: {
    encode(message: History, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): History;
    fromJSON(object: any): History;
    toJSON(message: History): unknown;
    fromPartial(object: DeepPartial<History>): History;
};
export declare const NameCard: {
    encode(message: NameCard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NameCard;
    fromJSON(object: any): NameCard;
    toJSON(message: NameCard): unknown;
    fromPartial(object: DeepPartial<NameCard>): NameCard;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
