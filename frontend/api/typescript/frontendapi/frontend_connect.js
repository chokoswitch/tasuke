// @generated by protoc-gen-connect-es v1.4.0
// @generated from file frontendapi/frontend.proto (package frontendapi, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { CreateUserRequest, CreateUserResponse } from "./frontend_pb.js";

/**
 * The service for the frontend.
 *
 * @generated from service frontendapi.FrontendService
 */
export const FrontendService = {
  typeName: "frontendapi.FrontendService",
  methods: {
    /**
     * Creates a new user. Tasuke is a GitHub application and users strongly
     * tied to GitHub users. For example, the user ID used in the system is
     * the GitHub user ID itself.
     *
     * @generated from rpc frontendapi.FrontendService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: CreateUserRequest,
      O: CreateUserResponse,
      kind: MethodKind.Unary,
    },
  },
};
