package main

import (
	"context"
	"database/sql"

	"rpc"

	"github.com/golang/protobuf/proto"
	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	errBadContext = runtime.NewError("bad context", 3)
	errMarshal    = runtime.NewError("cannot marshal response", 13)
	errUnmarshal  = runtime.NewError("cannot unmarshal request", 13)
)

type sessionContext struct {
	UserID    string
	SessionID string
}

func unpackContext(ctx context.Context) (*sessionContext, error) {
	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		return nil, errBadContext
	}
	sessionID, ok := ctx.Value(runtime.RUNTIME_CTX_SESSION_ID).(string)
	if !ok {
		return nil, errBadContext
	}
	return &sessionContext{UserID: userID, SessionID: sessionID}, nil
}
func rpcCreateMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	// logger.Info("received payload %v", payload)
	var session *sessionContext
	var err error
	if session, err = unpackContext(ctx); err != nil {
		logger.Error("unpack context failed: %v", err)
		return "", errBadContext
	}

	var request rpc.CreateMatchRequest
	if err = proto.Unmarshal([]byte(payload), &request); err != nil {
		logger.Error("unmarshalling failed: %v", err)
		return "", errUnmarshal
	}
	// logger.Info("session: %v, req %v", session, request)
	// Create the party match.
	matchID, err := nk.MatchCreate(ctx, "Niwrad", map[string]interface{}{
		"creator": session.UserID,
		// "type_id": request.MatchType,
		// "config":  request.Configuration,
	})

	if err != nil {
		logger.Error("match created failed: %v", err)
		response := &rpc.CreateMatchResponse{
			Result: rpc.CreateMatchCompletionResult_createMatchCompletionResultUnknownInternalFailure,
		}
		responseBytes, err := proto.Marshal(response)
		if err != nil {
			return "", errMarshal
		}
		return string(responseBytes), nil
	}

	// Return result to user.
	response := &rpc.CreateMatchResponse{
		MatchId: matchID,
		Result:  rpc.CreateMatchCompletionResult_createMatchCompletionResultSucceeded,
	}
	responseBytes, err := proto.Marshal(response)
	if err != nil {
		return "", errMarshal
	}
	// var err error
	// var payloadJson map[string]string
	// json.Unmarshal([]byte(payload), &payloadJson)

	// // Create the party match.
	// matchID, err := nk.MatchCreate(ctx, "Niwrad", map[string]interface{}{})

	// response, err := proto.Marshal(&rpc.CreateMatchResponse{
	// 	MatchId: matchID,
	// 	Result:  rpc.CreateMatchCompletionResult_createMatchCompletionResultSucceeded,
	// })

	// if err != nil {
	// 	return "", ErrJsonMarshal
	// }
	return string(responseBytes), nil
}
