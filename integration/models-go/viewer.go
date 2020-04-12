package models

import "git.sr.ht/~sircmpwn/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
