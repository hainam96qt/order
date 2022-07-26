package auth

import (
	"context"
	"database/sql"
	"log"
	"time"

	"order-gokomodo/configs"
	"order-gokomodo/pkg/db/mysql_db"
	"order-gokomodo/pkg/entities"
	sql_model "order-gokomodo/pkg/usecase/model"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey string
var expireTime time.Duration = 60

var _ entities.Authentication = &AuthenticationService{}

type AuthenticationService struct {
	DatabaseConn *sql.DB
	Query        *sql_model.Queries
}

func NewAuthenticationService(cfg *configs.Config) (*AuthenticationService, error) {
	jwtKey = cfg.SecretKey
	databaseConn, err := mysql_db.ConnectDatabase(cfg.Mysqldb)
	if err != nil {
		return nil, err
	}
	query := sql_model.New(databaseConn)
	return &AuthenticationService{
		DatabaseConn: databaseConn,
		Query:        query,
	}, nil
}

type Claims struct {
	Username string `json:"username"`
	USerID   int    `json:"user_id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (c Claims) Valid() error {
	return c.StandardClaims.Valid()
}

func (a AuthenticationService) Register(ctx context.Context, request *entities.RegisterRequest) (*entities.RegisterResponse, error) {
	hashPassword, err := hash(request.Password)
	if err != nil {
		return nil, err
	}

	roleID := 2
	if request.Role == "seller" {
		roleID = 1
	}
	_, err = a.Query.CreateUser(ctx, sql_model.CreateUserParams{
		Email:      request.Email,
		Password:   hashPassword,
		Name:       request.Name,
		Address:    request.Address,
		UserRoleID: int32(roleID),
	})
	if err != nil {
		return nil, err
	}

	loginInfo, err := a.Login(ctx, &entities.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &entities.RegisterResponse{Token: loginInfo.Token}, nil
}

func (a AuthenticationService) Login(ctx context.Context, request *entities.LoginRequest) (*entities.LoginResponse, error) {
	// TODO validate user
	// get user from database
	user, err := a.Query.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	err = CheckPasswordHash(user.Password, request.Password)
	if err != nil {
		return nil, err
	}
	role, err := a.Query.GetRoleByID(ctx, user.UserRoleID)
	if err != nil {
		return nil, err
	}
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: user.Name,
		USerID:   int(user.ID),
		Role:     role.RoleName.String,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}
	log.Print(tokenString)
	return &entities.LoginResponse{Token: tokenString}, nil
}

func DecodeToken(token string) (*Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
