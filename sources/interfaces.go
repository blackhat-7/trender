package sources

type Post struct {
	ID       string
	Title    string
	Body     string
	Username string
	Upvotes  int
}

type Comment struct {
	ID       string
	Body     string
	Username string
	Upvotes  int
}

type TrendGetter interface {
	GetTrendingPosts(topic string, limit int) ([]Post, error)
	GetTrendingComments(postID string, limit int) ([]Post, error)
}
