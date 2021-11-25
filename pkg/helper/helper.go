package helper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"time"

	"bitbucket.org/udevs/ur_go_user_service/pkg/logger"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/runtime/protoiface"
)

func HandleError(log logger.Logger, err error, message string, req interface{}, code codes.Code) error {
	if code == codes.Canceled {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(code, message)
	} else if err == sql.ErrNoRows {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message)
	}
	return nil
}

func ParseTime(timeString string) (time.Time, error) {
	resp, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return time.Time{}, err
	}

	return resp, err
}

func ParseToStruct(data interface{}, m protoiface.MessageV1) error {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	js, err := jspbMarshal.MarshalToString(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(js), data)
	return err
}

func ParseToProto(m protoiface.MessageV1, data interface{}) error {
	var jspbUnmarshal jsonpb.Unmarshaler
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = jspbUnmarshal.Unmarshal(bytes.NewBuffer(js), m)

	return err
}

func ParseToJson(object interface{}, data interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, object)
	return err
}
