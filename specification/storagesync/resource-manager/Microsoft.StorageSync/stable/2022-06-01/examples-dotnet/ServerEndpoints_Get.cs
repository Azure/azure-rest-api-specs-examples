using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageSync.Models;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/ServerEndpoints_Get.json
// this example is just showing the usage of "ServerEndpoints_Get" operation, for the dependent resources, they will have to be created separately.

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
NullableResponse<StorageSyncServerEndpointResource> response = await collection.GetIfExistsAsync(serverEndpointName);
StorageSyncServerEndpointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    StorageSyncServerEndpointData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
