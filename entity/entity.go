package entity

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/30 12:37 下午
 */

type (
	// 数据库实体映射
	Rover struct {
		Id   int
		Name string
		Pass string
		Dog  Dog
	}

	User struct {
		UserId string `json:"user_id"`
		Status int64  `json:"status"`
	}

	Dog struct {
		No   int64
		Name string
	}
)
