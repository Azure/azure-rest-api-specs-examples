package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/ResourceGuardProxyCRUD/ListResourceGuardProxy.json
func ExampleResourceGuardProxiesClient_NewGetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceGuardProxiesClient().NewGetPager("sampleVault", "SampleResourceGroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ResourceGuardProxyBaseResourceList = armrecoveryservicesbackup.ResourceGuardProxyBaseResourceList{
		// 	Value: []*armrecoveryservicesbackup.ResourceGuardProxyBaseResource{
		// 		{
		// 			Name: to.Ptr("swaggerExample"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupResourceGuardProxies"),
		// 			ID: to.Ptr("/backupmanagement/resources/sampleVault/backupResourceGuardProxies/swaggerExample"),
		// 			Properties: &armrecoveryservicesbackup.ResourceGuardProxyBase{
		// 				Description: to.Ptr("Please take JIT access before performing any of the critical operation"),
		// 				LastUpdatedTime: to.Ptr("2021-02-11T12:20:47.8210031Z"),
		// 				ResourceGuardOperationDetails: []*armrecoveryservicesbackup.ResourceGuardOperationDetail{
		// 					{
		// 						DefaultResourceRequest: to.Ptr("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew/deleteResourceGuardProxyRequests/default"),
		// 						VaultCriticalOperation: to.Ptr("Microsoft.DataProtection/resourceGuards/deleteResourceGuardProxyRequests"),
		// 					},
		// 					{
		// 						DefaultResourceRequest: to.Ptr("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew/disableSoftDeleteRequests/default"),
		// 						VaultCriticalOperation: to.Ptr("Microsoft.DataProtection/resourceGuards/disableSoftDeleteRequests"),
		// 				}},
		// 				ResourceGuardResourceID: to.Ptr("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew"),
		// 			},
		// 	}},
		// }
	}
}
