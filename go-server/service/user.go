package service

import (
	"context"
	"go-server/pb"
	"go-server/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAllUser(ctx context.Context, empty *emptypb.Empty) (res *pb.UserList, err error) {
	userS, err := s.userRepo.GetAllUser(ctx)
	var resUsers []*pb.User
	for _, user := range userS {
		resUsers = append(resUsers, &pb.User{
			Name:    user.Name,
			Address: user.Address,
			Phone:   user.Phone,
			Email:   user.Email,
		})
	}
	return &pb.UserList{
		Users: resUsers,
	}, err
}

func (s *UserService) Register(ctx context.Context, user *pb.User) (*pb.User, error) {
	resUser, err := s.userRepo.Register(ctx, user.Name, user.Address, user.Phone, user.Email)
	return &pb.User{
		Name:    resUser.Name,
		Address: resUser.Address,
		Phone:   resUser.Phone,
		Email:   resUser.Email,
	}, err
}
