// Package clrpcserver implements client and server RPC methods to call Cypherlock.
package clrpcserver

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/cypherlock-pf/clv1/ratchetserver"
	"github.com/cypherlock-pf/clv1/types"
)

// RPCServer implements a Cypherlock RPC server over http(s).
type RPCServer struct {
	rpcmethods *RPCMethods
}

// NewRPCServer creates a new RPC server and starts it.
func NewRPCServer(server *ratchetserver.RatchetServer, listenAddr string) (*RPCServer, error) {
	rs := &RPCServer{
		rpcmethods: &RPCMethods{
			server: server,
		},
	}
	rs.rpcmethods.server.StartService()
	rpc.Register(rs.rpcmethods)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	go http.Serve(l, nil)
	return rs, nil
}

// RPCMethods defines the methods for a Cypherlock RPC server.
type RPCMethods struct {
	server *ratchetserver.RatchetServer
}

// GetKeys returns the current pregenerated keys.
func (rm *RPCMethods) GetKeys(params types.RPCTypeNone, reply *types.RPCTypeGetKeysResponse) error {
	reply.Keys = rm.server.GetKeys()
	return nil
}

// Decrypt the message and return it's payload. Only use over TLS.
func (rm *RPCMethods) Decrypt(params types.RPCTypeDecrypt, reply *types.RPCTypeDecryptResponse) error {
	r, err := rm.server.Decrypt(params.OracleMessage)
	if err != nil {
		return err
	}
	reply.ResponseMessage = r
	return nil
}
