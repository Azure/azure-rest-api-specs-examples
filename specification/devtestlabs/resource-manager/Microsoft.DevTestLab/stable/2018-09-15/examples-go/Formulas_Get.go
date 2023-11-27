package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_Get.json
func ExampleFormulasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFormulasClient().Get(ctx, "resourceGroupName", "{labName}", "{formulaName}", &armdevtestlabs.FormulasClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Formula = armdevtestlabs.Formula{
	// 	Name: to.Ptr("{formulaName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/formulas"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/formulas/{formulaName}"),
	// 	Properties: &armdevtestlabs.FormulaProperties{
	// 		Description: to.Ptr("Formula used to create a Linux VM"),
	// 		Author: to.Ptr("user@contoso.com"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-22T19:13:35.992Z"); return t}()),
	// 		FormulaContent: &armdevtestlabs.LabVirtualMachineCreationParameter{
	// 			Location: to.Ptr("{location}"),
	// 			Properties: &armdevtestlabs.LabVirtualMachineCreationParameterProperties{
	// 				AllowClaim: to.Ptr(false),
	// 				Artifacts: []*armdevtestlabs.ArtifactInstallProperties{
	// 					{
	// 						ArtifactID: to.Ptr("/artifactsources/{artifactSourceName}/artifacts/linux-install-nodejs"),
	// 						Parameters: []*armdevtestlabs.ArtifactParameterProperties{
	// 						},
	// 				}},
	// 				DisallowPublicIPAddress: to.Ptr(true),
	// 				GalleryImageReference: &armdevtestlabs.GalleryImageReference{
	// 					Offer: to.Ptr("0001-com-ubuntu-server-groovy"),
	// 					OSType: to.Ptr("Linux"),
	// 					Publisher: to.Ptr("canonical"),
	// 					SKU: to.Ptr("20_10"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				IsAuthenticationWithSSHKey: to.Ptr(false),
	// 				LabSubnetName: to.Ptr("Dtl{labName}Subnet"),
	// 				LabVirtualNetworkID: to.Ptr("/virtualnetworks/dtl{labName}"),
	// 				NetworkInterface: &armdevtestlabs.NetworkInterfaceProperties{
	// 					SharedPublicIPAddressConfiguration: &armdevtestlabs.SharedPublicIPAddressConfiguration{
	// 						InboundNatRules: []*armdevtestlabs.InboundNatRule{
	// 							{
	// 								BackendPort: to.Ptr[int32](22),
	// 								TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
	// 						}},
	// 					},
	// 				},
	// 				Notes: to.Ptr("Ubuntu Server 20.10"),
	// 				Size: to.Ptr("Standard_B1ms"),
	// 				StorageType: to.Ptr("Standard"),
	// 				UserName: to.Ptr("user"),
	// 			},
	// 		},
	// 		OSType: to.Ptr("Linux"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UniqueIdentifier: to.Ptr("7a7d20ab-8a50-46fb-8dc6-7c6c1443a01b"),
	// 	},
	// }
}
