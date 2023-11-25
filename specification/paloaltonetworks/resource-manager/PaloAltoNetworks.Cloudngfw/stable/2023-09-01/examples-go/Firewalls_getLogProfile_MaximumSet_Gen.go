package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/Firewalls_getLogProfile_MaximumSet_Gen.json
func ExampleFirewallsClient_GetLogProfile_firewallsGetLogProfileMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallsClient().GetLogProfile(ctx, "firewall-rg", "firewall1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LogSettings = armpanngfw.LogSettings{
	// 	ApplicationInsights: &armpanngfw.ApplicationInsights{
	// 		ID: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 		Key: to.Ptr("aaaaaaaaaaaaa"),
	// 	},
	// 	CommonDestination: &armpanngfw.LogDestination{
	// 		EventHubConfigurations: &armpanngfw.EventHub{
	// 			Name: to.Ptr("aaaaaaaa"),
	// 			ID: to.Ptr("aaaaaaaaaa"),
	// 			NameSpace: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			PolicyName: to.Ptr("aaaaaaaaaaaa"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaaa"),
	// 		},
	// 		MonitorConfigurations: &armpanngfw.MonitorLog{
	// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 			PrimaryKey: to.Ptr("aaaaaaaaaaaaa"),
	// 			SecondaryKey: to.Ptr("a"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
	// 			Workspace: to.Ptr("aaaaaaaaaaa"),
	// 		},
	// 		StorageConfigurations: &armpanngfw.StorageAccount{
	// 			AccountName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 			ID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaa"),
	// 		},
	// 	},
	// 	DecryptLogDestination: &armpanngfw.LogDestination{
	// 		EventHubConfigurations: &armpanngfw.EventHub{
	// 			Name: to.Ptr("aaaaaaaa"),
	// 			ID: to.Ptr("aaaaaaaaaa"),
	// 			NameSpace: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			PolicyName: to.Ptr("aaaaaaaaaaaa"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaaa"),
	// 		},
	// 		MonitorConfigurations: &armpanngfw.MonitorLog{
	// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 			PrimaryKey: to.Ptr("aaaaaaaaaaaaa"),
	// 			SecondaryKey: to.Ptr("a"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
	// 			Workspace: to.Ptr("aaaaaaaaaaa"),
	// 		},
	// 		StorageConfigurations: &armpanngfw.StorageAccount{
	// 			AccountName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 			ID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaa"),
	// 		},
	// 	},
	// 	LogOption: to.Ptr(armpanngfw.LogOptionSAMEDESTINATION),
	// 	LogType: to.Ptr(armpanngfw.LogTypeTRAFFIC),
	// 	ThreatLogDestination: &armpanngfw.LogDestination{
	// 		EventHubConfigurations: &armpanngfw.EventHub{
	// 			Name: to.Ptr("aaaaaaaa"),
	// 			ID: to.Ptr("aaaaaaaaaa"),
	// 			NameSpace: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			PolicyName: to.Ptr("aaaaaaaaaaaa"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaaa"),
	// 		},
	// 		MonitorConfigurations: &armpanngfw.MonitorLog{
	// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 			PrimaryKey: to.Ptr("aaaaaaaaaaaaa"),
	// 			SecondaryKey: to.Ptr("a"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
	// 			Workspace: to.Ptr("aaaaaaaaaaa"),
	// 		},
	// 		StorageConfigurations: &armpanngfw.StorageAccount{
	// 			AccountName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 			ID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaa"),
	// 		},
	// 	},
	// 	TrafficLogDestination: &armpanngfw.LogDestination{
	// 		EventHubConfigurations: &armpanngfw.EventHub{
	// 			Name: to.Ptr("aaaaaaaa"),
	// 			ID: to.Ptr("aaaaaaaaaa"),
	// 			NameSpace: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			PolicyName: to.Ptr("aaaaaaaaaaaa"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaaa"),
	// 		},
	// 		MonitorConfigurations: &armpanngfw.MonitorLog{
	// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 			PrimaryKey: to.Ptr("aaaaaaaaaaaaa"),
	// 			SecondaryKey: to.Ptr("a"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
	// 			Workspace: to.Ptr("aaaaaaaaaaa"),
	// 		},
	// 		StorageConfigurations: &armpanngfw.StorageAccount{
	// 			AccountName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 			ID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 			SubscriptionID: to.Ptr("aaaaaaaaa"),
	// 		},
	// 	},
	// }
}
