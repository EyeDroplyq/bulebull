/**
  @author: lyq
  @since: 2023-08-13
  @desc:
**/

package model

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Gender     int    `json:"gender" binding:"min=0,max=1"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}
