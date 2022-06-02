package repository

import (
	"blockchain/utils"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(post *utils.User, blockchainAddress string) (*utils.User, error)
	FindAll() ([]utils.User, error)
}

type repo struct{}

//NewPostRepository creates a new repo
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "pragmatic-reviews-e8660"
	collectionName string = "users"
)

func (*repo) Save(post *utils.User, blockchainAddress string) (*utils.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"name":               post.Name,
		"email":              post.Email,
		"password":           post.Password,
		"company":            post.Company,
		"reputation":         post.Reputation,
		"field":              post.Field,
		"optional":           post.Optional,
		"blockchain_address": blockchainAddress,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]utils.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []utils.User
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := utils.User{
			Name:       doc.Data()["name"].(string),
			Email:      doc.Data()["email"].(string),
			Password:   doc.Data()["password"].(string),
			Reputation: doc.Data()["reputation"].(string),
			Company:    doc.Data()["Title"].(bool),
			Field:      doc.Data()["Title"].(string),
			Optional:   doc.Data()["Title"].(map[string]string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
