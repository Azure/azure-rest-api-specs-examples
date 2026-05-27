package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: 2025-01-01-preview/WebPubSub_CreateOrUpdate.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myWebPubSubService", armwebpubsub.ResourceInfo{
		Identity: &armwebpubsub.ManagedIdentity{
			Type: to.Ptr(armwebpubsub.ManagedIdentityTypeSystemAssigned),
		},
		Kind:     to.Ptr(armwebpubsub.ServiceKindWebPubSub),
		Location: to.Ptr("eastus"),
		Properties: &armwebpubsub.Properties{
			DisableAADAuth:   to.Ptr(false),
			DisableLocalAuth: to.Ptr(false),
			LiveTraceConfiguration: &armwebpubsub.LiveTraceConfiguration{
				Categories: []*armwebpubsub.LiveTraceCategory{
					{
						Name:    to.Ptr("ConnectivityLogs"),
						Enabled: to.Ptr("true"),
					},
				},
				Enabled: to.Ptr("false"),
			},
			NetworkACLs: &armwebpubsub.NetworkACLs{
				DefaultAction: to.Ptr(armwebpubsub.ACLActionDeny),
				PrivateEndpoints: []*armwebpubsub.PrivateEndpointACL{
					{
						Name: to.Ptr("mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
						Allow: []*armwebpubsub.WebPubSubRequestType{
							to.Ptr(armwebpubsub.WebPubSubRequestTypeServerConnection),
						},
					},
				},
				PublicNetwork: &armwebpubsub.NetworkACL{
					Allow: []*armwebpubsub.WebPubSubRequestType{
						to.Ptr(armwebpubsub.WebPubSubRequestTypeClientConnection),
					},
				},
			},
			PublicNetworkAccess: to.Ptr("Enabled"),
			SocketIO: &armwebpubsub.SocketIOSettings{
				ServiceMode: to.Ptr("Serverless"),
			},
			TLS: &armwebpubsub.TLSSettings{
				ClientCertEnabled: to.Ptr(false),
			},
		},
		SKU: &armwebpubsub.ResourceSKU{
			Name:     to.Ptr("Premium_P1"),
			Capacity: to.Ptr[int32](1),
			Tier:     to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armwebpubsub.ClientCreateOrUpdateResponse{
	// 	ResourceInfo: armwebpubsub.ResourceInfo{
	// 		Name: to.Ptr("myWebPubSubService"),
	// 		Type: to.Ptr("Microsoft.SignalRService/WebPubSub"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService"),
	// 		Identity: &armwebpubsub.ManagedIdentity{
	// 			Type: to.Ptr(armwebpubsub.ManagedIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Kind: to.Ptr(armwebpubsub.ServiceKindWebPubSub),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armwebpubsub.Properties{
	// 			ApplicationFirewall: &armwebpubsub.ApplicationFirewallSettings{
	// 				ClientConnectionCountRules: []armwebpubsub.ClientConnectionCountRuleClassification{
	// 					&armwebpubsub.ThrottleByJwtSignatureRule{
	// 						Type: to.Ptr(armwebpubsub.ClientConnectionCountRuleDiscriminatorThrottleByJwtSignatureRule),
	// 						MaxCount: to.Ptr[int32](20),
	// 					},
	// 					&armwebpubsub.ThrottleByUserIDRule{
	// 						Type: to.Ptr(armwebpubsub.ClientConnectionCountRuleDiscriminatorThrottleByUserIDRule),
	// 						MaxCount: to.Ptr[int32](20),
	// 					},
	// 					&armwebpubsub.ThrottleByJwtCustomClaimRule{
	// 						Type: to.Ptr(armwebpubsub.ClientConnectionCountRuleDiscriminatorThrottleByJwtCustomClaimRule),
	// 						ClaimName: to.Ptr("claimName"),
	// 						MaxCount: to.Ptr[int32](20),
	// 					},
	// 				},
	// 				ClientTrafficControlRules: []armwebpubsub.ClientTrafficControlRuleClassification{
	// 					&armwebpubsub.TrafficThrottleByJwtSignatureRule{
	// 						Type: to.Ptr(armwebpubsub.ClientTrafficControlRuleDiscriminatorTrafficThrottleByJwtSignatureRule),
	// 						AggregationWindowInSeconds: to.Ptr[int32](60),
	// 						MaxInboundMessageBytes: to.Ptr[int64](1073741824),
	// 					},
	// 					&armwebpubsub.TrafficThrottleByUserIDRule{
	// 						Type: to.Ptr(armwebpubsub.ClientTrafficControlRuleDiscriminatorTrafficThrottleByUserIDRule),
	// 						AggregationWindowInSeconds: to.Ptr[int32](60),
	// 						MaxInboundMessageBytes: to.Ptr[int64](1073741824),
	// 					},
	// 					&armwebpubsub.TrafficThrottleByJwtCustomClaimRule{
	// 						Type: to.Ptr(armwebpubsub.ClientTrafficControlRuleDiscriminatorTrafficThrottleByJwtCustomClaimRule),
	// 						AggregationWindowInSeconds: to.Ptr[int32](60),
	// 						ClaimName: to.Ptr("claimName"),
	// 						MaxInboundMessageBytes: to.Ptr[int64](1073741824),
	// 					},
	// 				},
	// 				MaxClientConnectionLifetimeInSeconds: to.Ptr[int64](5),
	// 			},
	// 			DisableAADAuth: to.Ptr(false),
	// 			DisableLocalAuth: to.Ptr(false),
	// 			ExternalIP: to.Ptr("10.0.0.1"),
	// 			HostName: to.Ptr("mywebpubsubservice.webpubsub.azure.com"),
	// 			LiveTraceConfiguration: &armwebpubsub.LiveTraceConfiguration{
	// 				Categories: []*armwebpubsub.LiveTraceCategory{
	// 					{
	// 						Name: to.Ptr("ConnectivityLogs"),
	// 						Enabled: to.Ptr("true"),
	// 					},
	// 				},
	// 				Enabled: to.Ptr("false"),
	// 			},
	// 			NetworkACLs: &armwebpubsub.NetworkACLs{
	// 				DefaultAction: to.Ptr(armwebpubsub.ACLActionDeny),
	// 				IPRules: []*armwebpubsub.IPRule{
	// 					{
	// 						Action: to.Ptr(armwebpubsub.ACLActionAllow),
	// 						Value: to.Ptr("123.456.789.123/24"),
	// 					},
	// 					{
	// 						Action: to.Ptr(armwebpubsub.ACLActionAllow),
	// 						Value: to.Ptr("123.456.789.123"),
	// 					},
	// 					{
	// 						Action: to.Ptr(armwebpubsub.ACLActionAllow),
	// 						Value: to.Ptr("AppService"),
	// 					},
	// 				},
	// 				PrivateEndpoints: []*armwebpubsub.PrivateEndpointACL{
	// 					{
	// 						Name: to.Ptr("mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 						Allow: []*armwebpubsub.WebPubSubRequestType{
	// 							to.Ptr(armwebpubsub.WebPubSubRequestTypeServerConnection),
	// 						},
	// 					},
	// 				},
	// 				PublicNetwork: &armwebpubsub.NetworkACL{
	// 					Allow: []*armwebpubsub.WebPubSubRequestType{
	// 						to.Ptr(armwebpubsub.WebPubSubRequestTypeClientConnection),
	// 					},
	// 				},
	// 			},
	// 			PrivateEndpointConnections: []*armwebpubsub.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 					Type: to.Ptr("Microsoft.SignalRService/WebPubSub/privateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/privateEndpointConnections/mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 					Properties: &armwebpubsub.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armwebpubsub.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armwebpubsub.PrivateLinkServiceConnectionState{
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armwebpubsub.PrivateLinkServiceConnectionStatusApproved),
	// 						},
	// 						ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
	// 					},
	// 					SystemData: &armwebpubsub.SystemData{
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
	// 						CreatedBy: to.Ptr("string"),
	// 						CreatedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("string"),
	// 						LastModifiedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr("Enabled"),
	// 			PublicPort: to.Ptr[int32](443),
	// 			RegionEndpointEnabled: to.Ptr("Enabled"),
	// 			ResourceLogConfiguration: &armwebpubsub.ResourceLogConfiguration{
	// 				Categories: []*armwebpubsub.ResourceLogCategory{
	// 					{
	// 						Name: to.Ptr("ConnectivityLogs"),
	// 						Enabled: to.Ptr("true"),
	// 					},
	// 				},
	// 			},
	// 			ResourceStopped: to.Ptr("false"),
	// 			ServerPort: to.Ptr[int32](443),
	// 			SocketIO: &armwebpubsub.SocketIOSettings{
	// 				ServiceMode: to.Ptr("Serverless"),
	// 			},
	// 			TLS: &armwebpubsub.TLSSettings{
	// 				ClientCertEnabled: to.Ptr(true),
	// 			},
	// 			Version: to.Ptr("1.0"),
	// 		},
	// 		SKU: &armwebpubsub.ResourceSKU{
	// 			Name: to.Ptr("Premium_P1"),
	// 			Capacity: to.Ptr[int32](1),
	// 			Size: to.Ptr("P1"),
	// 			Tier: to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
	// 		},
	// 		SystemData: &armwebpubsub.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
