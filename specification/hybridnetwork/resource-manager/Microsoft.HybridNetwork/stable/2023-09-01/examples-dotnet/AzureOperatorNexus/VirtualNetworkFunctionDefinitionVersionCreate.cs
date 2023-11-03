using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureOperatorNexus/VirtualNetworkFunctionDefinitionVersionCreate.json
// this example is just showing the usage of "NetworkFunctionDefinitionVersions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFunctionDefinitionGroupResource created on azure
// for more information of creating NetworkFunctionDefinitionGroupResource, please refer to the document of NetworkFunctionDefinitionGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string publisherName = "TestPublisher";
string networkFunctionDefinitionGroupName = "TestNetworkFunctionDefinitionGroupName";
ResourceIdentifier networkFunctionDefinitionGroupResourceId = NetworkFunctionDefinitionGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, networkFunctionDefinitionGroupName);
NetworkFunctionDefinitionGroupResource networkFunctionDefinitionGroup = client.GetNetworkFunctionDefinitionGroupResource(networkFunctionDefinitionGroupResourceId);

// get the collection of this NetworkFunctionDefinitionVersionResource
NetworkFunctionDefinitionVersionCollection collection = networkFunctionDefinitionGroup.GetNetworkFunctionDefinitionVersions();

// invoke the operation
string networkFunctionDefinitionVersionName = "1.0.0";
NetworkFunctionDefinitionVersionData data = new NetworkFunctionDefinitionVersionData(new AzureLocation("eastus"))
{
    Properties = new VirtualNetworkFunctionDefinitionVersion()
    {
        NetworkFunctionTemplate = new AzureOperatorNexusNetworkFunctionTemplate()
        {
            NetworkFunctionApplications =
            {
            new AzureOperatorNexusNetworkFunctionImageApplication()
            {
            ArtifactProfile = new AzureOperatorNexusImageArtifactProfile()
            {
            ImageArtifactProfile = new ImageArtifactProfile()
            {
            ImageName = "test-image",
            ImageVersion = "1.0.0",
            },
            ArtifactStoreId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"),
            },
            DeployParametersMappingRuleProfile = new AzureOperatorNexusImageDeployMappingRuleProfile()
            {
            ImageMappingRuleUserConfiguration = "",
            ApplicationEnablement = ApplicationEnablement.Unknown,
            },
            Name = "testImageRole",
            DependsOnProfile = new DependsOnProfile()
            {
            InstallDependsOn =
            {
            },
            UninstallDependsOn =
            {
            },
            UpdateDependsOn =
            {
            },
            },
            },new AzureOperatorNexusNetworkFunctionArmTemplateApplication()
            {
            ArtifactProfile = new AzureOperatorNexusArmTemplateArtifactProfile()
            {
            TemplateArtifactProfile = new ArmTemplateArtifactProfile()
            {
            TemplateName = "test-template",
            TemplateVersion = "1.0.0",
            },
            ArtifactStoreId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"),
            },
            DeployParametersMappingRuleProfile = new AzureOperatorNexusArmTemplateDeployMappingRuleProfile()
            {
            TemplateParameters = "{\"virtualMachineName\":\"{deployParameters.virtualMachineName}\",\"extendedLocationName\":\"{deployParameters.extendedLocationName}\",\"cpuCores\":\"{deployParameters.cpuCores}\",\"memorySizeGB\":\"{deployParameters.memorySizeGB}\",\"cloudServicesNetworkAttachment\":\"{deployParameters.cloudServicesNetworkAttachment}\",\"networkAttachments\":\"{deployParameters.networkAttachments}\",\"sshPublicKeys\":\"{deployParameters.sshPublicKeys}\",\"storageProfile\":\"{deployParameters.storageProfile}\",\"isolateEmulatorThread\":\"{deployParameters.isolateEmulatorThread}\",\"virtioInterface\":\"{deployParameters.virtioInterface}\",\"userData\":\"{deployParameters.userData}\",\"adminUsername\":\"{deployParameters.adminUsername}\",\"bootMethod\":\"{deployParameters.bootMethod}\",\"placementHints\":\"{deployParameters.placementHints}\"}",
            ApplicationEnablement = ApplicationEnablement.Unknown,
            },
            Name = "testTemplateRole",
            DependsOnProfile = new DependsOnProfile()
            {
            InstallDependsOn =
            {
            "testImageRole"
            },
            UninstallDependsOn =
            {
            "testImageRole"
            },
            UpdateDependsOn =
            {
            "testImageRole"
            },
            },
            }
            },
        },
        Description = "test NFDV for AzureOperatorNexus",
        DeployParameters = "{\"virtualMachineName\":{\"type\":\"string\"},\"extendedLocationName\":{\"type\":\"string\"},\"cpuCores\":{\"type\":\"int\"},\"memorySizeGB\":{\"type\":\"int\"},\"cloudServicesNetworkAttachment\":{\"type\":\"object\",\"properties\":{\"networkAttachmentName\":{\"type\":\"string\"},\"attachedNetworkId\":{\"type\":\"string\"},\"ipAllocationMethod\":{\"type\":\"string\"},\"ipv4Address\":{\"type\":\"string\"},\"ipv6Address\":{\"type\":\"string\"},\"defaultGateway\":{\"type\":\"string\"}},\"required\":[\"attachedNetworkId\",\"ipAllocationMethod\"]},\"networkAttachments\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"networkAttachmentName\":{\"type\":\"string\"},\"attachedNetworkId\":{\"type\":\"string\"},\"ipAllocationMethod\":{\"type\":\"string\"},\"ipv4Address\":{\"type\":\"string\"},\"ipv6Address\":{\"type\":\"string\"},\"defaultGateway\":{\"type\":\"string\"}},\"required\":[\"attachedNetworkId\",\"ipAllocationMethod\"]}},\"storageProfile\":{\"type\":\"object\",\"properties\":{\"osDisk\":{\"type\":\"object\",\"properties\":{\"createOption\":{\"type\":\"string\"},\"deleteOption\":{\"type\":\"string\"},\"diskSizeGB\":{\"type\":\"integer\"}},\"required\":[\"diskSizeGB\"]}},\"required\":[\"osDisk\"]},\"sshPublicKeys\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"keyData\":{\"type\":\"string\"}},\"required\":[\"keyData\"]}},\"userData\":{\"type\":\"string\"},\"adminUsername\":{\"type\":\"string\"},\"bootMethod\":{\"type\":\"string\",\"default\":\"UEFI\",\"enum\":[\"UEFI\",\"BIOS\"]},\"isolateEmulatorThread\":{\"type\":\"string\"},\"virtioInterface\":{\"type\":\"string\"},\"placementHints\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"hintType\":{\"type\":\"string\",\"enum\":[\"Affinity\",\"AntiAffinity\"]},\"resourceId\":{\"type\":\"string\"},\"schedulingExecution\":{\"type\":\"string\",\"enum\":[\"Soft\",\"Hard\"]},\"scope\":{\"type\":\"string\"}},\"required\":[\"hintType\",\"schedulingExecution\",\"resourceId\",\"scope\"]}}}",
    },
};
ArmOperation<NetworkFunctionDefinitionVersionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkFunctionDefinitionVersionName, data);
NetworkFunctionDefinitionVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFunctionDefinitionVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
