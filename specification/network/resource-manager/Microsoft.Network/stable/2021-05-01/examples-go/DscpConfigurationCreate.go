package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DscpConfigurationCreate.json
func ExampleDscpConfigurationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewDscpConfigurationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<dscp-configuration-name>",
		armnetwork.DscpConfiguration{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.DscpConfigurationPropertiesFormat{
				QosDefinitionCollection: []*armnetwork.QosDefinition{
					{
						DestinationIPRanges: []*armnetwork.QosIPRange{
							{
								EndIP:   to.Ptr("<end-ip>"),
								StartIP: to.Ptr("<start-ip>"),
							}},
						DestinationPortRanges: []*armnetwork.QosPortRange{
							{
								End:   to.Ptr[int32](15),
								Start: to.Ptr[int32](15),
							}},
						Markings: []*int32{
							to.Ptr[int32](1)},
						SourceIPRanges: []*armnetwork.QosIPRange{
							{
								EndIP:   to.Ptr("<end-ip>"),
								StartIP: to.Ptr("<start-ip>"),
							}},
						SourcePortRanges: []*armnetwork.QosPortRange{
							{
								End:   to.Ptr[int32](11),
								Start: to.Ptr[int32](10),
							},
							{
								End:   to.Ptr[int32](21),
								Start: to.Ptr[int32](20),
							}},
						Protocol: to.Ptr(armnetwork.ProtocolTypeTCP),
					},
					{
						DestinationIPRanges: []*armnetwork.QosIPRange{
							{
								EndIP:   to.Ptr("<end-ip>"),
								StartIP: to.Ptr("<start-ip>"),
							}},
						DestinationPortRanges: []*armnetwork.QosPortRange{
							{
								End:   to.Ptr[int32](52),
								Start: to.Ptr[int32](51),
							}},
						Markings: []*int32{
							to.Ptr[int32](2)},
						SourceIPRanges: []*armnetwork.QosIPRange{
							{
								EndIP:   to.Ptr("<end-ip>"),
								StartIP: to.Ptr("<start-ip>"),
							}},
						SourcePortRanges: []*armnetwork.QosPortRange{
							{
								End:   to.Ptr[int32](12),
								Start: to.Ptr[int32](11),
							}},
						Protocol: to.Ptr(armnetwork.ProtocolTypeUDP),
					}},
			},
		},
		&armnetwork.DscpConfigurationClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
