package armhealthdataaiservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
)

// Generated from example definition: 2024-09-20/PrivateEndpointConnections_ListByDeidService_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_NewListByDeidServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByDeidServicePager("rgopenapi", "deidTest", nil)
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
		// page = armhealthdataaiservices.PrivateEndpointConnectionsClientListByDeidServiceResponse{
		// 	PrivateEndpointConnectionResourceListResult: armhealthdataaiservices.PrivateEndpointConnectionResourceListResult{
		// 		Value: []*armhealthdataaiservices.PrivateEndpointConnectionResource{
		// 			{
		// 				Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
		// 						ID: to.Ptr("gpnxxbbtsysdhhclm"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
		// 						Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
		// 						ActionsRequired: to.Ptr("ulb"),
		// 						Description: to.Ptr("mmvcleuufspfrojjveuith"),
		// 					},
		// 					GroupIDs: []*string{
		// 						to.Ptr("xbdyjqg"),
		// 					},
		// 					ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/deidTest/privateEndpointConnections/aduyb"),
		// 				Name: to.Ptr("aduyb"),
		// 				Type: to.Ptr("umjjkodjuhccrngl"),
		// 				SystemData: &armhealthdataaiservices.SystemData{
		// 					CreatedBy: to.Ptr("p"),
		// 					CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
		// 					LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
