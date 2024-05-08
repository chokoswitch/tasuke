// @generated by protoc-gen-es v1.9.0
// @generated from file frontendapi/frontend.proto (package frontendapi, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * The settings for a user.
 *
 * @generated from message frontendapi.User
 */
export const User = /*@__PURE__*/ proto3.makeMessageType(
  "frontendapi.User",
  () => [
    {
      no: 1,
      name: "programming_language_ids",
      kind: "scalar",
      T: 13 /* ScalarType.UINT32 */,
      repeated: true,
    },
    {
      no: 2,
      name: "max_open_reviews",
      kind: "scalar",
      T: 13 /* ScalarType.UINT32 */,
      opt: true,
    },
  ],
);

/**
 * A request for FrontendService.CreateUser.
 *
 * @generated from message frontendapi.CreateUserRequest
 */
export const CreateUserRequest = /*@__PURE__*/ proto3.makeMessageType(
  "frontendapi.CreateUserRequest",
  () => [{ no: 1, name: "user", kind: "message", T: User, opt: true }],
);

/**
 * A response for FrontendService.CreateUser.
 *
 * Empty to allow future extension.
 *
 * @generated from message frontendapi.CreateUserResponse
 */
export const CreateUserResponse = /*@__PURE__*/ proto3.makeMessageType(
  "frontendapi.CreateUserResponse",
  [],
);
