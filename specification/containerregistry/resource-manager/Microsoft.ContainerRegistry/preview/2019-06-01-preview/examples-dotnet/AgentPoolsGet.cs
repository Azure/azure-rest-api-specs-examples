using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsGet.json
// this example is just showing the usage of "AgentPools_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryResource created on azure
// for more information of creating ContainerRegistryResource, please refer to the document of ContainerRegistryResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
ResourceIdentifier containerRegistryResourceId = ContainerRegistryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName);
ContainerRegistryResource containerRegistry = client.GetContainerRegistryResource(containerRegistryResourceId);

// get the collection of this ContainerRegistryAgentPoolResource
ContainerRegistryAgentPoolCollection collection = containerRegistry.GetContainerRegistryAgentPools();

// invoke the operation
string agentPoolName = "myAgentPool";
NullableResponse<ContainerRegistryAgentPoolResource> response = await collection.GetIfExistsAsync(agentPoolName);
ContainerRegistryAgentPoolResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerRegistryAgentPoolData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
