package sources

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type RedditClient struct {
	client *reddit.Client
}

func NewRedditClient() (*RedditClient, error) {
	credentials := reddit.Credentials{
		ID: "id", Secret: "secret", Username: "username", Password: "password",
	}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, err
	}
	return &RedditClient{client: client}, nil
}

func (r RedditClient) GetTrendingPosts(topic string, limit int) ([]Post, error) {
	posts, _, err := r.client.Subreddit.TopPosts(context.Background(), topic, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{Limit: limit},
		Time:        "day",
	})
	if err != nil {
		return nil, err
	}

	var result []Post
	for _, p := range posts {
		result = append(result, Post{
			Title:    p.Title,
			Body:     p.Body,
			Username: p.Author,
			Upvotes:  p.Score,
		})
	}
	return result, nil
}

func (r RedditClient) GetTrendingComments(postID string, limit int) ([]Comment, error) {
	postAndComments, _, err := r.client.Post.Get(
		context.Background(), postID,
	)
	if err != nil {
		return nil, err
	}
	var comments []Comment
	for _, comment := range postAndComments.Comments {
		comments = append(comments, Comment{
			Body:     comment.Body,
			Username: comment.Author,
			Upvotes:  comment.Score,
		})
		if len(comments) >= limit {
			break
		}
	}
	return comments, nil
}
