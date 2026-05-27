package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateClientApplication.json
func ExampleClientApplicationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClientApplicationClient().CreateOrUpdate(ctx, "rg1", "apimService1", "testAppId", armapimanagement.ClientApplicationContract{
		Properties: &armapimanagement.ClientApplicationContractProperties{
			Description: to.Ptr("This is just an example application"),
			DisplayName: to.Ptr("Test Application"),
			OwnerID:     to.Ptr("/users/userId"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.ClientApplicationClientCreateOrUpdateResponse{
	// 	ClientApplicationContract: armapimanagement.ClientApplicationContract{
	// 		Name: to.Ptr("testAppId"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/clientApplications"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/clientApplications/testAppId"),
	// 		Properties: &armapimanagement.ClientApplicationContractProperties{
	// 			Description: to.Ptr("This is just an example application"),
	// 			DisplayName: to.Ptr("Test Application"),
	// 			EntraApplicationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			EntraTenantID: to.Ptr("00000000-0000-0000-0000-000000000010"),
	// 			OwnerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/userId"),
	// 			State: to.Ptr(armapimanagement.ClientApplicationStateActive),
	// 		},
	// 	},
	// }
}
