package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/Monitors_Get.json
func ExampleMonitorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().Get(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 			Cspm: to.Ptr(false),
	// 			ID: to.Ptr("myOrg123"),
	// 		},
	// 		LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
	// 		LiftrResourcePreference: to.Ptr[int32](1),
	// 		MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
	// 		ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}
