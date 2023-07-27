package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoClustersUpdate.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "kustorptest", "kustoCluster2", armkusto.ClusterUpdate{
		Location: to.Ptr("westus"),
	}, &armkusto.ClustersClientBeginUpdateOptions{IfMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armkusto.Cluster{
	// 	Name: to.Ptr("kustoCluster2"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("abcd123"),
	// 	Identity: &armkusto.Identity{
	// 		Type: to.Ptr(armkusto.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("faabad1f-4876-463c-af9d-6ba2d2d2394c"),
	// 		TenantID: to.Ptr("b932977f-6277-4ab7-a2cd-5bd21f07aaf4"),
	// 		UserAssignedIdentities: map[string]*armkusto.ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 		},
	// 	},
	// 	Properties: &armkusto.ClusterProperties{
	// 		EnableAutoStop: to.Ptr(true),
	// 		EnableDiskEncryption: to.Ptr(false),
	// 		EnablePurge: to.Ptr(true),
	// 		EnableStreamingIngest: to.Ptr(true),
	// 		EngineType: to.Ptr(armkusto.EngineTypeV3),
	// 		KeyVaultProperties: &armkusto.KeyVaultProperties{
	// 			KeyName: to.Ptr("keyName"),
	// 			KeyVaultURI: to.Ptr("https://dummy.keyvault.com"),
	// 			KeyVersion: to.Ptr("keyVersion"),
	// 		},
	// 		ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 		PublicIPType: to.Ptr(armkusto.PublicIPTypeIPv4),
	// 		RestrictOutboundNetworkAccess: to.Ptr(armkusto.ClusterNetworkAccessFlagDisabled),
	// 	},
	// 	SKU: &armkusto.AzureSKU{
	// 		Name: to.Ptr(armkusto.AzureSKUNameStandardL16AsV3),
	// 		Capacity: to.Ptr[int32](2),
	// 		Tier: to.Ptr(armkusto.AzureSKUTierStandard),
	// 	},
	// }
}
