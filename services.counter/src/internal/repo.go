package internal

type Repo struct {
	count int
}

type RepoParams struct {}

func NewRepo(params *RepoParams) *Repo {
	return &Repo{
		count: 0,
	}
}

func (r *Repo) Get() int {
	return r.count
}

func (r *Repo) Inc() {
	r.count++
}

func (r *Repo) Dec() {
	r.count--
}

func (r *Repo) Reset() {
	r.count = 0
}

func (r *Repo) Set(count int) {
	r.count = count
}