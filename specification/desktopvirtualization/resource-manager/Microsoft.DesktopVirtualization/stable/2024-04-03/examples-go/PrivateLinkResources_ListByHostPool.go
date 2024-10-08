package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/PrivateLinkResources_ListByHostPool.json
func ExamplePrivateLinkResourcesClient_NewListByHostPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByHostPoolPager("resourceGroup1", "hostPool1", &armdesktopvirtualization.PrivateLinkResourcesClientListByHostPoolOptions{PageSize: nil,
		IsDescending: nil,
		InitialSkip:  nil,
	})
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
		// page.PrivateLinkResourceListResult = armdesktopvirtualization.PrivateLinkResourceListResult{
		// 	Value: []*armdesktopvirtualization.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("hostpool"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostPool1/privateLinkResources/hostpool"),
		// 			Properties: &armdesktopvirtualization.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("hostpool"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("rdbroker"),
		// 					to.Ptr("rddiagnostics"),
		// 					to.Ptr("rdweb"),
		// 					to.Ptr("rdgateway")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.wvd.microsoft.com")},
		// 					},
		// 			}},
		// 		}
	}
}
