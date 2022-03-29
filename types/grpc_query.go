package types

import (
	"context"
	gocontext "context"
	"errors"
	"fmt"
	gogogrpc "github.com/gogo/protobuf/grpc"
	codecTypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	sdkerrors "github.com/oracleNetworkProtocol/plugchain-sdk-go/types/errors"
	grpctypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/types/grpc"
	"github.com/tendermint/tendermint/abci/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
)

var _ gogogrpc.ClientConn = Context{}

var protoCodec = encoding.GetCodec(proto.Name)

type Context struct {
	Client            BaseClient
	InterfaceRegistry codecTypes.InterfaceRegistry
}

// Invoke implements the grpc ClientConn.Invoke method
func (ctx Context) Invoke(grpcCtx gocontext.Context, method string, req, reply interface{}, opts ...grpc.CallOption) (err error) {
	reqBa, err := protoCodec.Marshal(req)
	if err != nil {
		return err
	}

	md, _ := metadata.FromOutgoingContext(grpcCtx)
	abciReq := types.RequestQuery{
		Data:   reqBa,
		Path:   method,
		Height: 0,
	}
	res, err := ctx.QueryABCI(abciReq)
	if err != nil {
		return err
	}
	err = protoCodec.Unmarshal(res.Value, reply)
	if err != nil {
		return err
	}
	md = metadata.Pairs(grpctypes.GRPCBlockHeightHeader, strconv.FormatInt(res.Height, 10))
	for _, callOpt := range opts {
		header, ok := callOpt.(grpc.HeaderCallOption)
		if !ok {
			continue
		}

		*header.HeaderAddr = md
	}
	if ctx.InterfaceRegistry != nil {
		return codecTypes.UnpackInterfaces(reply, ctx.InterfaceRegistry)
	}

	return nil
}

func (ctx Context) QueryABCI(req types.RequestQuery) (types.ResponseQuery, error) {
	node, err := ctx.GetNode()
	if err != nil {
		return types.ResponseQuery{}, err
	}

	var queryHeight int64
	queryHeight = req.Height

	opts := rpcclient.ABCIQueryOptions{
		Height: queryHeight,
		Prove:  req.Prove,
	}

	result, err := node.ABCIQueryWithOptions(context.Background(), req.Path, req.Data, opts)
	if err != nil {
		return types.ResponseQuery{}, err
	}

	if !result.Response.IsOK() {
		return types.ResponseQuery{}, sdkErrorToGRPCError(result.Response)
	}

	return result.Response, nil
}

func sdkErrorToGRPCError(resp types.ResponseQuery) error {
	switch resp.Code {
	case sdkerrors.ErrInvalidRequest.ABCICode():
		return status.Error(codes.InvalidArgument, resp.Log)
	case sdkerrors.ErrUnauthorized.ABCICode():
		return status.Error(codes.Unauthenticated, resp.Log)
	case sdkerrors.ErrKeyNotFound.ABCICode():
		return status.Error(codes.NotFound, resp.Log)
	default:
		return status.Error(codes.Unknown, resp.Log)
	}
}

func (ctx Context) GetNode() (BaseClient, error) {
	if ctx.Client == nil {
		return nil, errors.New("no RPC client is defined in offline mode")
	}

	return ctx.Client, nil
}

// NewStream implements the grpc ClientConn.NewStream method
func (Context) NewStream(gocontext.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("streaming rpc not supported")
}
