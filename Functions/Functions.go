package Functions

import (
	h "ProductService/Helpers"
)

func StartFunctions() bool {
	functions := []string{CreateUserTable(), CreateCodeTable(), CreateUuidEx(), SignIn(), SignUp(), IsVerifiedAccount(), SetCode(), VerifyAccount(), CreateProductTable(), Createproduct(), Deleteproduct(), Updateproduct(), InsertFakeData()}
	for i := 0; i < len(functions); i++ {
		statu := h.RunQuery(functions[i])
		if !statu {
			return false
		}
	}
	return true
}
