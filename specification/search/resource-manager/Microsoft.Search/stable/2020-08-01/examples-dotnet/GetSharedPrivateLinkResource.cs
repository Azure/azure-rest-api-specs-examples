using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Search;
using Azure.ResourceManager.Search.Models;

// Generated from example definition: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/GetSharedPrivateLinkResource.json
// this example is just showing the usage of "SharedPrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
string sharedPrivateLinkResourceName = "testResource";
bool result = await collection.ExistsAsync(sharedPrivateLinkResourceName);

Console.WriteLine($"Succeeded: {result}");
