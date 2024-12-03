using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/Racks_ListByResourceGroup.json
// this example is just showing the usage of "Racks_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkCloudRackResource
NetworkCloudRackCollection collection = resourceGroupResource.GetNetworkCloudRacks();

// invoke the operation and iterate over the result
await foreach (NetworkCloudRackResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkCloudRackData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
