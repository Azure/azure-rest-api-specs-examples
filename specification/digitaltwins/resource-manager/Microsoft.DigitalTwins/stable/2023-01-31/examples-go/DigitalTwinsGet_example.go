package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsGet_example.json
func ExampleClient_Get_getADigitalTwinsInstanceResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdigitaltwins.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "resRg", "myDigitalTwinsService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Description = armdigitaltwins.Description{
	// 	Name: to.Ptr("myDigitalTwinsService"),
	// 	Type: to.Ptr("Microsoft.DigitalTwins/digitalTwinsInstances"),
	// 	ID: to.Ptr("/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourcegroups/resRg/providers/Microsoft.DigitalTwins/digitalTwinsInstances/myDigitalTwinsService"),
	// 	Location: to.Ptr("westus2"),
	// 	SystemData: &armdigitaltwins.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T17:13:59.403Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@example.com"),
	// 		CreatedByType: to.Ptr(armdigitaltwins.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T17:14:02.528Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("d2baee3d-44c0-41b6-9961-92563de66a97"),
	// 		LastModifiedByType: to.Ptr(armdigitaltwins.CreatedByTypeApplication),
	// 	},
	// 	Properties: &armdigitaltwins.Properties{
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-19T12:55:05.229Z"); return t}()),
	// 		HostName: to.Ptr("https://myDigitalTwinsService.api.wus2.ss.azuredigitaltwins-test.net"),
	// 		LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-06T12:21:58.610Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armdigitaltwins.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armdigitaltwins.PublicNetworkAccessEnabled),
	// 	},
	// }
}
