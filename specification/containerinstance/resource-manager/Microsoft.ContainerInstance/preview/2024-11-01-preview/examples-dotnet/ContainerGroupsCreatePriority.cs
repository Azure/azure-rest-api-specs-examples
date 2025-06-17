using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerInstance.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerInstance;

// Generated from example definition: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-11-01-preview/examples/ContainerGroupsCreatePriority.json
// this example is just showing the usage of "ContainerGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "demo";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerGroupResource
ContainerGroupCollection collection = resourceGroupResource.GetContainerGroups();

// invoke the operation
string containerGroupName = "demo1";
ContainerGroupData data = new ContainerGroupData(new AzureLocation("eastus"), new ContainerInstanceContainer[]
{
new ContainerInstanceContainer("test-container-001", "alpine:latest", new ContainerResourceRequirements(new ContainerResourceRequestsContent(1, 1)))
{
Command = {"/bin/sh", "-c", "sleep 10"},
}
}, ContainerInstanceOperatingSystemType.Linux)
{
    RestartPolicy = ContainerGroupRestartPolicy.Never,
    Sku = ContainerGroupSku.Standard,
    Priority = ContainerGroupPriority.Spot,
};
ArmOperation<ContainerGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, containerGroupName, data);
ContainerGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
