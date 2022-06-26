package middlewares

import (
	"blog/app/infrastractures"
	"blog/app/interfaces"
	"blog/app/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var URL = os.Getenv("UserAddress") + "/users/detail"

type VipMiddleware struct {
	logger interfaces.Logger
}

func NewVipMiddleware(logger infrastractures.PasargadLogger) VipMiddleware {
	return VipMiddleware{
		logger: &logger,
	}
}

func (m VipMiddleware) Handle(reqFunc httptransport.DecodeRequestFunc) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var buf bytes.Buffer
		body := io.TeeReader(r.Body, &buf)

		var request models.UserIDRequest
		if err := json.NewDecoder(body).Decode(&request); err != nil {
			return nil, err
		}
		user := new(models.User)
		if err := m.getUser(request.UserID, user); err != nil {
			return nil, err
		}
		if user.VIP {
			r.Body = ioutil.NopCloser(&buf)
			return reqFunc(ctx, r)
		}
		return nil, errors.New("your level is not vip")
	}
}

func (m VipMiddleware) getUser(id int64, user *models.User) error {
	var jsonStr = []byte(fmt.Sprintf(`{"id":%d}`, id))
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		m.logger.Error("failed to send request:" + err.Error())
		return err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		m.logger.Error(err.Error())
		return err
	}

	return nil
}
