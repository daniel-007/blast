// Copyright (c) 2018 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/mosuka/blast/node/data/client"
	"github.com/mosuka/blast/node/data/protobuf"
	"github.com/urfave/cli"
)

func join(c *cli.Context) error {
	grpcAddr := c.String("grpc-addr")
	targetRaftNodeID := c.String("target-raft-node-id")
	targetRaftAddr := c.String("target-raft-addr")
	targetGrpcAddr := c.String("target-grpc-addr")
	targetHttpAddr := c.String("target-http-addr")

	dataClient, err := client.NewGRPCClient(grpcAddr)
	if err != nil {
		return err
	}
	defer dataClient.Close()

	req := &protobuf.PutNodeRequest{
		Id:       targetRaftNodeID,
		RaftAddr: targetRaftAddr,
		GrpcAddr: targetGrpcAddr,
		HttpAddr: targetHttpAddr,
	}

	_, err = dataClient.PutNode(req)
	if err != nil {
		return err
	}

	return nil
}
