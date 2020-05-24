package inventory

//Service ...
type Service struct {
	repo Repo
}

//ProvideInventoryService ... service to talk to the repo
func ProvideInventoryService(r Repo) Service {
	return Service{repo: r}
}

//GetAll ... gets all the the records matching the where condition
func (s *Service) GetAll() {
	// return s.repo.GetClient()
	s.repo.GetClient()
}

//Ping ... Ping to check connectivity
func (s *Service) Ping() {
	s.repo.Ping()
}

//Add Item ... Ping to check connectivity
func (s *Service) Add(item Item) {
	s.repo.Add(item)
}

//Find Item ... Ping to check connectivity
func (s *Service) Find(ID string) []*Item {
	return s.repo.Find(ID)
}
