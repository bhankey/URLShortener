package URLShortener

import (
	"URLShortener/internal/api"
	"URLShortener/internal/store"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type GRPCServer struct {
	config *Config
	store  store.Store
	logger *logrus.Logger
	api.UnimplementedURLShortenerServer
}

func NewGRPCServer(config *Config) *GRPCServer {
	return &GRPCServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start start grpc server
func (s *GRPCServer) Start() error {
	if err := s.ConfigureStore(); err != nil {
		return fmt.Errorf("could not configure storage: %v", err)
	}
	if err := s.ConfigureLogger(); err != nil {
		return fmt.Errorf("could not configure logger: %v", err)
	}
	srv := grpc.NewServer()
	api.RegisterURLShortenerServer(srv, s)

	l, err := net.Listen("tcp", s.config.Server.BindAddr)
	if err != nil {
		return fmt.Errorf("could not listen by port %v, because of %v", s.config.Server.BindAddr, err)
	}
	s.logger.Info("starting gRPC server")
	if err := srv.Serve(l); err != nil {
		return fmt.Errorf("could not start gRPC server because of %v", err)
	}
	return nil
}

// ConfigureStore configures storage of current store realisation
func (s *GRPCServer) ConfigureStore() error {
	st, err := store.New(s.config.Store)
	if err != nil {
		return err
	}
	s.store = st
	return nil
}

// ConfigureLogger configures logger
func (s *GRPCServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.Server.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	s.logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return nil
}

// Create will take a long url and generate for it short url
func (s *GRPCServer) Create(ctx context.Context, req *api.OriginalURL) (*api.ShortCode, error) {
	s.logger.Infof("Create: request started with url \"%s\"", req.Url)
	metadata.FromIncomingContext(ctx)
	if !isURL(req.Url) {
		s.logger.Warnf("Create: incorrect url was received: \"%s\"", req.Url)
		return nil, fmt.Errorf("URL is not correct")
	}
	repository := s.store.URL()
	shortURL, err := repository.Create(req.Url)
	if err != nil {
		s.logger.Warnf("Create: request failed with url \"%s\": \"%s\"", req.Url, err)
	} else {
		s.logger.Infof("Create: request finished: url \"%s\", short code: \"%s\"", req.Url, shortURL)
	}
	return &api.ShortCode{ShortURLCode: shortURL}, err
}

//Get will take in a short url and return the original url
func (s *GRPCServer) Get(ctx context.Context, req *api.ShortCode) (*api.OriginalURL, error) {
	s.logger.Infof("Get: request with code \"%s\" started", req.ShortURLCode)
	repository := s.store.URL()
	originalURL, err := repository.Get(req.ShortURLCode)
	if err != nil {
		s.logger.Warnf("Get: request failed with short code \"%s\": \"%s\"", req.ShortURLCode, err)
	} else {
		s.logger.Infof("Get: request finish: short code \"%s\", url: \"%s\"", req.ShortURLCode, originalURL)
	}
	return &api.OriginalURL{Url: originalURL}, err
}
