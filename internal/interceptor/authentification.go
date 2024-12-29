package interceptor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	jwtutils "github.com/BelyaevEI/playlist/internal/jwt_utils"
	"github.com/BelyaevEI/playlist/internal/logger"
	"github.com/BelyaevEI/playlist/internal/storage/postgre"
	"github.com/BelyaevEI/playlist/pkg/playlist_v1"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor interface {
	Authentification() grpc.UnaryServerInterceptor
}

type authInterceptor struct {
	db *sql.DB
}

func NewAuthInterceptor(ctx context.Context, dsn string) AuthInterceptor {
	pg, err := postgre.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed connet to database: %s", err.Error())
	}

	err = pg.Ping()
	if err != nil {
		log.Fatalf("ping error: %s", err.Error())
	}

	closer.Add(pg.Close)

	return &authInterceptor{db: pg}
}

func (ai *authInterceptor) Authentification() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		// check jwt for playlist methods
		protectedMethods := map[string]bool{
			"/PlaylistV1/AddSong":    true,
			"/PlaylistV1/DeleteSong": true,
			"/PlaylistV1/Play":       true,
			"/PlaylistV1/Pause":      true,
			"/PlaylistV1/Prev":       true,
			"/PlaylistV1/Next":       true,
		}

		if protectedMethods[info.FullMethod] {

			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
			}

			authHeader := md["authorization"]
			if len(authHeader) == 0 {
				return nil, status.Errorf(codes.Unauthenticated, "authorization token is required")
			}

			token := authHeader[0]
			if len(token) > 7 && token[:7] == "Bearer " {
				token = token[7:]
			} else {
				return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
			}

			// Check JWT for other method
			switch info.FullMethod {
			case "/PlaylistV1/AddSong":
				if addSongReq, ok := req.(*playlist_v1.AddSongRequest); ok {
					newToken, err := ai.checkJWT(ctx, token, addSongReq.GetLogin())
					if err != nil {
						logger.Error(fmt.Sprintf("check JWT is failed %s", err.Error()))
						return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
					}

					grpc.SetTrailer(ctx, metadata.Pairs("new-authorization", newToken))
				}
			case "/PlaylistV1/DeleteSong":
				if delete, ok := req.(*playlist_v1.AddSongRequest); ok {
					newToken, err := ai.checkJWT(ctx, token, delete.GetLogin())
					if err != nil {
						logger.Error(fmt.Sprintf("check JWT is failed %s", err.Error()))
						return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
					}

					grpc.SetTrailer(ctx, metadata.Pairs("new-authorization", newToken))
				}
			case "/PlaylistV1/Play":
				if play, ok := req.(*playlist_v1.Request); ok {
					newToken, err := ai.checkJWT(ctx, token, play.GetLogin())
					if err != nil {
						logger.Error(fmt.Sprintf("check JWT is failed %s", err.Error()))
						return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
					}

					grpc.SetTrailer(ctx, metadata.Pairs("new-authorization", newToken))
				}
			case "/PlaylistV1/Pause":
				if pause, ok := req.(*playlist_v1.Request); ok {
					newToken, err := ai.checkJWT(ctx, token, pause.GetLogin())
					if err != nil {
						logger.Error(fmt.Sprintf("check JWT is failed %s", err.Error()))
						return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
					}

					grpc.SetTrailer(ctx, metadata.Pairs("new-authorization", newToken))
				}
			case "/PlaylistV1/Prev":
				if prev, ok := req.(*playlist_v1.Request); ok {
					newToken, err := ai.checkJWT(ctx, token, prev.GetLogin())
					if err != nil {
						logger.Error(fmt.Sprintf("check JWT is failed %s", err.Error()))
						return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
					}

					grpc.SetTrailer(ctx, metadata.Pairs("new-authorization", newToken))
				}
			case "/PlaylistV1/Next":
				if next, ok := req.(*playlist_v1.Request); ok {
					newToken, err := ai.checkJWT(ctx, token, next.GetLogin())
					if err != nil {
						logger.Error(fmt.Sprintf("check JWT is failed %s", err.Error()))
						return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
					}

					grpc.SetTrailer(ctx, metadata.Pairs("new-authorization", newToken))
				}
			}

		}

		return handler(ctx, req)
	}
}

// Check input JWT and if token success then generate new token
func (ai *authInterceptor) checkJWT(ctx context.Context, token, login string) (string, error) {

	if len(token) == 0 || len(login) == 0 {
		logger.Info("empty input data")
		return "", errors.New("invalid input data")
	}

	secretKey, err := ai.getUserSecretKey(ctx, login)
	if err != nil {
		logger.Info(fmt.Sprintf("getting user secret key is failed %s", err.Error()))
		return "", err
	}

	ok, err := jwtutils.ValidateToken(token, secretKey)
	if err != nil {
		logger.Info(err.Error())
		return "", err
	}

	if !ok {
		logger.Info("token is expired")
		return "", errors.New("token is expired")
	}

	newToken, err := jwtutils.GenerateToken(login, secretKey)
	if err != nil {
		logger.Debug(err.Error())
		return "", err
	}

	return newToken, nil
}

func (ai *authInterceptor) getUserSecretKey(ctx context.Context, login string) (string, error) {
	var secretKey string

	query := `
		SELECT secret_word FROM users WHERE user_login = $1
		`

	row := ai.db.QueryRowContext(ctx, query, login)
	if err := row.Scan(&secretKey); err != nil {
		return "", err
	}

	return secretKey, nil
}
