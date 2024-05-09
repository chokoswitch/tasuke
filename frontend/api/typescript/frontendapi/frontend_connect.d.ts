// @generated by protoc-gen-connect-es v1.4.0
// @generated from file frontendapi/frontend.proto (package frontendapi, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { SaveUserRequest, SaveUserResponse } from "./frontend_pb.js";
import type { MethodKind } from "@bufbuild/protobuf";

/**
 * The service for the frontend.
 *
 * @generated from service frontendapi.FrontendService
 */
export declare const FrontendService: {
  readonly typeName: "frontendapi.FrontendService";
  readonly methods: {
    /**
     * Saves information for a user. This method works both for a new or existing user.
     * The user is identified by the firebase ID token included in the authorization header.
     *
     * @generated from rpc frontendapi.FrontendService.SaveUser
     */
    readonly saveUser: {
      readonly name: "SaveUser";
      readonly I: typeof SaveUserRequest;
      readonly O: typeof SaveUserResponse;
      readonly kind: MethodKind.Unary;
    };
  };
};
