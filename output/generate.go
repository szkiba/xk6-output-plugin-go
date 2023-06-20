// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package output

//go:generate protoc -I .. --go_opt=module=github.com/szkiba/xk6-output-plugin-go/output --go-grpc_opt=module=github.com/szkiba/xk6-output-plugin-go/output --go-grpc_out=. --go_out=.  ../proto/output.proto
