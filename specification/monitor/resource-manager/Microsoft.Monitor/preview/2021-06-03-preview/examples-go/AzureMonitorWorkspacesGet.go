package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Monitor/preview/2021-06-03-preview/examples/AzureMonitorWorkspacesGet.json
func ExampleAzureMonitorWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureMonitorWorkspacesClient().Get(ctx, "myResourceGroup", "myAzureMonitorWorkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureMonitorWorkspaceResource = armmonitor.AzureMonitorWorkspaceResource{
	// 	Name: to.Ptr("myAzureMonitorWorkspace"),
	// 	Type: to.Ptr("Microsoft.Monitor/accounts"),
	// 	ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace"),
	// 	SystemData: &armmonitor.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("A"),
	// 		"tag2": to.Ptr("B"),
	// 	},
	// 	Etag: to.Ptr("070057da-0000-0000-0000-5ba70d6c0000"),
	// 	Properties: &armmonitor.AzureMonitorWorkspaceResourceProperties{
	// 		AccountID: to.Ptr("2df515bf-c3ce-4920-84d4-1d9d16542d9f"),
	// 		DefaultIngestionSettings: &armmonitor.AzureMonitorWorkspaceDefaultIngestionSettings{
	// 			DataCollectionEndpointResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/MA_myAzureMonitorWorkspace_eastus_managed/providers/Microsoft.Insights/dataCollectionEndpoints/myAzureMonitorWorkspace"),
	// 			DataCollectionRuleResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/MA_myAzureMonitorWorkspace_eastus_managed/providers/Microsoft.Insights/dataCollectionRules/myAzureMonitorWorkspace"),
	// 		},
	// 		Metrics: &armmonitor.AzureMonitorWorkspaceMetrics{
	// 			InternalID: to.Ptr("mac_2df515bf-c3ce-4920-84d4-1d9d16542d9f"),
	// 			PrometheusQueryEndpoint: to.Ptr("https://myAzureMonitorWorkspace-v8hx.eastus.prometheus.monitor.azure.com"),
	// 		},
	// 		ProvisioningState: to.Ptr(armmonitor.ProvisioningStateSucceeded),
	// 	},
	// }
}
