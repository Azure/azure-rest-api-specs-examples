package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_Update.json
func ExampleMonitorsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginUpdate(ctx, "myResourceGroup", "myMonitor", &armdatadog.MonitorsClientBeginUpdateOptions{Body: nil})
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
	// res.MonitorResource = armdatadog.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Microsoft.Datadog/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armdatadog.MonitorProperties{
	// 		DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
	// 			Name: to.Ptr("myOrg"),
	// 			ID: to.Ptr("myOrg123"),
	// 		},
	// 		LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
	// 		LiftrResourcePreference: to.Ptr[int32](1),
	// 		MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
	// 		ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armdatadog.ResourceSKU{
	// 		Name: to.Ptr("free_Monthly"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}
