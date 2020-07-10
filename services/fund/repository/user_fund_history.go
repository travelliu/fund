// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	_utils "github.com/travelliu/fund/utils"
)

func (r *repo) ReplaceUserFundHistory(ctx context.Context, code string, userFundHistories []*_fundMod.UserFundHistory) error {
	tx := r.db.Begin()

	err := tx.Where("code = ?", code).Delete(&_fundMod.UserFundHistory{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, u := range userFundHistories {
		if u.ID == 0 {
			u.ID = _utils.GenerateID()
		}
		err = tx.Create(u).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
