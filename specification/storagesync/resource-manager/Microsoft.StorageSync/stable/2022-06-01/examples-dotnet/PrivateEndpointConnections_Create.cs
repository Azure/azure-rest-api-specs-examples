using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageSync.Models;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/PrivateEndpointConnections_Create.json
// this example is just showing the usage of "PrivateEndpointConnections_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncServiceResource created on azure
// for more information of creating StorageSyncServiceResource, please refer to the document of StorageSyncServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res7687";
string storageSyncServiceName = "sss2527";
ResourceIdentifier storageSyncServiceResourceId = StorageSyncServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName);
StorageSyncServiceResource storageSyncService = client.GetStorageSyncServiceResource(storageSyncServiceResourceId);

// get the collection of this StorageSyncPrivateEndpointConnectionResource
StorageSyncPrivateEndpointConnectionCollection collection = storageSyncService.GetStorageSyncPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "{privateEndpointConnectionName}";
StorageSyncPrivateEndpointConnectionData data = new StorageSyncPrivateEndpointConnectionData
{
    ConnectionState = new StorageSyncPrivateLinkServiceConnectionState
    {
        Status = StorageSyncPrivateEndpointServiceConnectionStatus.Approved,
        Description = "Auto-Approved",
    },
};
ArmOperation<StorageSyncPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
StorageSyncPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageSyncPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
