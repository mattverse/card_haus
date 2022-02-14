import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../namecard/params";
import { NameCard } from "../namecard/nameCard";
export declare const protobufPackage = "mattverse.cardhaus.namecard";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryCardInfoRequest {
    address: string;
}
export interface QueryCardInfoResponse {
    nameCard: NameCard | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryCardInfoRequest: {
    encode(message: QueryCardInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryCardInfoRequest;
    fromJSON(object: any): QueryCardInfoRequest;
    toJSON(message: QueryCardInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryCardInfoRequest>): QueryCardInfoRequest;
};
export declare const QueryCardInfoResponse: {
    encode(message: QueryCardInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryCardInfoResponse;
    fromJSON(object: any): QueryCardInfoResponse;
    toJSON(message: QueryCardInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryCardInfoResponse>): QueryCardInfoResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of CardInfo items. */
    CardInfo(request: QueryCardInfoRequest): Promise<QueryCardInfoResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    CardInfo(request: QueryCardInfoRequest): Promise<QueryCardInfoResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
