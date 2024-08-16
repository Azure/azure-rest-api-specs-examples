package armhealthdataaiservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/PrivateEndpointConnections_Create_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreate(ctx, "rgopenapi", "deidTest", "kgwgrrpabvrsrrvpcgcnfmyfgyrl", armhealthdataaiservices.PrivateEndpointConnectionResource{
		Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{},
			PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
				Status:          to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
				ActionsRequired: to.Ptr("ulb"),
				Description:     to.Ptr("xr"),
			},
		},
	}, nil)
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
	// res = armhealthdataaiservices.PrivateEndpointConnectionsClientCreateResponse{
	// 	PrivateEndpointConnectionResource: &armhealthdataaiservices.PrivateEndpointConnectionResource{
	// 		Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
	// 				ID: to.Ptr("gpnxxbbtsysdhhclm"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
	// 				ActionsRequired: to.Ptr("ulb"),
	// 				Description: to.Ptr("xr"),
	// 			},
	// 			GroupIDs: []*string{
	// 				to.Ptr("xbdyjqg"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/deidTest/privateEndpointConnections/aduyb"),
	// 		Name: to.Ptr("aduyb"),
	// 		Type: to.Ptr("umjjkodjuhccrngl"),
	// 		SystemData: &armhealthdataaiservices.SystemData{
	// 			CreatedBy: to.Ptr("p"),
	// 			CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 			LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 		},
	// 	},
	// }
}
