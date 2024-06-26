package armedgeorderpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/SearchInventories.json
func ExampleAPISClient_NewSearchInventoriesPager_searchInventories() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorderpartner.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPISClient().NewSearchInventoriesPager(armedgeorderpartner.SearchInventoriesRequest{
		FamilyIdentifier: to.Ptr("AzureStackEdge"),
		SerialNumber:     to.Ptr("SerialNumber1"),
	}, nil)
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
		// page.PartnerInventoryList = armedgeorderpartner.PartnerInventoryList{
		// 	Value: []*armedgeorderpartner.PartnerInventory{
		// 		{
		// 			Properties: &armedgeorderpartner.InventoryProperties{
		// 				Configuration: &armedgeorderpartner.ConfigurationData{
		// 					ConfigurationIdentifier: to.Ptr("EdgeP_Base"),
		// 					ConfigurationIdentifierOnDevice: to.Ptr("EdgeP_High"),
		// 					FamilyIdentifier: to.Ptr("AzureStackEdge"),
		// 					ProductIdentifier: to.Ptr("AzureStackEdgeProGPU"),
		// 					ProductLineIdentifier: to.Ptr("AzureStackEdgePL"),
		// 				},
		// 				Inventory: &armedgeorderpartner.InventoryData{
		// 					Location: to.Ptr("Rack"),
		// 					RegistrationAllowed: to.Ptr(true),
		// 					Status: to.Ptr("Healthy"),
		// 				},
		// 				Location: to.Ptr("westus"),
		// 				ManagementResource: &armedgeorderpartner.ManagementResourceData{
		// 					ArmID: to.Ptr("/subscriptions/c783ea86-c85c-4175-b76d-3992656af50d/resourceGroups/EdgeTestRG/providers/Microsoft.DataBoxEdge/DataBoxEdgeDevices/TestEdgeDeviceName1"),
		// 					TenantID: to.Ptr("a783ea86-c85c-4175-b76d-3992656af50d"),
		// 				},
		// 				OrderItem: &armedgeorderpartner.OrderItemData{
		// 					ArmID: to.Ptr("/subscriptions/b783ea86-c85c-4175-b76d-3992656af50d/resourceGroups/TestRG/providers/Microsoft.EdgeOrder/orders/TestOrderName1"),
		// 					OrderItemType: to.Ptr(armedgeorderpartner.OrderItemTypeRental),
		// 				},
		// 				SerialNumber: to.Ptr("SerialNumber1"),
		// 			},
		// 	}},
		// }
	}
}
