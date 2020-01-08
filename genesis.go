package errors

type genesisErr struct {
	err string
}

func (e genesisErr) Error() string {
	return e.err
}

func (e genesisErr) Is(target error) bool {
	return e == target
}

func (e genesisErr) Unwrap() error {
	return nil
}