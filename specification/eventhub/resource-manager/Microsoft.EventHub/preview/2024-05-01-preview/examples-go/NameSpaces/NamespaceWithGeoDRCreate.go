package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5759c77eee2d57bdb9e47aa1805d0ffb61704f2d/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/NameSpaces/NamespaceWithGeoDRCreate.json
func ExampleNamespacesClient_BeginCreateOrUpdate_namespaceWithGeoDrCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginCreateOrUpdate(ctx, "ResurceGroupSample", "NamespaceGeoDRCreateSample", armeventhub.EHNamespace{
		Location: to.Ptr("East US"),
		Properties: &armeventhub.EHNamespaceProperties{
			GeoDataReplication: &armeventhub.GeoDataReplicationProperties{
				Locations: []*armeventhub.NamespaceReplicaLocation{
					{
						LocationName: to.Ptr("eastus"),
						RoleType:     to.Ptr(armeventhub.GeoDRRoleTypePrimary),
					},
					{
						LocationName: to.Ptr("westus"),
						RoleType:     to.Ptr(armeventhub.GeoDRRoleTypeSecondary),
					},
					{
						LocationName: to.Ptr("centralus"),
						RoleType:     to.Ptr(armeventhub.GeoDRRoleTypeSecondary),
					}},
				MaxReplicationLagDurationInSeconds: to.Ptr[int32](60),
			},
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
	// res.EHNamespace = armeventhub.EHNamespace{
	// 	Name: to.Ptr("NamespaceGeoDRCreateSample"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces"),
	// 	ID: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.EventHub/namespaces/NamespaceGeoDRCreateSample"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armeventhub.EHNamespaceProperties{
	// 		ClusterArmID: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.EventHub/clusters/enc-test"),
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-16T22:36:06.107Z"); return t}()),
	// 		DisableLocalAuth: to.Ptr(false),
	// 		GeoDataReplication: &armeventhub.GeoDataReplicationProperties{
	// 			Locations: []*armeventhub.NamespaceReplicaLocation{
	// 				{
	// 					LocationName: to.Ptr("eastus"),
	// 					ReplicaState: to.Ptr("Active"),
	// 					RoleType: to.Ptr(armeventhub.GeoDRRoleTypePrimary),
	// 				},
	// 				{
	// 					LocationName: to.Ptr("westus"),
	// 					ReplicaState: to.Ptr("Active"),
	// 					RoleType: to.Ptr(armeventhub.GeoDRRoleTypeSecondary),
	// 				},
	// 				{
	// 					LocationName: to.Ptr("centralus"),
	// 					ReplicaState: to.Ptr("Creating"),
	// 					RoleType: to.Ptr(armeventhub.GeoDRRoleTypeSecondary),
	// 			}},
	// 			MaxReplicationLagDurationInSeconds: to.Ptr[int32](60),
	// 		},
	// 		IsAutoInflateEnabled: to.Ptr(false),
	// 		KafkaEnabled: to.Ptr(false),
	// 		MaximumThroughputUnits: to.Ptr[int32](0),
	// 		MetricID: to.Ptr("MetricGUID:NamespaceGeoDRCreateSample"),
	// 		MinimumTLSVersion: to.Ptr(armeventhub.TLSVersionOne2),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceBusEndpoint: to.Ptr("https://NamespaceGeoDRCreateSample.servicebus.windows-int.net:443/"),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-16T22:37:42.290Z"); return t}()),
	// 		ZoneRedundant: to.Ptr(false),
	// 	},
	// 	SKU: &armeventhub.SKU{
	// 		Name: to.Ptr(armeventhub.SKUNameStandard),
	// 		Capacity: to.Ptr[int32](1),
	// 		Tier: to.Ptr(armeventhub.SKUTierStandard),
	// 	},
	// }
}
