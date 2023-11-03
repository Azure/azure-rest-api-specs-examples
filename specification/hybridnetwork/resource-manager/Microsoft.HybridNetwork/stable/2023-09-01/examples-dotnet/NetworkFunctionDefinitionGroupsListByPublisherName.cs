using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionGroupsListByPublisherName.json
// this example is just showing the usage of "NetworkFunctionDefinitionGroups_ListByPublisher" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PublisherResource created on azure
// for more information of creating PublisherResource, please refer to the document of PublisherResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string publisherName = "TestPublisher";
ResourceIdentifier publisherResourceId = PublisherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName);
PublisherResource publisher = client.GetPublisherResource(publisherResourceId);

// get the collection of this NetworkFunctionDefinitionGroupResource
NetworkFunctionDefinitionGroupCollection collection = publisher.GetNetworkFunctionDefinitionGroups();

// invoke the operation and iterate over the result
await foreach (NetworkFunctionDefinitionGroupResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkFunctionDefinitionGroupData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
