package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsList.json
func ExampleComponentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewComponentsClient().NewListPager(nil)
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
		// page.ComponentListResult = armapplicationinsights.ComponentListResult{
		// 	Value: []*armapplicationinsights.Component{
		// 		{
		// 			Name: to.Ptr("my-component"),
		// 			Type: to.Ptr("Microsoft.Insights/components"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr("web"),
		// 			Properties: &armapplicationinsights.ComponentProperties{
		// 				AppID: to.Ptr("16526d1a-dfba-4362-a9e9-123456789abc"),
		// 				ApplicationID: to.Ptr("my-component"),
		// 				ApplicationType: to.Ptr(armapplicationinsights.ApplicationTypeWeb),
		// 				ConnectionString: to.Ptr("InstrumentationKey=dc5931c7-a7ad-4ad0-89d6-123456789abc"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-14T12:24:05.004Z"); return t}()),
		// 				DisableIPMasking: to.Ptr(false),
		// 				FlowType: to.Ptr(armapplicationinsights.FlowTypeBluefield),
		// 				HockeyAppID: to.Ptr(""),
		// 				HockeyAppToken: to.Ptr(""),
		// 				IngestionMode: to.Ptr(armapplicationinsights.IngestionModeLogAnalytics),
		// 				InstrumentationKey: to.Ptr("dc5931c7-a7ad-4ad0-89d6-123456789abc"),
		// 				RequestSource: to.Ptr(armapplicationinsights.RequestSourceRest),
		// 				SamplingPercentage: to.Ptr[float64](75),
		// 				TenantID: to.Ptr("f438d567-7177-4fe1-a5e3-123456789abc"),
		// 				WorkspaceResourceID: to.Ptr("/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PublicNetworkAccessForIngestion: to.Ptr(armapplicationinsights.PublicNetworkAccessTypeEnabled),
		// 				PublicNetworkAccessForQuery: to.Ptr(armapplicationinsights.PublicNetworkAccessTypeEnabled),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-other-component"),
		// 			Type: to.Ptr("Microsoft.Insights/components"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my-other-resource-group/providers/Microsoft.Insights/components/my-other-component"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr("web"),
		// 			Properties: &armapplicationinsights.ComponentProperties{
		// 				AppID: to.Ptr("887f4bfd-b5fd-40d7-9fc3-123456789abc"),
		// 				ApplicationID: to.Ptr("my-other-component"),
		// 				ApplicationType: to.Ptr(armapplicationinsights.ApplicationTypeWeb),
		// 				ConnectionString: to.Ptr("InstrumentationKey=bc095013-3cf2-45ac-ab47-123456789abc"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-24T01:05:38.593Z"); return t}()),
		// 				DisableIPMasking: to.Ptr(false),
		// 				FlowType: to.Ptr(armapplicationinsights.FlowTypeBluefield),
		// 				HockeyAppID: to.Ptr(""),
		// 				HockeyAppToken: to.Ptr(""),
		// 				IngestionMode: to.Ptr(armapplicationinsights.IngestionModeLogAnalytics),
		// 				InstrumentationKey: to.Ptr("bc095013-3cf2-45ac-ab47-123456789abc"),
		// 				RequestSource: to.Ptr(armapplicationinsights.RequestSourceRest),
		// 				SamplingPercentage: to.Ptr[float64](30),
		// 				TenantID: to.Ptr("f438d567-7177-4fe1-a5e3-123456789abc"),
		// 				WorkspaceResourceID: to.Ptr("/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PublicNetworkAccessForIngestion: to.Ptr(armapplicationinsights.PublicNetworkAccessTypeEnabled),
		// 				PublicNetworkAccessForQuery: to.Ptr(armapplicationinsights.PublicNetworkAccessTypeEnabled),
		// 			},
		// 	}},
		// }
	}
}
