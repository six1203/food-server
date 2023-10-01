package encoder

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdhttp "net/http"
	"strings"
	"time"
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
	data, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", contentType(codec.Name()))
	w.WriteHeader(stdhttp.StatusOK)
	w.Write(data)
	return nil
}
