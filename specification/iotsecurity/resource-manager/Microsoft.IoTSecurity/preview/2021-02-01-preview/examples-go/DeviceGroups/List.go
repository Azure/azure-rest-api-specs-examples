package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/DeviceGroups/List.json
func ExampleDeviceGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeviceGroupsClient().NewListPager("eastus", nil)
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
		// page.DeviceGroupList = armiotsecurity.DeviceGroupList{
		// 	Value: []*armiotsecurity.DeviceGroupModel{
		// 		{
		// 			Name: to.Ptr("myGroup"),
		// 			Type: to.Ptr("Microsoft.IoTSecurity/locations/deviceGroups"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.IoTSecurity/locations/eastus/deviceGroups/myGroup"),
		// 			Properties: map[string]any{
		// 			},
		// 			SystemData: &armiotsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-27T21:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armiotsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-27T21:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armiotsecurity.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
