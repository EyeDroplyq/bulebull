/**
  @author: lyq
  @since: 2023-08-13
  @desc:
**/

package model

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Gender   int    `db:"gender"`
	UserID   int64  `db:"userId"`
}
