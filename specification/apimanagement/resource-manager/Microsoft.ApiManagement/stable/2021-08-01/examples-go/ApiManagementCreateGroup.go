package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGroup.json
func ExampleGroupClient_CreateOrUpdate_apiManagementCreateGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGroupClient().CreateOrUpdate(ctx, "rg1", "apimService1", "tempgroup", armapimanagement.GroupCreateParameters{
		Properties: &armapimanagement.GroupCreateParametersProperties{
			DisplayName: to.Ptr("temp group"),
		},
	}, &armapimanagement.GroupClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GroupContract = armapimanagement.GroupContract{
	// 	Name: to.Ptr("tempgroup"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/groups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/groups/tempgroup"),
	// 	Properties: &armapimanagement.GroupContractProperties{
	// 		Type: to.Ptr(armapimanagement.GroupTypeCustom),
	// 		DisplayName: to.Ptr("temp group"),
	// 	},
	// }
}
