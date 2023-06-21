package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/ResourceGuardProxyCRUD/PutResourceGuardProxy.json
func ExampleDppResourceGuardProxyClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDppResourceGuardProxyClient().CreateOrUpdate(ctx, "SampleResourceGroup", "sampleVault", "swaggerExample", armdataprotection.ResourceGuardProxyBaseResource{
		Properties: &armdataprotection.ResourceGuardProxyBase{
			ResourceGuardResourceID: to.Ptr("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceGuardProxyBaseResource = armdataprotection.ResourceGuardProxyBaseResource{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.DataProtection/vaults/backupResourceGuardProxies"),
	// 	ID: to.Ptr("/subscriptions/5e13b949-1218-4d18-8b99-7e12155ec4f7/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/backupVaults/sampleVault/backupResourceGuardProxies/swaggerExample"),
	// 	Properties: &armdataprotection.ResourceGuardProxyBase{
	// 		Description: to.Ptr("Please take JIT access before performing any of the critical operation"),
	// 		LastUpdatedTime: to.Ptr("2022-09-16T11:44:37.6130487Z"),
	// 		ResourceGuardOperationDetails: []*armdataprotection.ResourceGuardOperationDetail{
	// 			{
	// 				DefaultResourceRequest: to.Ptr("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource/deleteBackupInstanceRequests/default"),
	// 				VaultCriticalOperation: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/delete"),
	// 			},
	// 			{
	// 				DefaultResourceRequest: to.Ptr("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource/deleteResourceGuardProxyRequests/default"),
	// 				VaultCriticalOperation: to.Ptr("Microsoft.DataProtection/backupVaults/backupResourceGuardProxies/delete"),
	// 		}},
	// 		ResourceGuardResourceID: to.Ptr("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource"),
	// 	},
	// }
}
