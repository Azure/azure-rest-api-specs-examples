package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsCreate.json
func ExampleComponentsClient_CreateOrUpdate_componentCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentsClient().CreateOrUpdate(ctx, "my-resource-group", "my-component", armapplicationinsights.Component{
		Location: to.Ptr("South Central US"),
		Kind:     to.Ptr("web"),
		Properties: &armapplicationinsights.ComponentProperties{
			ApplicationType:     to.Ptr(armapplicationinsights.ApplicationTypeWeb),
			FlowType:            to.Ptr(armapplicationinsights.FlowTypeBluefield),
			RequestSource:       to.Ptr(armapplicationinsights.RequestSourceRest),
			WorkspaceResourceID: to.Ptr("/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Component = armapplicationinsights.Component{
	// 	Name: to.Ptr("my-component"),
	// 	Type: to.Ptr("Microsoft.Insights/components"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component"),
	// 	Location: to.Ptr("South Central US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Kind: to.Ptr("web"),
	// 	Properties: &armapplicationinsights.ComponentProperties{
	// 		AppID: to.Ptr("887f4bfd-b5fd-40d7-9fc3-123456789abc"),
	// 		ApplicationID: to.Ptr("my-component"),
	// 		ApplicationType: to.Ptr(armapplicationinsights.ApplicationTypeWeb),
	// 		ConnectionString: to.Ptr("InstrumentationKey=bc095013-3cf2-45ac-ab47-123456789abc"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-24T01:05:38.593Z"); return t}()),
	// 		DisableIPMasking: to.Ptr(false),
	// 		FlowType: to.Ptr(armapplicationinsights.FlowTypeBluefield),
	// 		HockeyAppID: to.Ptr(""),
	// 		HockeyAppToken: to.Ptr(""),
	// 		IngestionMode: to.Ptr(armapplicationinsights.IngestionModeLogAnalytics),
	// 		InstrumentationKey: to.Ptr("bc095013-3cf2-45ac-ab47-123456789abc"),
	// 		RequestSource: to.Ptr(armapplicationinsights.RequestSourceRest),
	// 		SamplingPercentage: to.Ptr[float64](100),
	// 		TenantID: to.Ptr("f438d567-7177-4fe1-a5e3-123456789abc"),
	// 		WorkspaceResourceID: to.Ptr("/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccessForIngestion: to.Ptr(armapplicationinsights.PublicNetworkAccessTypeEnabled),
	// 		PublicNetworkAccessForQuery: to.Ptr(armapplicationinsights.PublicNetworkAccessTypeEnabled),
	// 	},
	// }
}
