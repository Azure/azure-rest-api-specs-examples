using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ListByResourceGroup.json
// this example is just showing the usage of "IntegrationServiceEnvironments_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "f34b22a3-2202-4fb1-b040-1332bd928c84";
string resourceGroup = "testResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroup);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IntegrationServiceEnvironmentResource
IntegrationServiceEnvironmentCollection collection = resourceGroupResource.GetIntegrationServiceEnvironments();

// invoke the operation and iterate over the result
await foreach (IntegrationServiceEnvironmentResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    IntegrationServiceEnvironmentData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
