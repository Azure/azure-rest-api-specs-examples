using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Search.Models;
using Azure.ResourceManager.Search;

// Generated from example definition: specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/DeleteSharedPrivateLinkResource.json
// this example is just showing the usage of "SharedPrivateLinkResources_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SharedSearchServicePrivateLinkResource created on azure
// for more information of creating SharedSearchServicePrivateLinkResource, please refer to the document of SharedSearchServicePrivateLinkResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string searchServiceName = "mysearchservice";
string sharedPrivateLinkResourceName = "testResource";
ResourceIdentifier sharedSearchServicePrivateLinkResourceId = SharedSearchServicePrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName);
SharedSearchServicePrivateLinkResource sharedSearchServicePrivateLinkResource = client.GetSharedSearchServicePrivateLinkResource(sharedSearchServicePrivateLinkResourceId);

// invoke the operation
await sharedSearchServicePrivateLinkResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
