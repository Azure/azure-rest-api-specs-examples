using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageCache.Models;
using Azure.ResourceManager.StorageCache;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/importJob_Update.json
// this example is just showing the usage of "importJobs_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageCacheImportJobResource created on azure
// for more information of creating StorageCacheImportJobResource, please refer to the document of StorageCacheImportJobResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string amlFileSystemName = "fs1";
string importJobName = "job1";
ResourceIdentifier storageCacheImportJobResourceId = StorageCacheImportJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, amlFileSystemName, importJobName);
StorageCacheImportJobResource storageCacheImportJob = client.GetStorageCacheImportJobResource(storageCacheImportJobResourceId);

// invoke the operation
StorageCacheImportJobPatch patch = new StorageCacheImportJobPatch
{
    Tags =
    {
    ["Dept"] = "ContosoAds"
    },
};
ArmOperation<StorageCacheImportJobResource> lro = await storageCacheImportJob.UpdateAsync(WaitUntil.Completed, patch);
StorageCacheImportJobResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageCacheImportJobData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
