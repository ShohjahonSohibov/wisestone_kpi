package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	RoleId *LoginResponseRole `json:"role_id"`
}

type LoginResponseRole struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	NameUz        string    `bson:"name_uz" json:"name_uz" binding:"required"`
	NameEn        string    `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string    `bson:"name_kr" json:"name_kr" binding:"required"`
}
