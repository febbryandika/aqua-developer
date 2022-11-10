package response

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateProductResponse struct {
	Product_Name string `json:"product_name"`
	Picture      string `json:"picture"`
	Price        int64  `json:"price"`
	Category     string `json:"category"`
	Description  string `json:"description"`
}

type GetAllProductResponse struct {
	Product_Name string `json:"product_name"`
	Picture      string `json:"picture"`
	Price        int64  `json:"price"`
	Category     string `json:"category"`
}

type GetDetailedProductResponse struct {
	Product_Name string `json:"product_name"`
	Picture      string `json:"picture"`
	Price        int64  `json:"price"`
	Description  string `json:"description"`
}

type AddProductToCartResponse struct {
	User_ID    int64 `json:"user_id"`
	Product_ID int64 `json:"product_id"`
	Quantity   int64 `json:"quantity"`
}

type GetCartProductResponse struct {
	Product  GetAllProductResponse `json:"product"`
	Quantity int64                 `json:"quantity"`
}

type CheckoutCartResponse struct {
	Checkout_ID int64 `json:"checkout_id"`
	Product_ID  int64 `json:"product_id"`
	Quantity    int64 `json:"quantity"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type LoginUserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}
