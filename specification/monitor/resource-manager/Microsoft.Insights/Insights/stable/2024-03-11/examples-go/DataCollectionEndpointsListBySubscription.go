package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-03-11/DataCollectionEndpointsListBySubscription.json
func ExampleDataCollectionEndpointsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataCollectionEndpointsClient().NewListBySubscriptionPager(nil)
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
		// page = armmonitor.DataCollectionEndpointsClientListBySubscriptionResponse{
		// 	DataCollectionEndpointResourceListResult: armmonitor.DataCollectionEndpointResourceListResult{
		// 		Value: []*armmonitor.DataCollectionEndpointResource{
		// 			{
		// 				Name: to.Ptr("myDataCollectionEndpoint"),
		// 				Type: to.Ptr("Microsoft.Insights/dataCollectionEndpoints"),
		// 				Etag: to.Ptr("0000ae31-0000-0100-0000-65fd441c0000"),
		// 				ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/myDataCollectionEndpoint"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armmonitor.DataCollectionEndpointResourceProperties{
		// 					ConfigurationAccess: &armmonitor.DataCollectionEndpointConfigurationAccess{
		// 						Endpoint: to.Ptr("https://mydatacollectionendpoint-71ib.eastus-1.handler.control.monitor.azure.com"),
		// 					},
		// 					ImmutableID: to.Ptr("dce-cd85a330eb664aa3bc7883cc394b0000"),
		// 					LogsIngestion: &armmonitor.DataCollectionEndpointLogsIngestion{
		// 						Endpoint: to.Ptr("https://mydatacollectionendpoint-71ib.eastus-1.ingest.monitor.azure.com"),
		// 					},
		// 					MetricsIngestion: &armmonitor.DataCollectionEndpointMetricsIngestion{
		// 						Endpoint: to.Ptr("https://mydatacollectionendpoint-71ib.eastus-1.metrics.ingest.monitor.azure.com"),
		// 					},
		// 					NetworkACLs: &armmonitor.DataCollectionEndpointNetworkACLs{
		// 						PublicNetworkAccess: to.Ptr(armmonitor.KnownPublicNetworkAccessOptionsEnabled),
		// 					},
		// 					ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionEndpointProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armmonitor.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T08:40:59.2431682Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T08:40:59.2431682Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user1"),
		// 					LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("herDataCollectionEndpoint"),
		// 				Type: to.Ptr("Microsoft.Insights/dataCollectionEndpoints"),
		// 				Etag: to.Ptr("0000f831-0000-0100-0000-65fd4c290000"),
		// 				ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/herDataCollectionEndpoint"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armmonitor.DataCollectionEndpointResourceProperties{
		// 					ConfigurationAccess: &armmonitor.DataCollectionEndpointConfigurationAccess{
		// 						Endpoint: to.Ptr("https://herdatacollectionendpoint-6uww.eastus-1.handler.control.monitor.azure.com"),
		// 					},
		// 					ImmutableID: to.Ptr("dce-b4ab5f4638d54d38932ac69699e40000"),
		// 					LogsIngestion: &armmonitor.DataCollectionEndpointLogsIngestion{
		// 						Endpoint: to.Ptr("https://herdatacollectionendpoint-6uww.eastus-1.ingest.monitor.azure.com"),
		// 					},
		// 					MetricsIngestion: &armmonitor.DataCollectionEndpointMetricsIngestion{
		// 						Endpoint: to.Ptr("https://herdatacollectionendpoint-6uww.eastus-1.metrics.ingest.monitor.azure.com"),
		// 					},
		// 					NetworkACLs: &armmonitor.DataCollectionEndpointNetworkACLs{
		// 						PublicNetworkAccess: to.Ptr(armmonitor.KnownPublicNetworkAccessOptionsEnabled),
		// 					},
		// 					ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionEndpointProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armmonitor.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T09:15:20.578873Z"); return t}()),
		// 					CreatedBy: to.Ptr("user2"),
		// 					CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T09:15:20.578873Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("A"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
