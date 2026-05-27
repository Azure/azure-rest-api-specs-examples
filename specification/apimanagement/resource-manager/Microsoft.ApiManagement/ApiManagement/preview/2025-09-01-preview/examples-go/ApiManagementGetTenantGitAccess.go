package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementGetTenantGitAccess.json
func ExampleTenantAccessClient_Get_apiManagementGetTenantGitAccess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTenantAccessClient().Get(ctx, "rg1", "apimService1", armapimanagement.AccessIDNameGitAccess, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.TenantAccessClientGetResponse{
	// 	AccessInformationContract: armapimanagement.AccessInformationContract{
	// 		Name: to.Ptr("gitAccess"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/tenant"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tenant/gitAccess"),
	// 		Properties: &armapimanagement.AccessInformationContractProperties{
	// 			Enabled: to.Ptr(true),
	// 			PrincipalID: to.Ptr("git"),
	// 		},
	// 	},
	// }
}
