using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageSync.Models;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/ServerEndpoints_Create.json
// this example is just showing the usage of "ServerEndpoints_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncGroupResource created on azure
// for more information of creating StorageSyncGroupResource, please refer to the document of StorageSyncGroupResource
string subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
string resourceGroupName = "SampleResourceGroup_1";
string storageSyncServiceName = "SampleStorageSyncService_1";
string syncGroupName = "SampleSyncGroup_1";
ResourceIdentifier storageSyncGroupResourceId = StorageSyncGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName, syncGroupName);
StorageSyncGroupResource storageSyncGroup = client.GetStorageSyncGroupResource(storageSyncGroupResourceId);

// get the collection of this StorageSyncServerEndpointResource
StorageSyncServerEndpointCollection collection = storageSyncGroup.GetStorageSyncServerEndpoints();

// invoke the operation
string serverEndpointName = "SampleServerEndpoint_1";
StorageSyncServerEndpointCreateOrUpdateContent content = new StorageSyncServerEndpointCreateOrUpdateContent
{
    ServerLocalPath = "D:\\SampleServerEndpoint_1",
    CloudTiering = StorageSyncFeatureStatus.Off,
    VolumeFreeSpacePercent = 100,
    TierFilesOlderThanDays = 0,
    ServerResourceId = new ResourceIdentifier("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
    OfflineDataTransfer = StorageSyncFeatureStatus.On,
    OfflineDataTransferShareName = "myfileshare",
    InitialDownloadPolicy = InitialDownloadPolicy.NamespaceThenModifiedFiles,
    LocalCacheMode = LocalCacheMode.UpdateLocallyCachedFiles,
    InitialUploadPolicy = InitialUploadPolicy.ServerAuthoritative,
};
ArmOperation<StorageSyncServerEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverEndpointName, content);
StorageSyncServerEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageSyncServerEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
