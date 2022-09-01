package mail

type Driver interface {
	Send(mail Email, config map[string][]string) bool
}
