Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DscpConfigurationCreate.json
func ExampleDscpConfigurationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewDscpConfigurationClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<dscp-configuration-name>",
		armnetwork.DscpConfiguration{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.DscpConfigurationPropertiesFormat{
				QosDefinitionCollection: []*armnetwork.QosDefinition{
					{
						DestinationIPRanges: []*armnetwork.QosIPRange{
							{
								EndIP:   to.StringPtr("<end-ip>"),
								StartIP: to.StringPtr("<start-ip>"),
							}},
						DestinationPortRanges: []*armnetwork.QosPortRange{
							{
								End:   to.Int32Ptr(15),
								Start: to.Int32Ptr(15),
							}},
						Markings: []*int32{
							to.Int32Ptr(1)},
						SourceIPRanges: []*armnetwork.QosIPRange{
							{
								EndIP:   to.StringPtr("<end-ip>"),
								StartIP: to.StringPtr("<start-ip>"),
							}},
						SourcePortRanges: []*armnetwork.QosPortRange{
							{
								End:   to.Int32Ptr(11),
								Start: to.Int32Ptr(10),
							},
							{
								End:   to.Int32Ptr(21),
								Start: to.Int32Ptr(20),
							}},
						Protocol: armnetwork.ProtocolType("Tcp").ToPtr(),
					},
					{
						DestinationIPRanges: []*armnetwork.QosIPRange{
							{
								EndIP:   to.StringPtr("<end-ip>"),
								StartIP: to.StringPtr("<start-ip>"),
							}},
						DestinationPortRanges: []*armnetwork.QosPortRange{
							{
								End:   to.Int32Ptr(52),
								Start: to.Int32Ptr(51),
							}},
						Markings: []*int32{
							to.Int32Ptr(2)},
						SourceIPRanges: []*armnetwork.QosIPRange{
							{
								EndIP:   to.StringPtr("<end-ip>"),
								StartIP: to.StringPtr("<start-ip>"),
							}},
						SourcePortRanges: []*armnetwork.QosPortRange{
							{
								End:   to.Int32Ptr(12),
								Start: to.Int32Ptr(11),
							}},
						Protocol: armnetwork.ProtocolType("Udp").ToPtr(),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DscpConfigurationClientCreateOrUpdateResult)
}
```
