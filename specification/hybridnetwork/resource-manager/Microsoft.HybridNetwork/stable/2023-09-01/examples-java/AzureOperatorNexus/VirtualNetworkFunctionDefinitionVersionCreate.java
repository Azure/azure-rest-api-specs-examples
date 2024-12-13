
import com.azure.resourcemanager.hybridnetwork.models.ApplicationEnablement;
import com.azure.resourcemanager.hybridnetwork.models.ArmTemplateArtifactProfile;
import com.azure.resourcemanager.hybridnetwork.models.ArmTemplateMappingRuleProfile;
import com.azure.resourcemanager.hybridnetwork.models.AzureOperatorNexusArmTemplateArtifactProfile;
import com.azure.resourcemanager.hybridnetwork.models.AzureOperatorNexusArmTemplateDeployMappingRuleProfile;
import com.azure.resourcemanager.hybridnetwork.models.AzureOperatorNexusImageArtifactProfile;
import com.azure.resourcemanager.hybridnetwork.models.AzureOperatorNexusImageDeployMappingRuleProfile;
import com.azure.resourcemanager.hybridnetwork.models.AzureOperatorNexusNetworkFunctionArmTemplateApplication;
import com.azure.resourcemanager.hybridnetwork.models.AzureOperatorNexusNetworkFunctionImageApplication;
import com.azure.resourcemanager.hybridnetwork.models.AzureOperatorNexusNetworkFunctionTemplate;
import com.azure.resourcemanager.hybridnetwork.models.DependsOnProfile;
import com.azure.resourcemanager.hybridnetwork.models.ImageArtifactProfile;
import com.azure.resourcemanager.hybridnetwork.models.ImageMappingRuleProfile;
import com.azure.resourcemanager.hybridnetwork.models.ReferencedResource;
import com.azure.resourcemanager.hybridnetwork.models.VirtualNetworkFunctionDefinitionVersion;
import java.util.Arrays;

/**
 * Samples for NetworkFunctionDefinitionVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * AzureOperatorNexus/VirtualNetworkFunctionDefinitionVersionCreate.json
     */
    /**
     * Sample code: Create or update a network function definition version resource for AzureOperatorNexus VNF.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateANetworkFunctionDefinitionVersionResourceForAzureOperatorNexusVNF(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionVersions().define("1.0.0").withRegion("eastus")
            .withExistingNetworkFunctionDefinitionGroup("rg", "TestPublisher", "TestNetworkFunctionDefinitionGroupName")
            .withProperties(new VirtualNetworkFunctionDefinitionVersion()
                .withDescription("test NFDV for AzureOperatorNexus")
                .withDeployParameters(
                    "{\"virtualMachineName\":{\"type\":\"string\"},\"extendedLocationName\":{\"type\":\"string\"},\"cpuCores\":{\"type\":\"int\"},\"memorySizeGB\":{\"type\":\"int\"},\"cloudServicesNetworkAttachment\":{\"type\":\"object\",\"properties\":{\"networkAttachmentName\":{\"type\":\"string\"},\"attachedNetworkId\":{\"type\":\"string\"},\"ipAllocationMethod\":{\"type\":\"string\"},\"ipv4Address\":{\"type\":\"string\"},\"ipv6Address\":{\"type\":\"string\"},\"defaultGateway\":{\"type\":\"string\"}},\"required\":[\"attachedNetworkId\",\"ipAllocationMethod\"]},\"networkAttachments\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"networkAttachmentName\":{\"type\":\"string\"},\"attachedNetworkId\":{\"type\":\"string\"},\"ipAllocationMethod\":{\"type\":\"string\"},\"ipv4Address\":{\"type\":\"string\"},\"ipv6Address\":{\"type\":\"string\"},\"defaultGateway\":{\"type\":\"string\"}},\"required\":[\"attachedNetworkId\",\"ipAllocationMethod\"]}},\"storageProfile\":{\"type\":\"object\",\"properties\":{\"osDisk\":{\"type\":\"object\",\"properties\":{\"createOption\":{\"type\":\"string\"},\"deleteOption\":{\"type\":\"string\"},\"diskSizeGB\":{\"type\":\"integer\"}},\"required\":[\"diskSizeGB\"]}},\"required\":[\"osDisk\"]},\"sshPublicKeys\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"keyData\":{\"type\":\"string\"}},\"required\":[\"keyData\"]}},\"userData\":{\"type\":\"string\"},\"adminUsername\":{\"type\":\"string\"},\"bootMethod\":{\"type\":\"string\",\"default\":\"UEFI\",\"enum\":[\"UEFI\",\"BIOS\"]},\"isolateEmulatorThread\":{\"type\":\"string\"},\"virtioInterface\":{\"type\":\"string\"},\"placementHints\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"hintType\":{\"type\":\"string\",\"enum\":[\"Affinity\",\"AntiAffinity\"]},\"resourceId\":{\"type\":\"string\"},\"schedulingExecution\":{\"type\":\"string\",\"enum\":[\"Soft\",\"Hard\"]},\"scope\":{\"type\":\"string\"}},\"required\":[\"hintType\",\"schedulingExecution\",\"resourceId\",\"scope\"]}}}")
                .withNetworkFunctionTemplate(
                    new AzureOperatorNexusNetworkFunctionTemplate().withNetworkFunctionApplications(Arrays.asList(
                        new AzureOperatorNexusNetworkFunctionImageApplication().withName("testImageRole")
                            .withDependsOnProfile(new DependsOnProfile().withInstallDependsOn(Arrays.asList())
                                .withUninstallDependsOn(Arrays.asList()).withUpdateDependsOn(Arrays.asList()))
                            .withArtifactProfile(new AzureOperatorNexusImageArtifactProfile()
                                .withArtifactStore(new ReferencedResource().withId(
                                    "/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"))
                                .withImageArtifactProfile(
                                    new ImageArtifactProfile().withImageName("test-image").withImageVersion("1.0.0")))
                            .withDeployParametersMappingRuleProfile(
                                new AzureOperatorNexusImageDeployMappingRuleProfile()
                                    .withApplicationEnablement(ApplicationEnablement.UNKNOWN)
                                    .withImageMappingRuleProfile(
                                        new ImageMappingRuleProfile().withUserConfiguration(""))),
                        new AzureOperatorNexusNetworkFunctionArmTemplateApplication().withName("testTemplateRole")
                            .withDependsOnProfile(
                                new DependsOnProfile().withInstallDependsOn(Arrays.asList("testImageRole"))
                                    .withUninstallDependsOn(Arrays.asList("testImageRole"))
                                    .withUpdateDependsOn(Arrays.asList("testImageRole")))
                            .withArtifactProfile(new AzureOperatorNexusArmTemplateArtifactProfile()
                                .withArtifactStore(new ReferencedResource().withId(
                                    "/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"))
                                .withTemplateArtifactProfile(new ArmTemplateArtifactProfile()
                                    .withTemplateName("test-template").withTemplateVersion("1.0.0")))
                            .withDeployParametersMappingRuleProfile(
                                new AzureOperatorNexusArmTemplateDeployMappingRuleProfile()
                                    .withApplicationEnablement(ApplicationEnablement.UNKNOWN)
                                    .withTemplateMappingRuleProfile(
                                        new ArmTemplateMappingRuleProfile().withTemplateParameters(
                                            "{\"virtualMachineName\":\"{deployParameters.virtualMachineName}\",\"extendedLocationName\":\"{deployParameters.extendedLocationName}\",\"cpuCores\":\"{deployParameters.cpuCores}\",\"memorySizeGB\":\"{deployParameters.memorySizeGB}\",\"cloudServicesNetworkAttachment\":\"{deployParameters.cloudServicesNetworkAttachment}\",\"networkAttachments\":\"{deployParameters.networkAttachments}\",\"sshPublicKeys\":\"{deployParameters.sshPublicKeys}\",\"storageProfile\":\"{deployParameters.storageProfile}\",\"isolateEmulatorThread\":\"{deployParameters.isolateEmulatorThread}\",\"virtioInterface\":\"{deployParameters.virtioInterface}\",\"userData\":\"{deployParameters.userData}\",\"adminUsername\":\"{deployParameters.adminUsername}\",\"bootMethod\":\"{deployParameters.bootMethod}\",\"placementHints\":\"{deployParameters.placementHints}\"}")))))))
            .create();
    }
}
