// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package pb

import "storj.io/storj/pkg/storj"

// Path represents a object path
type Path = storj.Path

// NodeID is an alias to storj.NodeID for use in generated protobuf code
type NodeID = storj.NodeID

// NodeIDList is an alias to storj.NodeIDList for use in generated protobuf code
type NodeIDList = storj.NodeIDList

//go:generate protoc -I. --gogo_out=plugins=grpc:. meta.proto overlay.proto pointerdb.proto piecestore.proto bandwidth.proto inspector.proto datarepair.proto node.proto
