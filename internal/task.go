package internal

type Inputs struct {
}

func (i *Inputs) Get(key string) string {
	return ""
}

func (i *Inputs) Set(key, value string) {
}

type Outputs struct {
}

func (o *Outputs) Set(key, value string) {
}

type Task struct {
	Name    string
	Desc    string
	Inputs  Inputs
	Outputs Outputs
}
