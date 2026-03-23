using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-12-15-preview/examples/Caches_Update.json
// this example is just showing the usage of "Caches_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CacheResource created on azure
// for more information of creating CacheResource, please refer to the document of CacheResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string cacheName = "cache1";
ResourceIdentifier cacheResourceId = CacheResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, cacheName);
CacheResource cache = client.GetCacheResource(cacheResourceId);

// invoke the operation
CachePatch patch = new CachePatch
{
    Properties = new CacheUpdateProperties
    {
        Size = 214748364800L,
    },
};
ArmOperation<CacheResource> lro = await cache.UpdateAsync(WaitUntil.Completed, patch);
CacheResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CacheData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
