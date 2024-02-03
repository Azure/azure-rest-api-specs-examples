using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StorageCache;
using Azure.ResourceManager.StorageCache.Models;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_CreateOrUpdate_BlobNfs.json
// this example is just showing the usage of "StorageTargets_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageTargetResource created on azure
// for more information of creating StorageTargetResource, please refer to the document of StorageTargetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string cacheName = "sc1";
string storageTargetName = "st1";
ResourceIdentifier storageTargetResourceId = StorageTargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName, storageTargetName);
StorageTargetResource storageTarget = client.GetStorageTargetResource(storageTargetResourceId);

// invoke the operation
StorageTargetData data = new StorageTargetData()
{
    Junctions =
    {
    new NamespaceJunction()
    {
    NamespacePath = "/blobnfs",
    }
    },
    TargetType = StorageTargetType.BlobNfs,
    BlobNfs = new BlobNfsTarget()
    {
        Target = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/blofnfs/blobServices/default/containers/blobnfs"),
        UsageModel = "READ_WRITE",
        VerificationDelayInSeconds = 28800,
        WriteBackDelayInSeconds = 3600,
    },
};
ArmOperation<StorageTargetResource> lro = await storageTarget.UpdateAsync(WaitUntil.Completed, data);
StorageTargetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageTargetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
