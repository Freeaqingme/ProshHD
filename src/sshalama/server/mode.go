// Sshalama
//
// Copyright 2016-2017 Dolf Schimmel
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import "golang.org/x/crypto/ssh"

type mode interface {
	getPasswordCallback(dstClient **ssh.Client) func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error)
}

func (s *Server) getModeForUser(u string) (mode, error) {
	return &localMode{}, nil // TODO
	//return &proxyMode{}, nil // TODO
}
