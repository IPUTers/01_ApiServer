package response

import "try-standard-layout/internal/domain"

type TestApi struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ToTestApis(products []*domain.TestApi) []*TestApi {
	var res []*TestApi
	for _, p := range products {
		res = append(res, ToTestApi(p))
	}
	return res
}

func ToTestApi(testApi *domain.TestApi) *TestApi {
	return &TestApi{
		ID:   testApi.ID,
		Name: testApi.Name,
		Age:  testApi.Age,
	}
}
