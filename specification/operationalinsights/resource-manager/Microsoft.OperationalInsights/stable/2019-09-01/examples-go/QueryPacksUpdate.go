package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksUpdate.json
func ExampleQueryPacksClient_CreateOrUpdate_queryPackUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQueryPacksClient().CreateOrUpdate(ctx, "my-resource-group", "my-querypack", armoperationalinsights.LogAnalyticsQueryPack{
		Location: to.Ptr("South Central US"),
		Tags: map[string]*string{
			"Tag1": to.Ptr("Value1"),
		},
		Properties: &armoperationalinsights.LogAnalyticsQueryPackProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LogAnalyticsQueryPack = armoperationalinsights.LogAnalyticsQueryPack{
	// 	Name: to.Ptr("my-querypack"),
	// 	Type: to.Ptr("microsoft.operationalinsights/querypacks"),
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack"),
	// 	Location: to.Ptr("South Central US"),
	// 	Tags: map[string]*string{
	// 		"Tag1": to.Ptr("Value1"),
	// 	},
	// 	Properties: &armoperationalinsights.LogAnalyticsQueryPackProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		QueryPackID: to.Ptr("aac8fc00-2b68-441e-8f9b-ded8748dc635"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-04T12:37:56.8543652Z"); return t}()),
	// 		TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-04T12:37:56.8543652Z"); return t}()),
	// 	},
	// }
}
