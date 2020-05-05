package model

type ArticleStoreInMemory struct {
	ArticleMap []Article
}

func NewArticleStoreInMemory() *ArticleStoreInMemory { //Membuat instance dari articlesinmemory
	return &ArticleStoreInMemory{
		ArticleMap: []Article{
			Article{
				ID:    1,
				Title: "Membuat Website",
				Body:  "Hallo ini Body"},
		},
	}
}

func (store *ArticleStoreInMemory) Save(article *Article) error {
	lastID := len(store.ArticleMap)
	article.ID = lastID + 1
	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}
func (store *ArticleStoreInMemory) Remove(id int) error {
	store.ArticleMap = append(store.ArticleMap[:id-1], store.ArticleMap[id:]...)
	return nil
}
func (store *ArticleStoreInMemory) EditArticle(title, body string, id int) error {
	store.ArticleMap[id-1] = Article{ID: id, Title: title, Body: body}
	return nil
}
