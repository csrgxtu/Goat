package models

type (
  Recs interface {}

	Result struct {
		Msg   string `json:"msg" bson:"msg"`
		Data  []Recs `json:"data" bson:"data"`
	}

  // 保持和瑞珊同学的接口返回一致
  NewResult struct {
    Msg string  `json:"message" bson:"message"`
    Status int `json:"status" bson:"status"`
    Success bool `json:"success" bson:"success"`
    Data Recs `json:"data" bson:"data"`
  }
  NewResults struct {
    Msg string  `json:"message" bson:"message"`
    Status int `json:"status" bson:"status"`
    Success bool `json:"success" bson:"success"`
    Data []Recs `json:"data" bson:"data"`
  }
)
