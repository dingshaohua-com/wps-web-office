package api

// Body 是不额外包装业务信封的 Huma 响应结构。
type Body[T any] struct {
	Body T
}

func NewBody[T any](data T) *Body[T] {
	return &Body[T]{Body: data}
}

func NoContent() *struct{} {
	return &struct{}{}
}

// Body2 ---使用 Huma 识别的 Body 字段承载统一响应结构。

// Envelope 是包含业务状态码和数据的统一响应结构。
type Envelope[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

type Body2[T any] struct {
	Body Envelope[T]
}

func NewBody2[T any](data T) *Body2[T] {
	return &Body2[T]{
		Body: Envelope[T]{
			Code: 0,
			Data: data,
		},
	}
}
