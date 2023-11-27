package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsEndpointsGet_WithIdentity_example.json
func ExampleEndpointClient_NewListPager_getADigitalTwinsInstanceEndpointsWithIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdigitaltwins.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEndpointClient().NewListPager("resRg", "myDigitalTwinsService", nil)
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
		// page.EndpointResourceListResult = armdigitaltwins.EndpointResourceListResult{
		// 	Value: []*armdigitaltwins.EndpointResource{
		// 		{
		// 			Name: to.Ptr("myServiceBus"),
		// 			Type: to.Ptr("Microsoft.DigitalTwins/digitalTwinsInstances/endpoints"),
		// 			ID: to.Ptr("/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourcegroups/resRg/providers/Microsoft.DigitalTwins/digitalTwinsInstances/myDigitalTwinsService/endpoints/myServiceBus"),
		// 			SystemData: &armdigitaltwins.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T17:13:59.403Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@example.com"),
		// 				CreatedByType: to.Ptr(armdigitaltwins.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T17:14:02.528Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@example.com"),
		// 				LastModifiedByType: to.Ptr(armdigitaltwins.CreatedByTypeUser),
		// 			},
		// 			Properties: &armdigitaltwins.ServiceBus{
		// 				AuthenticationType: to.Ptr(armdigitaltwins.AuthenticationTypeIdentityBased),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-19T01:10:34.350Z"); return t}()),
		// 				EndpointType: to.Ptr(armdigitaltwins.EndpointTypeServiceBus),
		// 				ProvisioningState: to.Ptr(armdigitaltwins.EndpointProvisioningStateSucceeded),
		// 				EndpointURI: to.Ptr("sb://mysb.servicebus.windows.net/"),
		// 				EntityPath: to.Ptr("mysbtopic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myServiceBus"),
		// 			Type: to.Ptr("Microsoft.DigitalTwins/digitalTwinsInstances/endpoints"),
		// 			ID: to.Ptr("/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourcegroups/resRg/providers/Microsoft.DigitalTwins/digitalTwinsInstances/myDigitalTwinsService/endpoints/myServiceBus"),
		// 			SystemData: &armdigitaltwins.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T17:13:59.403Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@example.com"),
		// 				CreatedByType: to.Ptr(armdigitaltwins.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T17:14:02.528Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@example.com"),
		// 				LastModifiedByType: to.Ptr(armdigitaltwins.CreatedByTypeUser),
		// 			},
		// 			Properties: &armdigitaltwins.EventHub{
		// 				AuthenticationType: to.Ptr(armdigitaltwins.AuthenticationTypeKeyBased),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-19T01:10:34.350Z"); return t}()),
		// 				EndpointType: to.Ptr(armdigitaltwins.EndpointTypeEventHub),
		// 				ProvisioningState: to.Ptr(armdigitaltwins.EndpointProvisioningStateSucceeded),
		// 				ConnectionStringPrimaryKey: to.Ptr("************"),
		// 				ConnectionStringSecondaryKey: to.Ptr("************"),
		// 			},
		// 	}},
		// }
	}
}
