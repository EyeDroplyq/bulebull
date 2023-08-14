/**
  @author: lyq
  @since: 2023-08-13
  @desc:
**/

package model

type User struct {
	Id       int    `db:"id"`
	UserID   int64  `db:"userId"`
	Username string `db:"username"`
	Password string `db:"password"`
	Gender   int    `db:"gender"`
	Email    string `db:"email"`
}
