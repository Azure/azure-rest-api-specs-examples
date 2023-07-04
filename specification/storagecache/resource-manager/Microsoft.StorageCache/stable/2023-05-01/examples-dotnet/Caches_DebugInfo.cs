using System;
using System.Collections.Generic;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StorageCache;
using Azure.ResourceManager.StorageCache.Models;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/examples/Caches_DebugInfo.json
// this example is just showing the usage of "Caches_DebugInfo" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageCacheResource created on azure
// for more information of creating StorageCacheResource, please refer to the document of StorageCacheResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string cacheName = "sc";
ResourceIdentifier storageCacheResourceId = StorageCacheResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName);
StorageCacheResource storageCache = client.GetStorageCacheResource(storageCacheResourceId);

// invoke the operation
await storageCache.EnableDebugInfoAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
