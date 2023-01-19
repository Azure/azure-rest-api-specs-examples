using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StorageSync;
using Azure.ResourceManager.StorageSync.Models;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/PrivateEndpointConnections_ListByStorageSyncService.json
// this example is just showing the usage of "PrivateEndpointConnections_ListByStorageSyncService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncServiceResource created on azure
// for more information of creating StorageSyncServiceResource, please refer to the document of StorageSyncServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string storageSyncServiceName = "sss2527";
ResourceIdentifier storageSyncServiceResourceId = StorageSyncServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName);
StorageSyncServiceResource storageSyncService = client.GetStorageSyncServiceResource(storageSyncServiceResourceId);

// get the collection of this StorageSyncPrivateEndpointConnectionResource
StorageSyncPrivateEndpointConnectionCollection collection = storageSyncService.GetStorageSyncPrivateEndpointConnections();

// invoke the operation and iterate over the result
await foreach (StorageSyncPrivateEndpointConnectionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    StorageSyncPrivateEndpointConnectionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
