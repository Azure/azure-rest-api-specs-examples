package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d477c7caa09bf82e22c419be0a99d170552b5892/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-05-01/examples/PrivateLinkResourcesListByConfigurationStore.json
func ExamplePrivateLinkResourcesClient_NewListByConfigurationStorePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByConfigurationStorePager("myResourceGroup", "contoso", nil)
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
		// page.PrivateLinkResourceListResult = armappconfiguration.PrivateLinkResourceListResult{
		// 	Value: []*armappconfiguration.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("configurationStores"),
		// 			Type: to.Ptr("Microsoft.AppConfiguraiton/configurationStores/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/privateLinkResources/configurationStores"),
		// 			Properties: &armappconfiguration.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("configurationStores"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("configurationStores")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.azconfig.io")},
		// 					},
		// 			}},
		// 		}
	}
}
