using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionVersionListByNetworkFunctionDefinitionGroup.json
// this example is just showing the usage of "NetworkFunctionDefinitionVersions_ListByNetworkFunctionDefinitionGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFunctionDefinitionGroupResource created on azure
// for more information of creating NetworkFunctionDefinitionGroupResource, please refer to the document of NetworkFunctionDefinitionGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string publisherName = "TestPublisher";
string networkFunctionDefinitionGroupName = "TestNetworkFunctionDefinitionGroupNameName";
ResourceIdentifier networkFunctionDefinitionGroupResourceId = NetworkFunctionDefinitionGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, networkFunctionDefinitionGroupName);
NetworkFunctionDefinitionGroupResource networkFunctionDefinitionGroup = client.GetNetworkFunctionDefinitionGroupResource(networkFunctionDefinitionGroupResourceId);

// get the collection of this NetworkFunctionDefinitionVersionResource
NetworkFunctionDefinitionVersionCollection collection = networkFunctionDefinitionGroup.GetNetworkFunctionDefinitionVersions();

// invoke the operation and iterate over the result
await foreach (NetworkFunctionDefinitionVersionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkFunctionDefinitionVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
