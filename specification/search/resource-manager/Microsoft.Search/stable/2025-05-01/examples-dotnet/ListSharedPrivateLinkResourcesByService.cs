using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Search.Models;
using Azure.ResourceManager.Search;

// Generated from example definition: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/ListSharedPrivateLinkResourcesByService.json
// this example is just showing the usage of "SharedPrivateLinkResources_ListByService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SearchServiceResource created on azure
// for more information of creating SearchServiceResource, please refer to the document of SearchServiceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string searchServiceName = "mysearchservice";
ResourceIdentifier searchServiceResourceId = SearchServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, searchServiceName);
SearchServiceResource searchService = client.GetSearchServiceResource(searchServiceResourceId);

// get the collection of this SharedSearchServicePrivateLinkResource
SharedSearchServicePrivateLinkResourceCollection collection = searchService.GetSharedSearchServicePrivateLinkResources();

// invoke the operation and iterate over the result
await foreach (SharedSearchServicePrivateLinkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SharedSearchServicePrivateLinkResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
