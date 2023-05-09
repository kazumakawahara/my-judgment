package userdm

import (
	"context"
	"errors"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/common/vo/uservo"
)

type domainService struct {
	userRepository Repository
}

func NewUserDomainService(userRepository Repository) *domainService {
	return &domainService{
		userRepository: userRepository,
	}
}

func (s *domainService) ExistsUserByNameForCreate(ctx context.Context, nameVO uservo.Name) (bool, error) {
	if _, err := s.userRepository.FetchUserIDByName(ctx, nameVO); err != nil {
		if errors.Is(err, apperr.MjUserNotFound) {
			return false, nil
		}

		return false, mjerr.Wrap(err)
	}

	return true, nil
}

func (s *domainService) ExistsUserByNameForUpdate(ctx context.Context, idVO uservo.ID, nameVO uservo.Name) (bool, error) {
	if existingIDVO, err := s.userRepository.FetchUserIDByName(ctx, nameVO); err != nil {
		if errors.Is(err, apperr.MjUserNotFound) {
			return false, nil
		}

		return false, mjerr.Wrap(err)
	} else if idVO.Equals(existingIDVO) {
		return false, nil
	}

	return true, nil
}

func (s *domainService) ExistsUserByEmailForCreate(ctx context.Context, emailVO uservo.Email) (bool, error) {
	if _, err := s.userRepository.FetchUserIDByEmail(ctx, emailVO); err != nil {
		if errors.Is(err, apperr.MjUserNotFound) {
			return false, nil
		}

		return false, mjerr.Wrap(err)
	}

	return true, nil
}

func (s *domainService) ExistsUserByEmailForUpdate(ctx context.Context, idVO uservo.ID, emailVO uservo.Email) (bool, error) {
	if existingIDVO, err := s.userRepository.FetchUserIDByEmail(ctx, emailVO); err != nil {
		if errors.Is(err, apperr.MjUserNotFound) {
			return false, nil
		}

		return false, mjerr.Wrap(err)
	} else if idVO.Equals(existingIDVO) {
		return false, nil
	}

	return true, nil
}
