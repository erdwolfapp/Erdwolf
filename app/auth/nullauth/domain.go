package nullauth

type NullAuthDomain struct {}

func (f *NullAuthDomain) Name() string {
	return "without validation"
}

func (f *NullAuthDomain) DomainId() string {
	return "null"
}