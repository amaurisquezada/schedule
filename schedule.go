package schedule

import "errors"

type Client struct {
	Name string
	Age  int
}

func NewClient(name string, age int) (*Client, error) {
	if name == "" {
		return nil, errors.New("invalid blank name")
	}
	if age == 0 {
		return nil, errors.New("invalid zero age")
	}

	return &Client{name, age}, nil
}
