package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: 2025-02-14/KustoClusterUpdateCMKKeyRotation.json
func ExampleClustersClient_BeginUpdate_kustoClusterUpdateCmkKeyRotation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "kustoRgTest", "kustoClusterCMK", armkusto.ClusterUpdate{
		Location: to.Ptr("westus"),
		Properties: &armkusto.ClusterProperties{
			KeyVaultProperties: &armkusto.KeyVaultProperties{
				KeyVaultURI:  to.Ptr("https://myvault.vault.azure.net"),
				KeyName:      to.Ptr("myClusterCMKKey"),
				KeyVersion:   to.Ptr("87654321-4321-4321-4321-210987654321"),
				UserIdentity: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoClusterIdentity"),
			},
		},
	}, &armkusto.ClustersClientBeginUpdateOptions{
		IfMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkusto.ClustersClientUpdateResponse{
	// 	Cluster: armkusto.Cluster{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.Kusto/Clusters/kustoClusterCMK"),
	// 		Name: to.Ptr("kustoClusterCMK"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armkusto.ClusterProperties{
	// 			ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 			EnableDiskEncryption: to.Ptr(false),
	// 			EnableStreamingIngest: to.Ptr(true),
	// 			EnablePurge: to.Ptr(true),
	// 			EnableAutoStop: to.Ptr(true),
	// 			PublicIPType: to.Ptr(armkusto.PublicIPTypeIPv4),
	// 			KeyVaultProperties: &armkusto.KeyVaultProperties{
	// 				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net"),
	// 				KeyName: to.Ptr("myClusterCMKKey"),
	// 				KeyVersion: to.Ptr("87654321-4321-4321-4321-210987654321"),
	// 				UserIdentity: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoClusterIdentity"),
	// 			},
	// 			EngineType: to.Ptr(armkusto.EngineTypeV3),
	// 			RestrictOutboundNetworkAccess: to.Ptr(armkusto.ClusterNetworkAccessFlagDisabled),
	// 		},
	// 		SKU: &armkusto.AzureSKU{
	// 			Name: to.Ptr(armkusto.AzureSKUNameStandardL16AsV3),
	// 			Tier: to.Ptr(armkusto.AzureSKUTierStandard),
	// 			Capacity: to.Ptr[int32](2),
	// 		},
	// 		Identity: &armkusto.Identity{
	// 			Type: to.Ptr(armkusto.IdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armkusto.ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 				"/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoClusterIdentity": &armkusto.ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 					ClientID: to.Ptr("11111111-2222-3333-4444-555555555555"),
	// 					PrincipalID: to.Ptr("66666666-7777-8888-9999-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Etag: to.Ptr("abcd1234"),
	// 	},
	// }
}
