package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/AvailableSkusList.json
func ExampleAvailableSKUsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableSKUsClient().NewListPager(nil)
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
		// page.SKUList = armdataboxedge.SKUList{
		// 	Value: []*armdataboxedge.SKU{
		// 		{
		// 			Name: to.Ptr(armdataboxedge.SKUNameGateway),
		// 			Availability: to.Ptr(armdataboxedge.SKUAvailabilityAvailable),
		// 			Kind: to.Ptr("AzureDataBoxGateway"),
		// 			LocationInfo: []*armdataboxedge.SKULocationInfo{
		// 				{
		// 					Location: to.Ptr("West US"),
		// 			}},
		// 			Locations: []*string{
		// 				to.Ptr("West US")},
		// 				ResourceType: to.Ptr("dataBoxEdgeDevices"),
		// 				SignupOption: to.Ptr(armdataboxedge.SKUSignupOptionAvailable),
		// 				Tier: to.Ptr(armdataboxedge.SKUTierStandard),
		// 				Version: to.Ptr(armdataboxedge.SKUVersionStable),
		// 		}},
		// 	}
	}
}
