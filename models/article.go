package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	Article{ID: 0, Title: "Title0", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec venenatis mi rutrum lacinia rutrum. Nunc rutrum, est at mollis facilisis, ligula ligula tempor metus, sed venenatis risus lorem ac urna. Cras pretium lectus eget urna tristique auctor. Morbi vehicula orci id interdum molestie. Mauris non augue felis. Curabitur a risus in quam sollicitudin rutrum nec sit amet lacus. Vivamus lobortis fringilla tellus, eleifend viverra augue tincidunt sit amet. Vivamus posuere dui sit amet convallis fermentum. Nam rutrum ex interdum eros molestie, non rutrum libero facilisis. Fusce vel turpis euismod lacus molestie congue id a mi."},
	Article{ID: 1, Title: "Title1", Content: "Integer dictum dolor ligula. Vivamus egestas quam leo, id sagittis mauris posuere ut. Mauris hendrerit magna nec ultrices auctor. Vivamus ultrices pharetra egestas. Morbi vestibulum quam sit amet elit eleifend fermentum. Phasellus at venenatis felis. In lorem neque, pharetra dapibus hendrerit non, facilisis sed ligula. Cras dictum mi sed sodales placerat. Sed blandit nulla lorem. Duis pellentesque gravida sapien at euismod. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos."},
}

func GetAllArticles() (articles []Article) {
	return articleList
}

func GetArticleById(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func CreateArticle(Title, Content string) {
	articleList = append(articleList, Article{
		ID:      articleList[len(articleList)-1].ID + 1,
		Title:   Title,
		Content: Content,
	})
}
