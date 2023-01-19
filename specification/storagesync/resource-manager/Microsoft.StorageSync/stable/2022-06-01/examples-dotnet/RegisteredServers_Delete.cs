using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StorageSync;
using Azure.ResourceManager.StorageSync.Models;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/RegisteredServers_Delete.json
// this example is just showing the usage of "RegisteredServers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncRegisteredServerResource created on azure
// for more information of creating StorageSyncRegisteredServerResource, please refer to the document of StorageSyncRegisteredServerResource
string subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
string resourceGroupName = "SampleResourceGroup_1";
string storageSyncServiceName = "SampleStorageSyncService_1";
Guid serverId = Guid.Parse("41166691-ab03-43e9-ab3e-0330eda162ac");
ResourceIdentifier storageSyncRegisteredServerResourceId = StorageSyncRegisteredServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName, serverId);
StorageSyncRegisteredServerResource storageSyncRegisteredServer = client.GetStorageSyncRegisteredServerResource(storageSyncRegisteredServerResourceId);

// invoke the operation
await storageSyncRegisteredServer.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
