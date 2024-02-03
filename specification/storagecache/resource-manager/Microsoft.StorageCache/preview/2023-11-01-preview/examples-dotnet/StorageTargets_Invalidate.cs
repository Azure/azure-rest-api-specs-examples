using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StorageCache;
using Azure.ResourceManager.StorageCache.Models;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_Invalidate.json
// this example is just showing the usage of "StorageTargets_Invalidate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageTargetResource created on azure
// for more information of creating StorageTargetResource, please refer to the document of StorageTargetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string cacheName = "sc";
string storageTargetName = "st1";
ResourceIdentifier storageTargetResourceId = StorageTargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName, storageTargetName);
StorageTargetResource storageTarget = client.GetStorageTargetResource(storageTargetResourceId);

// invoke the operation
await storageTarget.InvalidateAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
