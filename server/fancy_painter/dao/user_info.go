package dao

import (
	"context"
	"log"
)

// 参考： https://cloud.tencent.com/developer/article/1493903
func GetUserInfoByMobile(ctx context.Context, mobile string) (*Users, error) {
	sql := "select id,phone,user_id  from users where phone = ? "
	rows, err := gFancyPainterDB.QueryContext(ctx, sql, mobile)
	if err != nil {
		log.Printf("GetUserInfoByMobile error, mobile:%s, sql:%s, err:%v", mobile, sql, err)
		return nil, err
	}
	defer rows.Close()

	var users []Users
	for rows.Next() {
		var user Users
		err = rows.Scan(&user.ID, &user.Phone, &user.UserId)
		if err != nil {
			log.Printf("GetUserInfoByMobile error, mobile:%s, sql:%s, err:%v", mobile, sql, err)
			continue
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

// InsertUserIDByMobile 通过手机号来插入用户id  ==>  https://blog.csdn.net/t894690230/article/details/77996355
func InsertUserIDByMobile(ctx context.Context, mobile string, user_id string) error {
	sql := "insert ignore into users (phone,user_id) values (?,?) "
	_, err := gFancyPainterDB.ExecContext(ctx, sql, mobile, user_id)
	if err != nil {
		log.Printf("InsertUserIDByMobile error, mobile:%s, user_id:%s, sql:%s, err:%v",
			mobile, user_id, sql, err)
		return err
	}
	return nil
}
