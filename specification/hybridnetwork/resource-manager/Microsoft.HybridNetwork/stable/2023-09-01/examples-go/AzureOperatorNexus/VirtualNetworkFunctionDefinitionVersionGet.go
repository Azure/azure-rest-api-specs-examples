package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureOperatorNexus/VirtualNetworkFunctionDefinitionVersionGet.json
func ExampleNetworkFunctionDefinitionVersionsClient_Get_getNetworkFunctionDefinitionVersionResourceForAzureOperatorNexusVnf() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkFunctionDefinitionVersionsClient().Get(ctx, "rg", "TestPublisher", "TestNetworkFunctionDefinitionGroupName", "1.0.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFunctionDefinitionVersion = armhybridnetwork.NetworkFunctionDefinitionVersion{
	// 	Name: to.Ptr("TestVersion"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/publishers/networkFunctionDefinitionGroups/networkFunctionDefinitionVersions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/TestPublisher/networkFunctionDefinitionGroups/TestNetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.0"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.VirtualNetworkFunctionDefinitionVersion{
	// 		Description: to.Ptr("test NFDV for AzureOperatorNexus"),
	// 		DeployParameters: to.Ptr("{\"virtualMachineName\":{\"type\":\"string\"},\"extendedLocationName\":{\"type\":\"string\"},\"cpuCores\":{\"type\":\"int\"},\"memorySizeGB\":{\"type\":\"int\"},\"cloudServicesNetworkAttachment\":{\"type\":\"object\",\"properties\":{\"networkAttachmentName\":{\"type\":\"string\"},\"attachedNetworkId\":{\"type\":\"string\"},\"ipAllocationMethod\":{\"type\":\"string\"},\"ipv4Address\":{\"type\":\"string\"},\"ipv6Address\":{\"type\":\"string\"},\"defaultGateway\":{\"type\":\"string\"}},\"required\":[\"attachedNetworkId\",\"ipAllocationMethod\"]},\"networkAttachments\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"networkAttachmentName\":{\"type\":\"string\"},\"attachedNetworkId\":{\"type\":\"string\"},\"ipAllocationMethod\":{\"type\":\"string\"},\"ipv4Address\":{\"type\":\"string\"},\"ipv6Address\":{\"type\":\"string\"},\"defaultGateway\":{\"type\":\"string\"}},\"required\":[\"attachedNetworkId\",\"ipAllocationMethod\"]}},\"storageProfile\":{\"type\":\"object\",\"properties\":{\"osDisk\":{\"type\":\"object\",\"properties\":{\"createOption\":{\"type\":\"string\"},\"deleteOption\":{\"type\":\"string\"},\"diskSizeGB\":{\"type\":\"integer\"}},\"required\":[\"diskSizeGB\"]}},\"required\":[\"osDisk\"]},\"sshPublicKeys\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"keyData\":{\"type\":\"string\"}},\"required\":[\"keyData\"]}},\"userData\":{\"type\":\"string\"},\"adminUsername\":{\"type\":\"string\"},\"bootMethod\":{\"type\":\"string\",\"default\":\"UEFI\",\"enum\":[\"UEFI\",\"BIOS\"]},\"isolateEmulatorThread\":{\"type\":\"string\"},\"virtioInterface\":{\"type\":\"string\"},\"placementHints\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"hintType\":{\"type\":\"string\",\"enum\":[\"Affinity\",\"AntiAffinity\"]},\"resourceId\":{\"type\":\"string\"},\"schedulingExecution\":{\"type\":\"string\",\"enum\":[\"Soft\",\"Hard\"]},\"scope\":{\"type\":\"string\"}},\"required\":[\"hintType\",\"schedulingExecution\",\"resourceId\",\"scope\"]}}}"),
	// 		NetworkFunctionType: to.Ptr(armhybridnetwork.NetworkFunctionTypeVirtualNetworkFunction),
	// 		VersionState: to.Ptr(armhybridnetwork.VersionStatePreview),
	// 		NetworkFunctionTemplate: &armhybridnetwork.AzureOperatorNexusNetworkFunctionTemplate{
	// 			NfviType: to.Ptr(armhybridnetwork.VirtualNetworkFunctionNFVITypeAzureOperatorNexus),
	// 			NetworkFunctionApplications: []armhybridnetwork.AzureOperatorNexusNetworkFunctionApplicationClassification{
	// 				&armhybridnetwork.AzureOperatorNexusNetworkFunctionImageApplication{
	// 					Name: to.Ptr("testImageRole"),
	// 					DependsOnProfile: &armhybridnetwork.DependsOnProfile{
	// 						InstallDependsOn: []*string{
	// 						},
	// 						UninstallDependsOn: []*string{
	// 						},
	// 						UpdateDependsOn: []*string{
	// 						},
	// 					},
	// 					ArtifactType: to.Ptr(armhybridnetwork.AzureOperatorNexusArtifactTypeImageFile),
	// 					ArtifactProfile: &armhybridnetwork.AzureOperatorNexusImageArtifactProfile{
	// 						ArtifactStore: &armhybridnetwork.ReferencedResource{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"),
	// 						},
	// 						ImageArtifactProfile: &armhybridnetwork.ImageArtifactProfile{
	// 							ImageName: to.Ptr("test-image"),
	// 							ImageVersion: to.Ptr("1.0.0"),
	// 						},
	// 					},
	// 					DeployParametersMappingRuleProfile: &armhybridnetwork.AzureOperatorNexusImageDeployMappingRuleProfile{
	// 						ApplicationEnablement: to.Ptr(armhybridnetwork.ApplicationEnablementUnknown),
	// 						ImageMappingRuleProfile: &armhybridnetwork.ImageMappingRuleProfile{
	// 							UserConfiguration: to.Ptr(""),
	// 						},
	// 					},
	// 				},
	// 				&armhybridnetwork.AzureOperatorNexusNetworkFunctionArmTemplateApplication{
	// 					Name: to.Ptr("testTemplateRole"),
	// 					DependsOnProfile: &armhybridnetwork.DependsOnProfile{
	// 						InstallDependsOn: []*string{
	// 							to.Ptr("testImageRole")},
	// 							UninstallDependsOn: []*string{
	// 								to.Ptr("testImageRole")},
	// 								UpdateDependsOn: []*string{
	// 									to.Ptr("testImageRole")},
	// 								},
	// 								ArtifactType: to.Ptr(armhybridnetwork.AzureOperatorNexusArtifactTypeArmTemplate),
	// 								ArtifactProfile: &armhybridnetwork.AzureOperatorNexusArmTemplateArtifactProfile{
	// 									ArtifactStore: &armhybridnetwork.ReferencedResource{
	// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"),
	// 									},
	// 									TemplateArtifactProfile: &armhybridnetwork.ArmTemplateArtifactProfile{
	// 										TemplateName: to.Ptr("test-template"),
	// 										TemplateVersion: to.Ptr("1.0.0"),
	// 									},
	// 								},
	// 								DeployParametersMappingRuleProfile: &armhybridnetwork.AzureOperatorNexusArmTemplateDeployMappingRuleProfile{
	// 									ApplicationEnablement: to.Ptr(armhybridnetwork.ApplicationEnablementUnknown),
	// 									TemplateMappingRuleProfile: &armhybridnetwork.ArmTemplateMappingRuleProfile{
	// 										TemplateParameters: to.Ptr("{\"virtualMachineName\":\"{deployParameters.virtualMachineName}\",\"extendedLocationName\":\"{deployParameters.extendedLocationName}\",\"cpuCores\":\"{deployParameters.cpuCores}\",\"memorySizeGB\":\"{deployParameters.memorySizeGB}\",\"cloudServicesNetworkAttachment\":\"{deployParameters.cloudServicesNetworkAttachment}\",\"networkAttachments\":\"{deployParameters.networkAttachments}\",\"sshPublicKeys\":\"{deployParameters.sshPublicKeys}\",\"storageProfile\":\"{deployParameters.storageProfile}\",\"isolateEmulatorThread\":\"{deployParameters.isolateEmulatorThread}\",\"virtioInterface\":\"{deployParameters.virtioInterface}\",\"userData\":\"{deployParameters.userData}\",\"adminUsername\":\"{deployParameters.adminUsername}\",\"bootMethod\":\"{deployParameters.bootMethod}\",\"placementHints\":\"{deployParameters.placementHints}\"}"),
	// 									},
	// 								},
	// 						}},
	// 					},
	// 				},
	// 			}
}
