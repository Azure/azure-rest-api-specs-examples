Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.2.1/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_CreateOrUpdate.json
func ExampleFormulasClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewFormulasClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.Formula{
			Location: to.StringPtr("<location>"),
			Properties: &armdevtestlabs.FormulaProperties{
				Description: to.StringPtr("<description>"),
				FormulaContent: &armdevtestlabs.LabVirtualMachineCreationParameter{
					Location: to.StringPtr("<location>"),
					Properties: &armdevtestlabs.LabVirtualMachineCreationParameterProperties{
						AllowClaim: to.BoolPtr(false),
						Artifacts: []*armdevtestlabs.ArtifactInstallProperties{
							{
								ArtifactID: to.StringPtr("<artifact-id>"),
								Parameters: []*armdevtestlabs.ArtifactParameterProperties{},
							}},
						DisallowPublicIPAddress: to.BoolPtr(true),
						GalleryImageReference: &armdevtestlabs.GalleryImageReference{
							Offer:     to.StringPtr("<offer>"),
							OSType:    to.StringPtr("<ostype>"),
							Publisher: to.StringPtr("<publisher>"),
							SKU:       to.StringPtr("<sku>"),
							Version:   to.StringPtr("<version>"),
						},
						IsAuthenticationWithSSHKey: to.BoolPtr(false),
						LabSubnetName:              to.StringPtr("<lab-subnet-name>"),
						LabVirtualNetworkID:        to.StringPtr("<lab-virtual-network-id>"),
						NetworkInterface: &armdevtestlabs.NetworkInterfaceProperties{
							SharedPublicIPAddressConfiguration: &armdevtestlabs.SharedPublicIPAddressConfiguration{
								InboundNatRules: []*armdevtestlabs.InboundNatRule{
									{
										BackendPort:       to.Int32Ptr(22),
										TransportProtocol: armdevtestlabs.TransportProtocol("Tcp").ToPtr(),
									}},
							},
						},
						Notes:       to.StringPtr("<notes>"),
						Size:        to.StringPtr("<size>"),
						StorageType: to.StringPtr("<storage-type>"),
						UserName:    to.StringPtr("<user-name>"),
					},
				},
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
	log.Printf("Response result: %#v\n", res.FormulasClientCreateOrUpdateResult)
}
```
