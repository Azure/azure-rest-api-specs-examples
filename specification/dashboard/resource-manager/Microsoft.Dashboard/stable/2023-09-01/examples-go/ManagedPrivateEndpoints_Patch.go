package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/78eac0bd58633028293cb1ec1709baa200bed9e2/specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/ManagedPrivateEndpoints_Patch.json
func ExampleManagedPrivateEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedPrivateEndpointsClient().BeginUpdate(ctx, "myResourceGroup", "myWorkspace", "myMPEName", armdashboard.ManagedPrivateEndpointUpdateParameters{
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev 2"),
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
	// res.ManagedPrivateEndpointModel = armdashboard.ManagedPrivateEndpointModel{
	// 	Name: to.Ptr("myMPEName"),
	// 	Type: to.Ptr("Microsoft.Dashboard/grafana/managedPrivateEndpoint"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/Microsoft.Dashboard/grafana/myWorkspace/managedPrivateEndpoints/myPrivateEndpointName"),
	// 	SystemData: &armdashboard.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armdashboard.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armdashboard.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armdashboard.ManagedPrivateEndpointModelProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("grafana")},
	// 			PrivateLinkResourceID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-000000000000/resourceGroups/xx-rg/providers/Microsoft.Kusto/Clusters/sampleKustoResource"),
	// 			PrivateLinkResourceRegion: to.Ptr("West US"),
	// 			PrivateLinkServicePrivateIP: to.Ptr("10.0.0.5"),
	// 			PrivateLinkServiceURL: to.Ptr("my-self-hosted-influxdb.westus.mydomain.com"),
	// 			ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
	// 			RequestMessage: to.Ptr("Example Request Message"),
	// 		},
	// 	}
}
