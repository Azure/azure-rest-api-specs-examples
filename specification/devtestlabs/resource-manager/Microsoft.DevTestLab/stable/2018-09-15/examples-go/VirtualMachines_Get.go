package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Get.json
func ExampleVirtualMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().Get(ctx, "resourceGroupName", "{labName}", "{vmName}", &armdevtestlabs.VirtualMachinesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LabVirtualMachine = armdevtestlabs.LabVirtualMachine{
	// 	Name: to.Ptr("{vmName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/{vmName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.LabVirtualMachineProperties{
	// 		AllowClaim: to.Ptr(true),
	// 		ArtifactDeploymentStatus: &armdevtestlabs.ArtifactDeploymentStatusProperties{
	// 			ArtifactsApplied: to.Ptr[int32](0),
	// 			TotalArtifacts: to.Ptr[int32](0),
	// 		},
	// 		ComputeID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{labName}-{vmName}-{randomSuffix}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
	// 		CreatedByUser: to.Ptr(""),
	// 		CreatedByUserID: to.Ptr(""),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T16:53:02.4830866-07:00"); return t}()),
	// 		DataDiskParameters: []*armdevtestlabs.DataDiskProperties{
	// 		},
	// 		DisallowPublicIPAddress: to.Ptr(true),
	// 		GalleryImageReference: &armdevtestlabs.GalleryImageReference{
	// 			Offer: to.Ptr("UbuntuServer"),
	// 			OSType: to.Ptr("Linux"),
	// 			Publisher: to.Ptr("Canonical"),
	// 			SKU: to.Ptr("16.04-LTS"),
	// 			Version: to.Ptr("Latest"),
	// 		},
	// 		LabSubnetName: to.Ptr("{virtualNetworkName}Subnet"),
	// 		LabVirtualNetworkID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}"),
	// 		NetworkInterface: &armdevtestlabs.NetworkInterfaceProperties{
	// 		},
	// 		OSType: to.Ptr("Linux"),
	// 		OwnerObjectID: to.Ptr(""),
	// 		OwnerUserPrincipalName: to.Ptr(""),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Size: to.Ptr("Standard_A2_v2"),
	// 		StorageType: to.Ptr("Standard"),
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 		UserName: to.Ptr("{userName}"),
	// 		VirtualMachineCreationSource: to.Ptr(armdevtestlabs.VirtualMachineCreationSourceFromGalleryImage),
	// 	},
	// }
}
