package notifier

type NotifierImplementation struct {
}

func New() Notifier {
	n := &NotifierImplementation{}
	return n
}

func (n NotifierImplementation) Email() error {

	return nil
}
