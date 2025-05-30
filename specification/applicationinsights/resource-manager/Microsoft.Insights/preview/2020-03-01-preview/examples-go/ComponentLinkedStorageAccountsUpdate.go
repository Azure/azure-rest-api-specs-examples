package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsUpdate.json
func ExampleComponentLinkedStorageAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentLinkedStorageAccountsClient().Update(ctx, "someResourceGroupName", "myComponent", armapplicationinsights.StorageTypeServiceProfiler, armapplicationinsights.ComponentLinkedStorageAccountsPatch{
		Properties: &armapplicationinsights.LinkedStorageAccountsProperties{
			LinkedStorageAccount: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/someResourceGroupName/providers/Microsoft.Storage/storageAccounts/storageaccountname"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentLinkedStorageAccounts = armapplicationinsights.ComponentLinkedStorageAccounts{
	// 	Name: to.Ptr("serviceprofile"),
	// 	Type: to.Ptr("microsoft.insights/components/linkedStorageAccounts"),
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/someResourceGroupName/providers/microsoft.insights/components/myComponent/linkedStorageAccounts/serviceprofiler"),
	// 	Properties: &armapplicationinsights.LinkedStorageAccountsProperties{
	// 		LinkedStorageAccount: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/someResourceGroupName/providers/Microsoft.Storage/storageAccounts/storageaccountname"),
	// 	},
	// }
}
