package model

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/teru01/image/server/database"
)

type Session struct {
	sess *sessions.Session
}

const SessionName = "auth"

func NewSession(user *User, context *database.DBContext) (*Session, error) {
	if err := user.ValidateLoginUser(context); err != nil {
		return nil, err
	}
	return IssueSession(user, context)
}

func IssueSession(user *User, context *database.DBContext) (*Session, error) {
	session, err := session.Get(SessionName, context)
	if err != nil {
		return nil, err
	}
	session.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	newSession := Session{
		sess: session,
	}

	user.SetSessionValue(&newSession)
	session.Save(context.Request(), context.Response())
	return &newSession, nil
}

func (s *Session) Set(key string, val interface{}) {
	s.sess.Values[key] = val
}

func GetAuthSessionData(c *database.DBContext, key string) interface{} {
	sess, err := session.Get(SessionName, c)
	if err != nil {
		return err
	}
	return sess.Values[key]
}
