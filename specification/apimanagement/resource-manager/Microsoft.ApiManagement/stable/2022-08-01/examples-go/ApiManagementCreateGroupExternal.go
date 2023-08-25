package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGroupExternal.json
func ExampleGroupClient_CreateOrUpdate_apiManagementCreateGroupExternal() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGroupClient().CreateOrUpdate(ctx, "rg1", "apimService1", "aadGroup", armapimanagement.GroupCreateParameters{
		Properties: &armapimanagement.GroupCreateParametersProperties{
			Type:        to.Ptr(armapimanagement.GroupTypeExternal),
			Description: to.Ptr("new group to test"),
			DisplayName: to.Ptr("NewGroup (samiraad.onmicrosoft.com)"),
			ExternalID:  to.Ptr("aad://samiraad.onmicrosoft.com/groups/83cf2753-5831-4675-bc0e-2f8dc067c58d"),
		},
	}, &armapimanagement.GroupClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GroupContract = armapimanagement.GroupContract{
	// 	Name: to.Ptr("aadGroup"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/groups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/groups/aadGroup"),
	// 	Properties: &armapimanagement.GroupContractProperties{
	// 		Type: to.Ptr(armapimanagement.GroupTypeExternal),
	// 		Description: to.Ptr("new group to test"),
	// 		DisplayName: to.Ptr("NewGroup (samiraad.onmicrosoft.com)"),
	// 		ExternalID: to.Ptr("aad://samiraad.onmicrosoft.com/groups/83cf2753-5831-4675-bc0e-2f8dc067c58d"),
	// 	},
	// }
}
