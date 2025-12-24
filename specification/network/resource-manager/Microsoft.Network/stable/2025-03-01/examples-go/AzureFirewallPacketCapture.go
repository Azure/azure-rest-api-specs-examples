package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/AzureFirewallPacketCapture.json
func ExampleAzureFirewallsClient_BeginPacketCapture() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureFirewallsClient().BeginPacketCapture(ctx, "rg1", "azureFirewall1", armnetwork.FirewallPacketCaptureParameters{
		DurationInSeconds: to.Ptr[int32](300),
		FileName:          to.Ptr("azureFirewallPacketCapture"),
		Filters: []*armnetwork.AzureFirewallPacketCaptureRule{
			{
				DestinationPorts: []*string{
					to.Ptr("4500")},
				Destinations: []*string{
					to.Ptr("20.1.2.0")},
				Sources: []*string{
					to.Ptr("20.1.1.0")},
			},
			{
				DestinationPorts: []*string{
					to.Ptr("123"),
					to.Ptr("80")},
				Destinations: []*string{
					to.Ptr("10.1.2.0")},
				Sources: []*string{
					to.Ptr("10.1.1.0"),
					to.Ptr("10.1.1.1")},
			}},
		Flags: []*armnetwork.AzureFirewallPacketCaptureFlags{
			{
				Type: to.Ptr(armnetwork.AzureFirewallPacketCaptureFlagsTypeSyn),
			},
			{
				Type: to.Ptr(armnetwork.AzureFirewallPacketCaptureFlagsTypeFin),
			}},
		NumberOfPacketsToCapture: to.Ptr[int32](5000),
		SasURL:                   to.Ptr("someSASURL"),
		Protocol:                 to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolAny),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
