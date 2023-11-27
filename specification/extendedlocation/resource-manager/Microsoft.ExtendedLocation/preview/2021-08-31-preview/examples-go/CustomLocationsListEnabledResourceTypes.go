package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb9c8e2ca33e9723c2b2fc849f627329067feb54/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsListEnabledResourceTypes.json
func ExampleCustomLocationsClient_NewListEnabledResourceTypesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomLocationsClient().NewListEnabledResourceTypesPager("testresourcegroup", "customLocation01", nil)
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
		// page.EnabledResourceTypesListResult = armextendedlocation.EnabledResourceTypesListResult{
		// 	Value: []*armextendedlocation.EnabledResourceType{
		// 		{
		// 			Name: to.Ptr("d016ecf26dae90594806aca3c1a6326c668357037f68103587edf2e657824737"),
		// 			Type: to.Ptr("Microsoft.ExtendedLocation/customLocations/enabledResourceTypes"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01/enabledResourceTypes/d016ecf26dae90594806aca3c1a6326c668357037f68103587edf2e657824737"),
		// 			Properties: &armextendedlocation.EnabledResourceTypeProperties{
		// 				ClusterExtensionID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/cldfe2econnectedcluster/providers/Microsoft.KubernetesConfiguration/extensions/vmware-extension"),
		// 				ExtensionType: to.Ptr("arc-vmware"),
		// 				TypesMetadata: []*armextendedlocation.EnabledResourceTypePropertiesTypesMetadataItem{
		// 					{
		// 						APIVersion: to.Ptr("2020-01-01-preview"),
		// 						ResourceProviderNamespace: to.Ptr("Microsoft.VMware"),
		// 						ResourceType: to.Ptr("virtualMachines"),
		// 					},
		// 					{
		// 						APIVersion: to.Ptr("2020-01-22-preview"),
		// 						ResourceProviderNamespace: to.Ptr("Microsoft.VMware"),
		// 						ResourceType: to.Ptr("virtualmachines"),
		// 				}},
		// 			},
		// 			SystemData: &armextendedlocation.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.092Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("266e9d31e5be6be1e919574e25780d5783586d502f0b2cc422e0a228a34e00a6"),
		// 			Type: to.Ptr("Microsoft.ExtendedLocation/customLocations/enabledResourceTypes"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01/enabledResourceTypes/266e9d31e5be6be1e919574e25780d5783586d502f0b2cc422e0a228a34e00a6"),
		// 			Properties: &armextendedlocation.EnabledResourceTypeProperties{
		// 				ClusterExtensionID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/cldfe2econnectedcluster/providers/Microsoft.KubernetesConfiguration/extensions/cassandra-extension"),
		// 				ExtensionType: to.Ptr("cassandradatacentersoperator"),
		// 				TypesMetadata: []*armextendedlocation.EnabledResourceTypePropertiesTypesMetadataItem{
		// 					{
		// 						APIVersion: to.Ptr("2020-01-01-preview"),
		// 						ResourceProviderNamespace: to.Ptr("Microsoft.Cassandra"),
		// 						ResourceType: to.Ptr("cassandraDataCenters"),
		// 					},
		// 					{
		// 						APIVersion: to.Ptr("2020-01-22-preview"),
		// 						ResourceProviderNamespace: to.Ptr("Microsoft.Cassandra"),
		// 						ResourceType: to.Ptr("cassandrabackups"),
		// 				}},
		// 			},
		// 			SystemData: &armextendedlocation.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.092Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 			},
		// 	}},
		// }
	}
}
