package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementTenantAccessSyncState.json
func ExampleTenantConfigurationClient_GetSyncState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTenantConfigurationClient().GetSyncState(ctx, "rg1", "apimService1", armapimanagement.ConfigurationIDNameConfiguration, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TenantConfigurationSyncStateContract = armapimanagement.TenantConfigurationSyncStateContract{
	// 	Name: to.Ptr("syncState"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/tenant/syncState"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tenant/configuration/syncState"),
	// 	Properties: &armapimanagement.TenantConfigurationSyncStateContractProperties{
	// 		Branch: to.Ptr("master"),
	// 		CommitID: to.Ptr("de891c2342c7058dde45e5e624eae7e558c94683"),
	// 		ConfigurationChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-13T00:11:43.862Z"); return t}()),
	// 		IsExport: to.Ptr(true),
	// 		IsGitEnabled: to.Ptr(true),
	// 		IsSynced: to.Ptr(true),
	// 		LastOperationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tenant/configuration/operationResults/6074f0bd093a9d0dac3d7347"),
	// 		SyncDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-13T01:15:53.982Z"); return t}()),
	// 	},
	// }
}
