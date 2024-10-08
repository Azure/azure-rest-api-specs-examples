package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/ResourceGuardProxyCRUD/ListResourceGuardProxy.json
func ExampleDppResourceGuardProxyClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDppResourceGuardProxyClient().NewListPager("SampleResourceGroup", "sampleVault", nil)
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
		// page.ResourceGuardProxyBaseResourceList = armdataprotection.ResourceGuardProxyBaseResourceList{
		// 	Value: []*armdataprotection.ResourceGuardProxyBaseResource{
		// 		{
		// 			Name: to.Ptr("swaggerExample"),
		// 			Type: to.Ptr("Microsoft.DataProtection/vaults/backupResourceGuardProxies"),
		// 			ID: to.Ptr("/subscriptions/5e13b949-1218-4d18-8b99-7e12155ec4f7/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/backupVaults/sampleVault/backupResourceGuardProxies/swaggerExample"),
		// 			Properties: &armdataprotection.ResourceGuardProxyBase{
		// 				Description: to.Ptr("Please take JIT access before performing any of the critical operation"),
		// 				LastUpdatedTime: to.Ptr("2022-09-16T11:44:37.6130487Z"),
		// 				ResourceGuardOperationDetails: []*armdataprotection.ResourceGuardOperationDetail{
		// 					{
		// 						DefaultResourceRequest: to.Ptr("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource/deleteBackupInstanceRequests/default"),
		// 						VaultCriticalOperation: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/delete"),
		// 					},
		// 					{
		// 						DefaultResourceRequest: to.Ptr("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource/deleteResourceGuardProxyRequests/default"),
		// 						VaultCriticalOperation: to.Ptr("Microsoft.DataProtection/backupVaults/backupResourceGuardProxies/delete"),
		// 				}},
		// 				ResourceGuardResourceID: to.Ptr("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource"),
		// 			},
		// 	}},
		// }
	}
}
