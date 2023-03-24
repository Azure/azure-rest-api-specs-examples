package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsEndpointDelete_WithIdentity_example.json
func ExampleEndpointClient_BeginDelete_deleteADigitalTwinsInstanceEndpointWithIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdigitaltwins.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEndpointClient().BeginDelete(ctx, "resRg", "myDigitalTwinsService", "myendpoint", nil)
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
	// res.EndpointResource = armdigitaltwins.EndpointResource{
	// 	Name: to.Ptr("myendpoint"),
	// 	Type: to.Ptr("Microsoft.DigitalTwins/digitalTwinsInstances/endpoints"),
	// 	ID: to.Ptr("/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourcegroups/resRg/providers/Microsoft.DigitalTwins/digitalTwinsInstances/myDigitalTwinsService/endpoints/myServiceBus"),
	// 	SystemData: &armdigitaltwins.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T17:13:59.4037715Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@example.com"),
	// 		CreatedByType: to.Ptr(armdigitaltwins.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T17:14:02.5281409Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@example.com"),
	// 		LastModifiedByType: to.Ptr(armdigitaltwins.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdigitaltwins.ServiceBus{
	// 		AuthenticationType: to.Ptr(armdigitaltwins.AuthenticationTypeKeyBased),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-19T01:10:34.350Z"); return t}()),
	// 		EndpointType: to.Ptr(armdigitaltwins.EndpointTypeServiceBus),
	// 		ProvisioningState: to.Ptr(armdigitaltwins.EndpointProvisioningStateSucceeded),
	// 		EndpointURI: to.Ptr("sb://mysb.servicebus.windows.net/"),
	// 		EntityPath: to.Ptr("mysbtopic"),
	// 	},
	// }
}
