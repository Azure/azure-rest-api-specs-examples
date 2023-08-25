package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiOperationTag.json
func ExampleTagClient_GetByOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTagClient().GetByOperation(ctx, "rg1", "apimService1", "59d6bb8f1f7fab13dc67ec9b", "59d6bb8f1f7fab13dc67ec9a", "59306a29e4bbd510dc24e5f9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TagContract = armapimanagement.TagContract{
	// 	Name: to.Ptr("59306a29e4bbd510dc24e5f9"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/tags"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tags/59306a29e4bbd510dc24e5f9"),
	// 	Properties: &armapimanagement.TagContractProperties{
	// 		DisplayName: to.Ptr("tag1"),
	// 	},
	// }
}
