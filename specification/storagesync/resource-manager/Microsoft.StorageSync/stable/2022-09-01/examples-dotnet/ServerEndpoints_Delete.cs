using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageSync.Models;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-09-01/examples/ServerEndpoints_Delete.json
// this example is just showing the usage of "ServerEndpoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncServerEndpointResource created on azure
// for more information of creating StorageSyncServerEndpointResource, please refer to the document of StorageSyncServerEndpointResource
string subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
string resourceGroupName = "SampleResourceGroup_1";
string storageSyncServiceName = "SampleStorageSyncService_1";
string syncGroupName = "SampleSyncGroup_1";
string serverEndpointName = "SampleServerEndpoint_1";
ResourceIdentifier storageSyncServerEndpointResourceId = StorageSyncServerEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName);
StorageSyncServerEndpointResource storageSyncServerEndpoint = client.GetStorageSyncServerEndpointResource(storageSyncServerEndpointResourceId);

// invoke the operation
await storageSyncServerEndpoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
