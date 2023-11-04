using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureOperatorNexus/VirtualNetworkFunctionCreate.json
// this example is just showing the usage of "NetworkFunctions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkFunctionResource
NetworkFunctionCollection collection = resourceGroupResource.GetNetworkFunctions();

// invoke the operation
string networkFunctionName = "testNf";
NetworkFunctionData data = new NetworkFunctionData(new AzureLocation("eastus"))
{
    Properties = new NetworkFunctionValueWithoutSecrets()
    {
        DeploymentValues = "{\"virtualMachineName\":\"test-VM\",\"extendedLocationName\":\"test-cluster\",\"cpuCores\":4,\"memorySizeGB\":8,\"cloudServicesNetworkAttachment\":{\"attachedNetworkId\":\"test-csnet\",\"ipAllocationMethod\":\"Dynamic\",\"networkAttachmentName\":\"test-cs-vlan\"},\"networkAttachments\":[{\"attachedNetworkId\":\"test-l3vlan\",\"defaultGateway\":\"True\",\"ipAllocationMethod\":\"Dynamic\",\"networkAttachmentName\":\"test-vlan\"}],\"sshPublicKeys\":[{\"keyData\":\"ssh-rsa CMIIIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0TqlveKKlc2MFvEmuXJiLGBsY1t4ML4uiRADGSZlnc+7Ugv3h+MCjkkwOKiOdsNo8k4KSBIG5GcQfKYOOd17AJvqCL6cGQbaLuqv0a64jeDm8oO8/xN/IM0oKw7rMr/2oAJOgIsfeXPkRxWWic9AVIS++H5Qi2r7bUFX+cqFsyUCAwEBBQ==\"}],\"storageProfile\":{\"osDisk\":{\"createOption\":\"Ephemeral\",\"deleteOption\":\"Delete\",\"diskSizeGB\":10}},\"userData\":\"testUserData\",\"adminUsername\":\"testUser\",\"virtioInterface\":\"Transitional\",\"isolateEmulatorThread\":\"False\",\"bootMethod\":\"BIOS\",\"placementHints\":[]}",
        NetworkFunctionDefinitionVersionResourceReference = new OpenDeploymentResourceReference()
        {
            Id = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1"),
        },
        NfviType = NfviType.AzureOperatorNexus,
        NfviId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation"),
        AllowSoftwareUpdate = false,
    },
};
ArmOperation<NetworkFunctionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkFunctionName, data);
NetworkFunctionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFunctionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
