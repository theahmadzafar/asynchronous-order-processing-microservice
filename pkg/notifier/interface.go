package notifier

type Notifier interface {
	Email() error
}
