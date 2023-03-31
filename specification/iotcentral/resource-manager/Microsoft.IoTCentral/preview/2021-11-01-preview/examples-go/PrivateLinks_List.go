package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/96e52e2b911d533f95a0ad8e324c828d556c5f2b/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateLinks_List.json
func ExamplePrivateLinksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinksClient().NewListPager("resRg", "myIoTCentralApp", nil)
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
		// page.PrivateLinkResourceListResult = armiotcentral.PrivateLinkResourceListResult{
		// 	Value: []*armiotcentral.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("iotApp"),
		// 			Type: to.Ptr("Microsoft.IoTCentral/iotApps/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp/privateLinkResources/iotApp"),
		// 			Properties: &armiotcentral.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("iotApp"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("myIoTCentralApp"),
		// 					to.Ptr("iotc-00000000-0000-0000-0000-000000000000"),
		// 					to.Ptr("iotc-00000000-0000-0000-0"),
		// 					to.Ptr("saas-dps-00000000-0000-0000-0000-000000000000")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.azureiotcentral-test.com"),
		// 						to.Ptr("privatelink.azure-devices.net"),
		// 						to.Ptr("privatelink.servicebus.windows.net"),
		// 						to.Ptr("privatelink.azure-devices-provisioning.net")},
		// 					},
		// 			}},
		// 		}
	}
}
