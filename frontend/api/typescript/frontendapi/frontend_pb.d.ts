// @generated by protoc-gen-es v1.9.0
// @generated from file frontendapi/frontend.proto (package frontendapi, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
  BinaryReadOptions,
  FieldList,
  JsonReadOptions,
  JsonValue,
  PartialMessage,
  PlainMessage,
} from "@bufbuild/protobuf";
import { Message, type proto3 } from "@bufbuild/protobuf";

/**
 * The settings for a user.
 *
 * @generated from message frontendapi.User
 */
export declare class User extends Message<User> {
  /**
   * IDs of programming languages that reviews can be created for.
   * IDs correspond to `language_id` from github-linguist.
   * https://github.com/github-linguist/linguist/blob/master/lib/linguist/languages.yml
   * Required.
   *
   * @generated from field: repeated uint32 programming_language_ids = 1;
   */
  programmingLanguageIds: number[];

  /**
   * The maximum number of reviews created by the app that can be open at once.
   * Required.
   *
   * @generated from field: uint32 max_open_reviews = 2;
   */
  maxOpenReviews: number;

  constructor(data?: PartialMessage<User>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "frontendapi.User";
  static readonly fields: FieldList;

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): User;

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): User;

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): User;

  static equals(
    a: User | PlainMessage<User> | undefined,
    b: User | PlainMessage<User> | undefined,
  ): boolean;
}

/**
 * A request for FrontendService.SaveUser.
 *
 * @generated from message frontendapi.SaveUserRequest
 */
export declare class SaveUserRequest extends Message<SaveUserRequest> {
  /**
   * The user to create.
   * Required.
   *
   * @generated from field: optional frontendapi.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<SaveUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "frontendapi.SaveUserRequest";
  static readonly fields: FieldList;

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): SaveUserRequest;

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): SaveUserRequest;

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): SaveUserRequest;

  static equals(
    a: SaveUserRequest | PlainMessage<SaveUserRequest> | undefined,
    b: SaveUserRequest | PlainMessage<SaveUserRequest> | undefined,
  ): boolean;
}

/**
 * A response for FrontendService.SaveUser.
 *
 * Empty to allow future extension.
 *
 * @generated from message frontendapi.SaveUserResponse
 */
export declare class SaveUserResponse extends Message<SaveUserResponse> {
  constructor(data?: PartialMessage<SaveUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "frontendapi.SaveUserResponse";
  static readonly fields: FieldList;

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): SaveUserResponse;

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): SaveUserResponse;

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): SaveUserResponse;

  static equals(
    a: SaveUserResponse | PlainMessage<SaveUserResponse> | undefined,
    b: SaveUserResponse | PlainMessage<SaveUserResponse> | undefined,
  ): boolean;
}
