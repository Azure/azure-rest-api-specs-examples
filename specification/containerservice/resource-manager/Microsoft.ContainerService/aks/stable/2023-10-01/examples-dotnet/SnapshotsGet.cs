using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-10-01/examples/SnapshotsGet.json
// this example is just showing the usage of "Snapshots_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AgentPoolSnapshotResource
AgentPoolSnapshotCollection collection = resourceGroupResource.GetAgentPoolSnapshots();

// invoke the operation
string resourceName = "snapshot1";
NullableResponse<AgentPoolSnapshotResource> response = await collection.GetIfExistsAsync(resourceName);
AgentPoolSnapshotResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AgentPoolSnapshotData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
