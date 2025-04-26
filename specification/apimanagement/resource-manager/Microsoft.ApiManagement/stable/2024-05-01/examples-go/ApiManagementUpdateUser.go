package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateUser.json
func ExampleUserClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().Update(ctx, "rg1", "apimService1", "5931a75ae4bbd512a88c680b", "*", armapimanagement.UserUpdateParameters{
		Properties: &armapimanagement.UserUpdateParametersProperties{
			Email:     to.Ptr("foobar@outlook.com"),
			FirstName: to.Ptr("foo"),
			LastName:  to.Ptr("bar"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserContract = armapimanagement.UserContract{
	// 	Name: to.Ptr("5931a75ae4bbd512a88c680b"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/users"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
	// 	Properties: &armapimanagement.UserContractProperties{
	// 		Identities: []*armapimanagement.UserIdentityContract{
	// 			{
	// 				ID: to.Ptr("*************"),
	// 				Provider: to.Ptr("Microsoft"),
	// 		}},
	// 		State: to.Ptr(armapimanagement.UserStateActive),
	// 		Email: to.Ptr("foobar@outlook.com"),
	// 		FirstName: to.Ptr("foo"),
	// 		LastName: to.Ptr("bar"),
	// 		RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:58:50.357Z"); return t}()),
	// 	},
	// }
}
