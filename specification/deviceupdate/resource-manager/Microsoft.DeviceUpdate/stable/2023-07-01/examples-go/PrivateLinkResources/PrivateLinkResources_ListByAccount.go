package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/PrivateLinkResources/PrivateLinkResources_ListByAccount.json
func ExamplePrivateLinkResourcesClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByAccountPager("test-rg", "contoso", nil)
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
		// page.PrivateLinkResourceListResult = armdeviceupdate.PrivateLinkResourceListResult{
		// 	Value: []*armdeviceupdate.GroupInformation{
		// 		{
		// 			Name: to.Ptr("adu"),
		// 			Type: to.Ptr("Microsoft.DeviceUpdate/accounts/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DeviceUpdate/accounts/contoso/privateLinkResources/adu"),
		// 			Properties: &armdeviceupdate.GroupInformationProperties{
		// 				GroupID: to.Ptr("adu"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("adu")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.adu.microsoft.com")},
		// 					},
		// 			}},
		// 		}
	}
}
