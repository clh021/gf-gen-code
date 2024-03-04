package model

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1"      v:"min:0#分页号码错误"              dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"999999" v:"max:999999#分页数量最大999999条" dc:"分页数量，最大50"`
}

type CommonPaginationRes struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}

type CommonMap struct {
  Key string `json:"key"`
	Value string `json:"value"`
}