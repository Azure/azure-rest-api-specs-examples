using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ResourceGraph;

// Generated from example definition: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2024-04-01/examples/GraphQueryAdd.json
// this example is just showing the usage of "GraphQuery_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "024e2271-06fa-46b6-9079-f1ed3c7b070e";
string resourceGroupName = "my-resource-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ResourceGraphQueryResource
ResourceGraphQueryCollection collection = resourceGroupResource.GetResourceGraphQueries();

// invoke the operation
string resourceName = "MyDockerVMs";
ResourceGraphQueryData data = new ResourceGraphQueryData(default)
{
    Description = "Docker VMs in PROD",
    Query = "where isnotnull(tags['Prod']) and properties.extensions[0].Name == 'docker'",
    Tags = { },
};
ArmOperation<ResourceGraphQueryResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
ResourceGraphQueryResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceGraphQueryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
