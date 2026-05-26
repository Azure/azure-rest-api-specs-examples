package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateGroup.json
func ExampleGroupClient_CreateOrUpdate_apiManagementCreateGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGroupClient().CreateOrUpdate(ctx, "rg1", "apimService1", "tempgroup", armapimanagement.GroupCreateParameters{
		Properties: &armapimanagement.GroupCreateParametersProperties{
			DisplayName: to.Ptr("temp group"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.GroupClientCreateOrUpdateResponse{
	// 	GroupContract: armapimanagement.GroupContract{
	// 		Name: to.Ptr("tempgroup"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/groups"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/groups/tempgroup"),
	// 		Properties: &armapimanagement.GroupContractProperties{
	// 			Type: to.Ptr(armapimanagement.GroupTypeCustom),
	// 			DisplayName: to.Ptr("temp group"),
	// 		},
	// 	},
	// }
}
