package server

import (
	"encoding/json"
	"github.com/juricaKenda/golang-rest-example/src/server/status"
	"net/http"
)

func (s *Server) getAllArticles() func(http.ResponseWriter, *http.Request) {
	return func(out http.ResponseWriter, _ *http.Request) {
		s.log.Trace("entered the server handler")

		articles, err := s.articleRepo.GetAllArticles()
		if err != nil {
			out.WriteHeader(status.INTERNAL_ERR)
			s.log.Error("internal repo err")
			return
		}

		body, err := json.Marshal(articles)
		if err != nil {
			out.WriteHeader(status.INTERNAL_ERR)
			s.log.Error("failed unmarshalling")
			return
		}

		out.WriteHeader(status.OK)
		_, err = out.Write(body)
		if err != nil {
			s.log.Error("failed network stuff")
			return
		}
	}
}
