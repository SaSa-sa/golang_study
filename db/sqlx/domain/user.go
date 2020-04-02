package domain

// User is a struct for user table.
type User struct { //※ユーザーテーブルのどのカラムが、structのどのカラムに値するか
	// sqlxのdbタグを使ってカラムとフィールドをマッピングする
	ID        int    `db:"id" json:"id"` //...みたいにjsonの値も定義できる？(jbuilder的な) addほにゃららってコマンド叩くらしい
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

// Users have some users.
type Users []User
