package internal

type Srv struct {
	r *Repo
}

type SrvParams struct {
	Repo *Repo
}

func NewSrv(params *SrvParams) *Srv {
	return &Srv{
		r: params.Repo,
	}
}

func (s *Srv) Get() int {
	return s.r.Get()
}

func (s *Srv) Inc() {
	s.r.Inc()
}

func (s *Srv) Dec() {
	s.r.Dec()
}

func (s *Srv) Reset() {
	s.r.Reset()
}

func (s *Srv) Set(count int) {
	s.r.Set(count)
}