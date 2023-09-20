package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalR_Get.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "myResourceGroup", "mySignalRService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceInfo = armsignalr.ResourceInfo{
	// 	Name: to.Ptr("mySignalRService"),
	// 	Type: to.Ptr("Microsoft.SignalRService/SignalR"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
	// 	SystemData: &armsignalr.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armsignalr.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armsignalr.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Identity: &armsignalr.ManagedIdentity{
	// 		Type: to.Ptr(armsignalr.ManagedIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	},
	// 	Kind: to.Ptr(armsignalr.ServiceKindSignalR),
	// 	Properties: &armsignalr.Properties{
	// 		Cors: &armsignalr.CorsSettings{
	// 			AllowedOrigins: []*string{
	// 				to.Ptr("https://foo.com"),
	// 				to.Ptr("https://bar.com")},
	// 			},
	// 			DisableAADAuth: to.Ptr(false),
	// 			DisableLocalAuth: to.Ptr(false),
	// 			ExternalIP: to.Ptr("10.0.0.1"),
	// 			Features: []*armsignalr.Feature{
	// 				{
	// 					Flag: to.Ptr(armsignalr.FeatureFlagsServiceMode),
	// 					Properties: map[string]*string{
	// 					},
	// 					Value: to.Ptr("Serverless"),
	// 				},
	// 				{
	// 					Flag: to.Ptr(armsignalr.FeatureFlagsEnableConnectivityLogs),
	// 					Properties: map[string]*string{
	// 					},
	// 					Value: to.Ptr("True"),
	// 				},
	// 				{
	// 					Flag: to.Ptr(armsignalr.FeatureFlagsEnableMessagingLogs),
	// 					Properties: map[string]*string{
	// 					},
	// 					Value: to.Ptr("False"),
	// 				},
	// 				{
	// 					Flag: to.Ptr(armsignalr.FeatureFlagsEnableLiveTrace),
	// 					Properties: map[string]*string{
	// 					},
	// 					Value: to.Ptr("False"),
	// 			}},
	// 			HostName: to.Ptr("mysignalrservice.service.signalr.net"),
	// 			LiveTraceConfiguration: &armsignalr.LiveTraceConfiguration{
	// 				Categories: []*armsignalr.LiveTraceCategory{
	// 					{
	// 						Name: to.Ptr("ConnectivityLogs"),
	// 						Enabled: to.Ptr("true"),
	// 				}},
	// 				Enabled: to.Ptr("false"),
	// 			},
	// 			NetworkACLs: &armsignalr.NetworkACLs{
	// 				DefaultAction: to.Ptr(armsignalr.ACLActionDeny),
	// 				PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
	// 					{
	// 						Allow: []*armsignalr.SignalRRequestType{
	// 							to.Ptr(armsignalr.SignalRRequestTypeServerConnection)},
	// 							Name: to.Ptr("mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 					}},
	// 					PublicNetwork: &armsignalr.NetworkACL{
	// 						Allow: []*armsignalr.SignalRRequestType{
	// 							to.Ptr(armsignalr.SignalRRequestTypeClientConnection)},
	// 						},
	// 					},
	// 					PrivateEndpointConnections: []*armsignalr.PrivateEndpointConnection{
	// 						{
	// 							Name: to.Ptr("mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 							Type: to.Ptr("Microsoft.SignalRService/SignalR/privateEndpointConnections"),
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/privateEndpointConnections/mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 							SystemData: &armsignalr.SystemData{
	// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
	// 								CreatedBy: to.Ptr("string"),
	// 								CreatedByType: to.Ptr(armsignalr.CreatedByTypeUser),
	// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
	// 								LastModifiedBy: to.Ptr("string"),
	// 								LastModifiedByType: to.Ptr(armsignalr.CreatedByTypeUser),
	// 							},
	// 							Properties: &armsignalr.PrivateEndpointConnectionProperties{
	// 								PrivateEndpoint: &armsignalr.PrivateEndpoint{
	// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 								},
	// 								PrivateLinkServiceConnectionState: &armsignalr.PrivateLinkServiceConnectionState{
	// 									ActionsRequired: to.Ptr("None"),
	// 									Status: to.Ptr(armsignalr.PrivateLinkServiceConnectionStatusApproved),
	// 								},
	// 								ProvisioningState: to.Ptr(armsignalr.ProvisioningStateSucceeded),
	// 							},
	// 					}},
	// 					ProvisioningState: to.Ptr(armsignalr.ProvisioningStateSucceeded),
	// 					PublicNetworkAccess: to.Ptr("Enabled"),
	// 					PublicPort: to.Ptr[int32](443),
	// 					ResourceLogConfiguration: &armsignalr.ResourceLogConfiguration{
	// 						Categories: []*armsignalr.ResourceLogCategory{
	// 							{
	// 								Name: to.Ptr("ConnectivityLogs"),
	// 								Enabled: to.Ptr("true"),
	// 						}},
	// 					},
	// 					ServerPort: to.Ptr[int32](443),
	// 					Serverless: &armsignalr.ServerlessSettings{
	// 						ConnectionTimeoutInSeconds: to.Ptr[int32](5),
	// 					},
	// 					TLS: &armsignalr.TLSSettings{
	// 						ClientCertEnabled: to.Ptr(true),
	// 					},
	// 					Upstream: &armsignalr.ServerlessUpstreamSettings{
	// 						Templates: []*armsignalr.UpstreamTemplate{
	// 							{
	// 								URLTemplate: to.Ptr("http://foo.com"),
	// 						}},
	// 					},
	// 					Version: to.Ptr("1.0"),
	// 				},
	// 				SKU: &armsignalr.ResourceSKU{
	// 					Name: to.Ptr("Premium_P1"),
	// 					Capacity: to.Ptr[int32](1),
	// 					Size: to.Ptr("P1"),
	// 					Tier: to.Ptr(armsignalr.SignalRSKUTierPremium),
	// 				},
	// 			}
}
