package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateUser.json
func ExampleUserClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().CreateOrUpdate(ctx, "rg1", "apimService1", "5931a75ae4bbd512288c680b", armapimanagement.UserCreateParameters{
		Properties: &armapimanagement.UserCreateParameterProperties{
			Confirmation: to.Ptr(armapimanagement.ConfirmationSignup),
			Email:        to.Ptr("foobar@outlook.com"),
			FirstName:    to.Ptr("foo"),
			LastName:     to.Ptr("bar"),
		},
	}, &armapimanagement.UserClientCreateOrUpdateOptions{Notify: nil,
		IfMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserContract = armapimanagement.UserContract{
	// 	Name: to.Ptr("5931a75ae4bbd512288c680b"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/users"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512288c680b"),
	// 	Properties: &armapimanagement.UserContractProperties{
	// 		Identities: []*armapimanagement.UserIdentityContract{
	// 			{
	// 				ID: to.Ptr("foobar@outlook.com"),
	// 				Provider: to.Ptr("Basic"),
	// 		}},
	// 		State: to.Ptr(armapimanagement.UserStateActive),
	// 		Email: to.Ptr("foobar@outlook.com"),
	// 		FirstName: to.Ptr("foo"),
	// 		Groups: []*armapimanagement.GroupContractProperties{
	// 		},
	// 		LastName: to.Ptr("bar"),
	// 		RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-07T21:21:29.160Z"); return t}()),
	// 	},
	// }
}
