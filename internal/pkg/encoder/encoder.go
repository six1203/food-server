package encoder

import (
	stdhttp "net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/tidwall/sjson"
)

type httpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Ts      string      `json:"ts"`
	Reason  string      `json:"reason"`
}

func contentType(subtype string) string {
	return strings.Join([]string{"application", subtype}, "/")
}

func ErrorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	se := errors.FromError(err)
	reply := new(httpResponse)
	reply.Code = int(se.Code)
	reply.Data = nil
	reply.Message = se.Message
	reply.Reason = se.Reason
	reply.Ts = time.Now().Format("2006-01-02 15:04:05.00000")
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(reply)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType(codec.Name()))
	if se.Code > 99 && se.Code < 600 {
		w.WriteHeader(stdhttp.StatusOK)
	} else {
		w.WriteHeader(stdhttp.StatusInternalServerError)
	}
	w.Write(body)
}

func ResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, v interface{}) error {
	reply := new(httpResponse)
	reply.Code = stdhttp.StatusOK
	reply.Data = v
	reply.Message = "success"
	reply.Reason = "success"
	reply.Ts = time.Now().Format("2006-01-02 15:04:05.00000")
	codec, _ := http.CodecForRequest(r, "Accept")
	// 解决零值字段因为自定义ResponseEncoder导致json序列化被忽略掉的问题
	// https://github.com/go-kratos/kratos/issues/1952
	replyBytes, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	jsonRes := string(replyBytes)

	resBytes, err := codec.Marshal(v)
	if err != nil {
		return err
	}
	raw, err := sjson.SetRaw(jsonRes, "data", string(resBytes))
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", contentType(codec.Name()))
	w.WriteHeader(stdhttp.StatusOK)
	w.Write([]byte(raw))
	return nil
}
