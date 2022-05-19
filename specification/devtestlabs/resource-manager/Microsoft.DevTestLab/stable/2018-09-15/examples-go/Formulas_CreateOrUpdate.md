Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv1.0.0/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_CreateOrUpdate.json
func ExampleFormulasClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewFormulasClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resourceGroupName",
		"{labName}",
		"{formulaName}",
		armdevtestlabs.Formula{
			Location: to.Ptr("{location}"),
			Properties: &armdevtestlabs.FormulaProperties{
				Description: to.Ptr("Formula using a Linux base"),
				FormulaContent: &armdevtestlabs.LabVirtualMachineCreationParameter{
					Location: to.Ptr("{location}"),
					Properties: &armdevtestlabs.LabVirtualMachineCreationParameterProperties{
						AllowClaim: to.Ptr(false),
						Artifacts: []*armdevtestlabs.ArtifactInstallProperties{
							{
								ArtifactID: to.Ptr("/artifactsources/{artifactSourceName}/artifacts/linux-install-nodejs"),
								Parameters: []*armdevtestlabs.ArtifactParameterProperties{},
							}},
						DisallowPublicIPAddress: to.Ptr(true),
						GalleryImageReference: &armdevtestlabs.GalleryImageReference{
							Offer:     to.Ptr("0001-com-ubuntu-server-groovy"),
							OSType:    to.Ptr("Linux"),
							Publisher: to.Ptr("canonical"),
							SKU:       to.Ptr("20_10"),
							Version:   to.Ptr("latest"),
						},
						IsAuthenticationWithSSHKey: to.Ptr(false),
						LabSubnetName:              to.Ptr("Dtl{labName}Subnet"),
						LabVirtualNetworkID:        to.Ptr("/virtualnetworks/dtl{labName}"),
						NetworkInterface: &armdevtestlabs.NetworkInterfaceProperties{
							SharedPublicIPAddressConfiguration: &armdevtestlabs.SharedPublicIPAddressConfiguration{
								InboundNatRules: []*armdevtestlabs.InboundNatRule{
									{
										BackendPort:       to.Ptr[int32](22),
										TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
									}},
							},
						},
						Notes:       to.Ptr("Ubuntu Server 20.10"),
						Size:        to.Ptr("Standard_B1ms"),
						StorageType: to.Ptr("Standard"),
						UserName:    to.Ptr("user"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
