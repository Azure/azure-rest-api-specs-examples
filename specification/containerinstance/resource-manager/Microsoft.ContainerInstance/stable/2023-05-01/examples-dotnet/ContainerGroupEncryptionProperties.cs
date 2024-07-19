using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerInstance.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerInstance;

// Generated from example definition: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerGroupEncryptionProperties.json
// this example is just showing the usage of "ContainerGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "demo";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerGroupResource
ContainerGroupCollection collection = resourceGroupResource.GetContainerGroups();

// invoke the operation
string containerGroupName = "demo1";
ContainerGroupData data = new ContainerGroupData(new AzureLocation("eastus2"), new ContainerInstanceContainer[]
{
new ContainerInstanceContainer("demo1","nginx",new ContainerResourceRequirements(new ContainerResourceRequestsContent(1.5,1)))
{
Command =
{
},
Ports =
{
new ContainerPort(80)
},
EnvironmentVariables =
{
},
}
}, ContainerInstanceOperatingSystemType.Linux)
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity")] = new UserAssignedIdentity(),
        },
    },
    ImageRegistryCredentials =
    {
    },
    IPAddress = new ContainerGroupIPAddress(new ContainerGroupPort[]
{
new ContainerGroupPort(80)
{
Protocol = ContainerGroupNetworkProtocol.Tcp,
}
}, ContainerGroupIPAddressType.Public),
    EncryptionProperties = new ContainerGroupEncryptionProperties(new Uri("https://testkeyvault.vault.azure.net"), "test-key", "<key version>")
    {
        Identity = "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity",
    },
};
ArmOperation<ContainerGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, containerGroupName, data);
ContainerGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
