package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/RedisEnterpriseUpdate.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdate(ctx, "rg1", "cache1", armredisenterprise.ClusterUpdate{
		Properties: &armredisenterprise.ClusterProperties{
			MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
		},
		SKU: &armredisenterprise.SKU{
			Name:     to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
			Capacity: to.Ptr[int32](9),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
		},
	}, nil)
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
	// res.Cluster = armredisenterprise.Cluster{
	// 	Name: to.Ptr("cache1"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 	},
	// 	Identity: &armredisenterprise.ManagedServiceIdentity{
	// 		Type: to.Ptr(armredisenterprise.ManagedServiceIdentityTypeNone),
	// 	},
	// 	Properties: &armredisenterprise.ClusterProperties{
	// 		Encryption: &armredisenterprise.ClusterPropertiesEncryption{
	// 		},
	// 		HostName: to.Ptr("cache1.westus.something.azure.com"),
	// 		MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		RedisVersion: to.Ptr("5"),
	// 		ResourceState: to.Ptr(armredisenterprise.ResourceStateUpdating),
	// 	},
	// 	SKU: &armredisenterprise.SKU{
	// 		Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
	// 		Capacity: to.Ptr[int32](9),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1"),
	// 		to.Ptr("2"),
	// 		to.Ptr("3")},
	// 	}
}
