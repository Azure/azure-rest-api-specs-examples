package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/QueryPacksListByResourceGroup.json
func ExampleQueryPacksClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQueryPacksClient().NewListByResourceGroupPager("my-resource-group", nil)
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
		// page.LogAnalyticsQueryPackListResult = armoperationalinsights.LogAnalyticsQueryPackListResult{
		// 	Value: []*armoperationalinsights.LogAnalyticsQueryPack{
		// 		{
		// 			Name: to.Ptr("my-querypack"),
		// 			Type: to.Ptr("microsoft.operationalinsights/querypacks"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armoperationalinsights.LogAnalyticsQueryPackProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				QueryPackID: to.Ptr("d1c8fc00-2b68-441e-8f9b-ded8748dc6aa"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-04T12:37:56.854Z"); return t}()),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-04T12:37:56.854Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-other-querypack"),
		// 			Type: to.Ptr("microsoft.operationalinsights/querypacks"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-other-querypack"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armoperationalinsights.LogAnalyticsQueryPackProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				QueryPackID: to.Ptr("aac8fc00-2b68-441e-8f9b-ded8748dc635"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-04T12:37:56.854Z"); return t}()),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-04T12:37:56.854Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
