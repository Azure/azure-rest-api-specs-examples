package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksGet.json
func ExampleQueryPacksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQueryPacksClient().Get(ctx, "my-resource-group", "my-querypack", nil)
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
	// 	Location: to.Ptr("southcentralus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armoperationalinsights.LogAnalyticsQueryPackProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		QueryPackID: to.Ptr("d1c8fc00-2b68-441e-8f9b-ded8748dc635"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-04T12:37:56.854Z"); return t}()),
	// 		TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-04T12:37:56.854Z"); return t}()),
	// 	},
	// }
}
