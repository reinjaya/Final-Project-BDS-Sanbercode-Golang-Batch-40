package structs

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Article struct {
	ID      int64  `json:"id"`
	IDUser  int64  `json:"id_user"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Content string `json:"content"`
}

type Comment struct {
	ID        int64  `json:"id"`
	IDUser    int64  `json:"id_user"`
	IDArticle int64  `json:"id_article"`
	IDReply   int64  `json:"id_reply_comment"`
	Comment   string `json:"comment"`
	Image     string `json:"image"`
}

type Like struct {
	ID        int64  `json:"id"`
	IDUser    int64  `json:"id_user"`
	IDArticle int64  `json:"id_article"`
	Respon    string `json:"respon"`
}

type LikeComment struct {
	ID        int64  `json:"id"`
	IDUser    int64  `json:"id_user"`
	IDComment int64  `json:"id_comment"`
	Respon    string `json:"respon"`
}
