package domain

//Stub Mock Adapter
type ClientRepositoryStub struct {
	clients []Client
}

func (s ClientRepositoryStub) FindAll() ([]Client, error) {
	return s.clients, nil
}

//Create dummy clients
func NewClientRepositoryStub() ClientRepositoryStub {
	clients := []Client{
		{"1", "Pierre", "Lyon", "69007", "1990-04-03", "1"},
		{"2", "Samuel", "Paris", "91000", "1993-11-01", "1"},
	}
	return ClientRepositoryStub{clients}
}
