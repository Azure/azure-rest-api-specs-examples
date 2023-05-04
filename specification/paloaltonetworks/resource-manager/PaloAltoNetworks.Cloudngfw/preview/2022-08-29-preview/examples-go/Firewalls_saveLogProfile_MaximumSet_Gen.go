package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fdf43f2fdacf17bd78c0621df44a5c024b61db82/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/Firewalls_saveLogProfile_MaximumSet_Gen.json
func ExampleFirewallsClient_SaveLogProfile_firewallsSaveLogProfileMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFirewallsClient().SaveLogProfile(ctx, "firewall-rg", "firewall1", &armpanngfw.FirewallsClientSaveLogProfileOptions{LogSettings: &armpanngfw.LogSettings{
		ApplicationInsights: &armpanngfw.ApplicationInsights{
			ID:  to.Ptr("aaaaaaaaaaaaaaaa"),
			Key: to.Ptr("aaaaaaaaaaaaa"),
		},
		CommonDestination: &armpanngfw.LogDestination{
			EventHubConfigurations: &armpanngfw.EventHub{
				Name:           to.Ptr("aaaaaaaa"),
				ID:             to.Ptr("aaaaaaaaaa"),
				NameSpace:      to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
				PolicyName:     to.Ptr("aaaaaaaaaaaa"),
				SubscriptionID: to.Ptr("aaaaaaaaaa"),
			},
			MonitorConfigurations: &armpanngfw.MonitorLog{
				ID:             to.Ptr("aaaaaaaaaaaaaaaaaaa"),
				PrimaryKey:     to.Ptr("aaaaaaaaaaaaa"),
				SecondaryKey:   to.Ptr("a"),
				SubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
				Workspace:      to.Ptr("aaaaaaaaaaa"),
			},
			StorageConfigurations: &armpanngfw.StorageAccount{
				AccountName:    to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
				ID:             to.Ptr("aaaaaaaaaaaaaaa"),
				SubscriptionID: to.Ptr("aaaaaaaaa"),
			},
		},
		DecryptLogDestination: &armpanngfw.LogDestination{
			EventHubConfigurations: &armpanngfw.EventHub{
				Name:           to.Ptr("aaaaaaaa"),
				ID:             to.Ptr("aaaaaaaaaa"),
				NameSpace:      to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
				PolicyName:     to.Ptr("aaaaaaaaaaaa"),
				SubscriptionID: to.Ptr("aaaaaaaaaa"),
			},
			MonitorConfigurations: &armpanngfw.MonitorLog{
				ID:             to.Ptr("aaaaaaaaaaaaaaaaaaa"),
				PrimaryKey:     to.Ptr("aaaaaaaaaaaaa"),
				SecondaryKey:   to.Ptr("a"),
				SubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
				Workspace:      to.Ptr("aaaaaaaaaaa"),
			},
			StorageConfigurations: &armpanngfw.StorageAccount{
				AccountName:    to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
				ID:             to.Ptr("aaaaaaaaaaaaaaa"),
				SubscriptionID: to.Ptr("aaaaaaaaa"),
			},
		},
		LogOption: to.Ptr(armpanngfw.LogOptionSAMEDESTINATION),
		LogType:   to.Ptr(armpanngfw.LogTypeTRAFFIC),
		ThreatLogDestination: &armpanngfw.LogDestination{
			EventHubConfigurations: &armpanngfw.EventHub{
				Name:           to.Ptr("aaaaaaaa"),
				ID:             to.Ptr("aaaaaaaaaa"),
				NameSpace:      to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
				PolicyName:     to.Ptr("aaaaaaaaaaaa"),
				SubscriptionID: to.Ptr("aaaaaaaaaa"),
			},
			MonitorConfigurations: &armpanngfw.MonitorLog{
				ID:             to.Ptr("aaaaaaaaaaaaaaaaaaa"),
				PrimaryKey:     to.Ptr("aaaaaaaaaaaaa"),
				SecondaryKey:   to.Ptr("a"),
				SubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
				Workspace:      to.Ptr("aaaaaaaaaaa"),
			},
			StorageConfigurations: &armpanngfw.StorageAccount{
				AccountName:    to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
				ID:             to.Ptr("aaaaaaaaaaaaaaa"),
				SubscriptionID: to.Ptr("aaaaaaaaa"),
			},
		},
		TrafficLogDestination: &armpanngfw.LogDestination{
			EventHubConfigurations: &armpanngfw.EventHub{
				Name:           to.Ptr("aaaaaaaa"),
				ID:             to.Ptr("aaaaaaaaaa"),
				NameSpace:      to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
				PolicyName:     to.Ptr("aaaaaaaaaaaa"),
				SubscriptionID: to.Ptr("aaaaaaaaaa"),
			},
			MonitorConfigurations: &armpanngfw.MonitorLog{
				ID:             to.Ptr("aaaaaaaaaaaaaaaaaaa"),
				PrimaryKey:     to.Ptr("aaaaaaaaaaaaa"),
				SecondaryKey:   to.Ptr("a"),
				SubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
				Workspace:      to.Ptr("aaaaaaaaaaa"),
			},
			StorageConfigurations: &armpanngfw.StorageAccount{
				AccountName:    to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
				ID:             to.Ptr("aaaaaaaaaaaaaaa"),
				SubscriptionID: to.Ptr("aaaaaaaaa"),
			},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
