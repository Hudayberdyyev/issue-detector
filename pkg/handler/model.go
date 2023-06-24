package handler

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"issue-detector/pkg/repository"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type RefreshTokenBody struct {
	RequestFingerprint   string `json:"request_fingerprint"`
	RequestRefreshToken  string `json:"request_refresh_token"`
	ResponseRefreshToken string `json:"response_refresh_token"`
	ResponseAccessToken  string `json:"response_access_token"`
	ResponseStatusCode   int    `json:"response_status_code"`
	ResponseMessage      string `json:"response_message"`
	ErrorCode            string `json:"error_code"`
}

func (r RefreshTokenBody) GetRefreshTokenDbModel() repository.RefreshTokenModel {
	return repository.RefreshTokenModel{
		UserId:               r.GetUserIdFromAccessToken(),
		RequestFingerprint:   r.RequestFingerprint,
		RequestRefreshToken:  r.RequestRefreshToken,
		ResponseRefreshToken: r.ResponseRefreshToken,
		ResponseAccessToken:  r.ResponseAccessToken,
		ResponseStatusCode:   r.ResponseStatusCode,
		ResponseMessage:      r.ResponseMessage,
		ErrorCode:            r.ErrorCode,
	}
}

func (r RefreshTokenBody) GetUserIdFromAccessToken() int {
	accessToken := r.ResponseAccessToken
	claims, err := parseToken(accessToken)
	if err != nil || claims == nil {
		return InvalidUserId
	}
	return claims.UserId
}

type CheckIpBody struct {
	AuthorizationHeader     string `json:"authorization_header"`
	UserAgentHeader         string `json:"user_agent_header"`
	MacAddressHeader        string `json:"mac_address_user_header"`
	XForwardedForHeader     string `json:"x_forwarded_for_header"`
	ResponseStatusCode      int    `json:"response_status_code"`
	ResponseIsActualVersion int    `json:"response_is_actual_version"`
	ResponseIsAccessAllowed int    `json:"response_is_access_allowed"`
	ErrorLog                string `json:"error_log"`
	ErrorCode               string `json:"error_code"`
}

func (c CheckIpBody) GetCheckIpModel() repository.CheckIpModel {
	return repository.CheckIpModel{
		UserId:                  c.GetUserIdFromAccessToken(),
		AuthorizationHeader:     c.AuthorizationHeader,
		UserAgentHeader:         c.UserAgentHeader,
		MacAddressHeader:        c.MacAddressHeader,
		XForwardedForHeader:     c.XForwardedForHeader,
		ResponseStatusCode:      c.ResponseStatusCode,
		ResponseIsActualVersion: c.ResponseIsActualVersion,
		ResponseIsAccessAllowed: c.ResponseIsAccessAllowed,
		DbIsIpAccessAllowed:     0,
		DbIsUserAccessAllowed:   0,
		ErrorLog:                c.ErrorLog,
		ErrorCode:               c.ErrorCode,
	}
}

func (c CheckIpBody) GetCheckIpModelWithArgs(dbIpAccess, dbUserAccess int) repository.CheckIpModel {
	return repository.CheckIpModel{
		UserId:                  c.GetUserIdFromAccessToken(),
		AuthorizationHeader:     c.AuthorizationHeader,
		UserAgentHeader:         c.UserAgentHeader,
		MacAddressHeader:        c.MacAddressHeader,
		XForwardedForHeader:     c.XForwardedForHeader,
		ResponseStatusCode:      c.ResponseStatusCode,
		ResponseIsActualVersion: c.ResponseIsActualVersion,
		ResponseIsAccessAllowed: c.ResponseIsAccessAllowed,
		DbIsIpAccessAllowed:     dbIpAccess,
		DbIsUserAccessAllowed:   dbUserAccess,
		ErrorLog:                c.ErrorLog,
		ErrorCode:               c.ErrorCode,
	}
}

func (c CheckIpBody) GetUserIdFromAccessToken() int {
	accessToken := c.AuthorizationHeader
	claims, err := parseToken(accessToken)
	if err != nil || claims == nil {
		return InvalidUserId
	}
	return claims.UserId
}

func parseToken(accessToken string) (*TokenClaims, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignInAccess), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("bad request")
	}
	return claims, nil
}
