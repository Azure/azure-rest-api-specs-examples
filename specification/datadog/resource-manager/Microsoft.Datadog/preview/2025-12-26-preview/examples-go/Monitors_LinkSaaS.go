package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog/v2"
)

// Generated from example definition: 2025-12-26-preview/Monitors_LinkSaaS.json
func ExampleMonitorResourcesClient_BeginLinkSaaS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("1a2b3c4d-5e6f-7a8b-9c0d-e1f2a3b4c5d6", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorResourcesClient().BeginLinkSaaS(ctx, "myResourceGroup", "myMonitor", armdatadog.SaaSData{
		SaaSResourceID: to.Ptr("/subscriptions/1a2b3c4d-5e6f-7a8b-9c0d-e1f2a3b4c5d6/resourceGroups/myResourceGroup/providers/Microsoft.SaaS/resources/mySaaSResource"),
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
	// res = armdatadog.MonitorResourcesClientLinkSaaSResponse{
	// 	MonitorResource: &armdatadog.MonitorResource{
	// 		ID: to.Ptr("/subscriptions/1a2b3c4d-5e6f-7a8b-9c0d-e1f2a3b4c5d6/resourceGroups/myResourceGroup/providers/Microsoft.Datadog/monitors/myMonitor"),
	// 		Name: to.Ptr("myMonitor"),
	// 		Type: to.Ptr("Microsoft.Datadog/monitors"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armdatadog.MonitorProperties{
	// 			DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
	// 				Name: to.Ptr("myOrg"),
	// 				ID: to.Ptr("myOrg123"),
	// 				Cspm: to.Ptr(false),
	// 				ResourceCollection: to.Ptr(false),
	// 			},
	// 			MarketplaceOfferDetails: &armdatadog.MarketplaceOfferDetails{
	// 				PublisherID: to.Ptr("datadog1591740804488"),
	// 				OfferID: to.Ptr("dd_liftr_v3_decoupled"),
	// 			},
	// 			SaaSData: &armdatadog.SaaSData{
	// 				SaaSResourceID: to.Ptr("/subscriptions/1a2b3c4d-5e6f-7a8b-9c0d-e1f2a3b4c5d6/resourceGroups/myResourceGroup/providers/Microsoft.SaaS/resources/mySaaSResource"),
	// 			},
	// 			LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
	// 			LiftrResourcePreference: to.Ptr[int32](1),
	// 			MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
	// 			ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Environment": to.Ptr("Dev"),
	// 		},
	// 	},
	// }
}
