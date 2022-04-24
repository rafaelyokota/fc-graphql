package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/rafaelyokota/fc-graphql/graph/generated"
	"github.com/rafaelyokota/fc-graphql/graph/model"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, input *model.NewCategory) (*model.Category, error) {
	category := model.Category{
		ID:          fmt.Sprintf("T%d", rand.Intn(200)),
		Name:        input.Name,
		Description: &input.Description,
	}
	r.Categories = append(r.Categories, &category)
	return &category, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, input *model.NewCourse) (*model.Course, error) {
	var category *model.Category

	for _, value := range r.Categories {
		if value.ID == input.CategoryID {
			category = value
		}
	}

	course := model.Course{
		ID:          fmt.Sprintf("T%d", rand.Intn(200)),
		Name:        input.Name,
		Description: &input.Description,
		Category:    category,
	}

	r.Courses = append(r.Courses, &course)
	return &course, nil
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input *model.NewChapter) (*model.Chapter, error) {
	var course *model.Course
	for _, value := range r.Courses {
		if value.ID == input.CourseID {
			course = value
		}
	}
	chapter := &model.Chapter{
		ID:     fmt.Sprintf("T%d", rand.Intn(200)),
		Name:   &input.Name,
		Course: course,
	}

	r.Chapters = append(r.Chapters, chapter)

	return chapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
// func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
// 	panic(fmt.Errorf("not implemented"))
// }
// func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
// 	panic(fmt.Errorf("not implemented"))
// }
