package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureCore/VirtualNetworkFunctionCreate.json
func ExampleNetworkFunctionsClient_BeginCreateOrUpdate_createVirtualNetworkFunctionResourceOnAzureCore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFunctionsClient().BeginCreateOrUpdate(ctx, "rg", "testNf", armhybridnetwork.NetworkFunction{
		Location: to.Ptr("eastus"),
		Properties: &armhybridnetwork.NetworkFunctionValueWithoutSecrets{
			AllowSoftwareUpdate: to.Ptr(false),
			ConfigurationType:   to.Ptr(armhybridnetwork.NetworkFunctionConfigurationTypeOpen),
			NetworkFunctionDefinitionVersionResourceReference: &armhybridnetwork.OpenDeploymentResourceReference{
				IDType: to.Ptr(armhybridnetwork.IDTypeOpen),
				ID:     to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1"),
			},
			NfviID:           to.Ptr("/subscriptions/subid/resourceGroups/testResourceGroup"),
			NfviType:         to.Ptr(armhybridnetwork.NFVITypeAzureCore),
			DeploymentValues: to.Ptr("{\"virtualMachineName\":\"test-VM\",\"cpuCores\":4,\"memorySizeGB\":8,\"cloudServicesNetworkAttachment\":{\"attachedNetworkId\":\"test-csnet\",\"ipAllocationMethod\":\"Dynamic\",\"networkAttachmentName\":\"test-cs-vlan\"},\"networkAttachments\":[{\"attachedNetworkId\":\"test-l3vlan\",\"defaultGateway\":\"True\",\"ipAllocationMethod\":\"Dynamic\",\"networkAttachmentName\":\"test-vlan\"}],\"sshPublicKeys\":[{\"keyData\":\"ssh-rsa CMIIIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0TqlveKKlc2MFvEmuXJiLGBsY1t4ML4uiRADGSZlnc+7Ugv3h+MCjkkwOKiOdsNo8k4KSBIG5GcQfKYOOd17AJvqCL6cGQbaLuqv0a64jeDm8oO8/xN/IM0oKw7rMr/2oAJOgIsfeXPkRxWWic9AVIS++H5Qi2r7bUFX+cqFsyUCAwEBBQ==\"}],\"storageProfile\":{\"osDisk\":{\"createOption\":\"Ephemeral\",\"deleteOption\":\"Delete\",\"diskSizeGB\":10}},\"userData\":\"testUserData\",\"adminUsername\":\"testUser\",\"virtioInterface\":\"Transitional\",\"isolateEmulatorThread\":\"False\",\"bootMethod\":\"BIOS\",\"placementHints\":[]}"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFunction = armhybridnetwork.NetworkFunction{
	// 	Name: to.Ptr("testNf"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.NetworkFunctionValueWithoutSecrets{
	// 		AllowSoftwareUpdate: to.Ptr(false),
	// 		ConfigurationType: to.Ptr(armhybridnetwork.NetworkFunctionConfigurationTypeOpen),
	// 		NetworkFunctionDefinitionGroupName: to.Ptr("TestNetworkFunctionDefinitionGroupName"),
	// 		NetworkFunctionDefinitionOfferingLocation: to.Ptr("eastus"),
	// 		NetworkFunctionDefinitionVersion: to.Ptr("1.0.0"),
	// 		NetworkFunctionDefinitionVersionResourceReference: &armhybridnetwork.OpenDeploymentResourceReference{
	// 			IDType: to.Ptr(armhybridnetwork.IDTypeOpen),
	// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1"),
	// 		},
	// 		NfviID: to.Ptr("/subscriptions/subid/resourceGroups/testResourceGroup"),
	// 		NfviType: to.Ptr(armhybridnetwork.NFVITypeAzureCore),
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 		PublisherName: to.Ptr("TestPublisher"),
	// 		PublisherScope: to.Ptr(armhybridnetwork.PublisherScopePrivate),
	// 		DeploymentValues: to.Ptr("{\"virtualMachineName\":\"test-VM\",\"cpuCores\":4,\"memorySizeGB\":8,\"cloudServicesNetworkAttachment\":{\"attachedNetworkId\":\"test-csnet\",\"ipAllocationMethod\":\"Dynamic\",\"networkAttachmentName\":\"test-cs-vlan\"},\"networkAttachments\":[{\"attachedNetworkId\":\"test-l3vlan\",\"defaultGateway\":\"True\",\"ipAllocationMethod\":\"Dynamic\",\"networkAttachmentName\":\"test-vlan\"}],\"sshPublicKeys\":[{\"keyData\":\"ssh-rsa CMIIIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0TqlveKKlc2MFvEmuXJiLGBsY1t4ML4uiRADGSZlnc+7Ugv3h+MCjkkwOKiOdsNo8k4KSBIG5GcQfKYOOd17AJvqCL6cGQbaLuqv0a64jeDm8oO8/xN/IM0oKw7rMr/2oAJOgIsfeXPkRxWWic9AVIS++H5Qi2r7bUFX+cqFsyUCAwEBBQ==\"}],\"storageProfile\":{\"osDisk\":{\"createOption\":\"Ephemeral\",\"deleteOption\":\"Delete\",\"diskSizeGB\":10}},\"userData\":\"testUserData\",\"adminUsername\":\"testUser\",\"virtioInterface\":\"Transitional\",\"isolateEmulatorThread\":\"False\",\"bootMethod\":\"BIOS\",\"placementHints\":[]}"),
	// 	},
	// }
}
