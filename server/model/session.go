package model

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/teru01/image/server/database"
)

type Session struct {
	sess *sessions.Session
}

const sessionName = "auth"

func NewSession(user *User, context *database.DBContext) (*Session, error) {
	if err := user.ValidateLoginUser(context); err != nil {
		return nil, err
	}

	session, err := session.Get(sessionName, context)
	if err != nil {
		return nil, err
	}
	session.Options = &sessions.Options{
		MaxAge: 86400 * 7,
	}

	newSession := Session {
		sess: session,
	}

	user.SetSessionValue(&newSession)
	session.Save(context.Request(), context.Response())
	return &newSession, nil
}

func (s *Session) Set(key string, val interface{}) {
	s.sess.Values[key] = val
}
