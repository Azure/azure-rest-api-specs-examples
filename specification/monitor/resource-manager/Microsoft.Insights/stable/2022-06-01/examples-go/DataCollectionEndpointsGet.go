package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionEndpointsGet.json
func ExampleDataCollectionEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataCollectionEndpointsClient().Get(ctx, "myResourceGroup", "myCollectionEndpoint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataCollectionEndpointResource = armmonitor.DataCollectionEndpointResource{
	// 	Name: to.Ptr("myCollectionEndpoint"),
	// 	Type: to.Ptr("Microsoft.Insights/dataCollectionEndpoints"),
	// 	Etag: to.Ptr("070057da-0000-0000-0000-5ba70d6c0000"),
	// 	ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/myCollectionEndpoint"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armmonitor.DataCollectionEndpointResourceProperties{
	// 		ConfigurationAccess: &armmonitor.DataCollectionEndpointConfigurationAccess{
	// 			Endpoint: to.Ptr("https://mycollectionendpoint-abcd.eastus-1.control.monitor.azure.com"),
	// 		},
	// 		LogsIngestion: &armmonitor.DataCollectionEndpointLogsIngestion{
	// 			Endpoint: to.Ptr("https://mycollectionendpoint-abcd.eastus-1.ingest.monitor.azure.com"),
	// 		},
	// 		NetworkACLs: &armmonitor.DataCollectionEndpointNetworkACLs{
	// 			PublicNetworkAccess: to.Ptr(armmonitor.KnownPublicNetworkAccessOptionsEnabled),
	// 		},
	// 	},
	// 	SystemData: &armmonitor.DataCollectionEndpointResourceSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("A"),
	// 		"tag2": to.Ptr("B"),
	// 	},
	// }
}
