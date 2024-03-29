package web

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"goclean/adapter/repository"
	"goclean/usecase"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"net/http"
	"time"
)

const PW_SALT_SIZE = 32
const HASHPASS_LEN = 32
const HASH_ITERATION_COUNT = 4096
const DEFAULT_TOKEN_EXPIRED_MINUTE = 30 * 24 * 60 // 30 days
const RESET_TOKEN_EXPIRED_MINUTE = 30             // 30 minutes
const RESET_PASS_AUD = "resetPassAud"

type AuthCtrl interface {
	LoginByEmail(w http.ResponseWriter, r *http.Request)
	RegisterByMail(w http.ResponseWriter, r *http.Request)
	RequestResetPassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}

// Interface for Infrastructure to implement the Jwt support module
type JwtAuth interface {
	CreateToken(uid, aud string, nExpiredDay int, signedKey string, now time.Time) (string, error)
	ParseToken(encryptedToken string, repoSignedKeyFunc func(uid, aud string) (string, error)) (string, error)
}

func NewAuthCtrl(
	userUseCase usecase.UserUseCase,
	authRepo repository.AuthRepo,
	jwtAuth JwtAuth,
	mailService usecase.MailService,
) AuthCtrl {
	return &authCtrlImpl{
		userUseCase: userUseCase,
		authRepo:    authRepo,
		jwtAuth:     jwtAuth,
		mailService: mailService,
	}
}

type authCtrlImpl struct {
	userUseCase usecase.UserUseCase
	authRepo    repository.AuthRepo
	jwtAuth     JwtAuth
	mailService usecase.MailService
}

type registerByMailReq struct {
	email string `json:"email"`
	pass  string `json:"pass"`
}

type loginByEmailReq struct {
	email string `json:"email"`
	pass  string `json:"pass"`
	aud   string `json:"aud"`
}

func (c *authCtrlImpl) RegisterByMail(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}
	req := registerByMailReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}

	// Validate data
	// TODO: validate email format

	// TODO: password should be hashed before sending to server

	// Check if auth exist
	auth, err := c.authRepo.GetByEmail(req.email)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if auth != nil {
		ResponseError(w, http.StatusBadRequest, errors.New("Email is used"))
		return
	}

	salt := GenSalt()
	hashedPass, err := HashPass(req.pass, salt)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Create a user
	userId, err := c.userUseCase.CreateUser()
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Create an auth data
	_, err = c.authRepo.CreateAuthByEmailAndHashPass(userId, req.email, hashedPass, salt)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: return body response here
	ResponseOk(w, "")
}

func (c *authCtrlImpl) LoginByEmail(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}
	req := loginByEmailReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}

	// TODO: validation
	aud := req.aud
	if aud == "" {
		aud = "defaultAud"
	}

	auth, err := c.authRepo.GetByEmail(req.email)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if auth == nil {
		ResponseError(w, http.StatusNotFound, errors.New("No Auth record"))
		return
	}

	// Validate pass
	if ValidatePass(req.pass, auth.HashedPass, auth.Salt) {
		ResponseError(w, http.StatusBadRequest, errors.New("Incorrect username & pass"))
		return
	}

	// Gen token and response
	signedKey := GenSalt()

	token, err := c.jwtAuth.CreateToken(auth.Uid, aud, DEFAULT_TOKEN_EXPIRED_MINUTE, signedKey, time.Now())
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Save signed key
	err = c.authRepo.SaveSignedKey(auth.Uid, aud, signedKey)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: Return token
	ResponseOk(w, token)
}

func (c *authCtrlImpl) RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}
	req := struct {
		email string `json:"email"`
	}{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}

	// Get auth
	auth, err := c.authRepo.GetByEmail(req.email)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if auth == nil {
		// Response ok even user is not exist to prevent privacy leak
		ResponseOk(w, "")
		return
	}

	// Gen token and response
	signedKey := GenSalt()
	resetToken, err := c.jwtAuth.CreateToken(auth.Uid, RESET_PASS_AUD, RESET_TOKEN_EXPIRED_MINUTE, signedKey, time.Now())
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Save signed key to database
	err = c.authRepo.SaveSignedKey(auth.Uid, RESET_PASS_AUD, signedKey)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: send mail to user
	c.mailService.SendMail(resetToken, auth.Uid)

	// TODO: response data later
	ResponseOk(w, "")
}

func (c *authCtrlImpl) ResetPassword(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}
	req := struct {
		pass       string `json:"pass"`
		resetToken string `json:"resetToken"`
	}{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}

	uid, err := c.jwtAuth.ParseToken(req.resetToken, func(uid, aud string) (string, error) {
		auth, err := c.authRepo.Get(uid)
		if err != nil {
			return "", err
		}
		// Only accept reset password aud
		signedKey, ok := auth.SignedKeys[RESET_PASS_AUD]
		if !ok {
			return "", nil
		}

		return signedKey.Key, nil
	})
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}

	// Get auth
	auth, err := c.authRepo.Get(uid)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	// Hash new pass
	auth.HashedPass, err = HashPass(req.pass, auth.Salt)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	// Update
	err = c.authRepo.Update(*auth)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: response data
	ResponseOk(w, "")
}

// Create a salt with size equal PW_SALT_SIZE
// salt is response as a hex number
func GenSalt() string {
	b := make([]byte, PW_SALT_SIZE)
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Errorf("Random generate salt failed: %v", err))
	}

	return base64.StdEncoding.EncodeToString(b)
}

// TODO: use byte array instead of string for pass to avoid security if memory leak
func HashPass(plainPass string, saltStr string) (string, error) {
	salt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		return "", err
	}
	hashedPass := pbkdf2.Key([]byte(plainPass), salt, HASH_ITERATION_COUNT, HASHPASS_LEN, sha3.New256)

	return base64.StdEncoding.EncodeToString(hashedPass), nil
}

func ValidatePass(pass string, hashedPass string, salt string) bool {
	encoded, err := HashPass(pass, salt)
	if err != nil {
		return false
	}
	return encoded == hashedPass
}
