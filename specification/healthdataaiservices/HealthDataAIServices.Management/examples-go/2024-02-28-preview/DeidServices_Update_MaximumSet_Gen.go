package armhealthdataaiservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/DeidServices_Update_MaximumSet_Gen.json
func ExampleDeidServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeidServicesClient().BeginUpdate(ctx, "rgopenapi", "deidTest", armhealthdataaiservices.DeidUpdate{
		Identity: &armhealthdataaiservices.ManagedServiceIdentityUpdate{
			Type:                   to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{},
		},
		Tags: map[string]*string{},
		Properties: &armhealthdataaiservices.DeidPropertiesUpdate{
			PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
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
	// res = armhealthdataaiservices.DeidServicesClientUpdateResponse{
	// 	DeidService: &armhealthdataaiservices.DeidService{
	// 		Properties: &armhealthdataaiservices.DeidServiceProperties{
	// 			ProvisioningState: to.Ptr(armhealthdataaiservices.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
	// 			PrivateEndpointConnections: []*armhealthdataaiservices.PrivateEndpointConnection{
	// 				{
	// 					Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
	// 							ID: to.Ptr("gpnxxbbtsysdhhclm"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
	// 							Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
	// 							ActionsRequired: to.Ptr("ulb"),
	// 							Description: to.Ptr("ddxuoved"),
	// 						},
	// 						GroupIDs: []*string{
	// 							to.Ptr("xbdyjqg"),
	// 						},
	// 						ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 					ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih/privateEndpointConnections/mdwvqjtwcjcvrh"),
	// 					Name: to.Ptr("mdwvqjtwcjcvrh"),
	// 					Type: to.Ptr("bzxabjlpbwreez"),
	// 					SystemData: &armhealthdataaiservices.SystemData{
	// 						CreatedBy: to.Ptr("p"),
	// 						CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 						LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			ServiceURL: to.Ptr("woponotsxaippkvhwmibffywnqcfru"),
	// 		},
	// 		Identity: &armhealthdataaiservices.ManagedServiceIdentity{
	// 			Type: to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("a82361f4-5320-4a26-8d1b-45832d2164dd"),
	// 			TenantID: to.Ptr("53a6a686-ae15-4a1d-badf-3e7947918893"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("qwyhvdwcsjulggagdqxlmazcl"),
	// 		ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih"),
	// 		Name: to.Ptr("nlrthrxaukih"),
	// 		Type: to.Ptr("slyfiibvwlhfdpzjynsywhbfauexk"),
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
