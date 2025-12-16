using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageSync.Models;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-09-01/examples/RegisteredServers_Get.json
// this example is just showing the usage of "RegisteredServers_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncRegisteredServerResource created on azure
// for more information of creating StorageSyncRegisteredServerResource, please refer to the document of StorageSyncRegisteredServerResource
string subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
string resourceGroupName = "SampleResourceGroup_1";
string storageSyncServiceName = "SampleStorageSyncService_1";
Guid serverId = Guid.Parse("080d4133-bdb5-40a0-96a0-71a6057bfe9a");
ResourceIdentifier storageSyncRegisteredServerResourceId = StorageSyncRegisteredServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName, serverId);
StorageSyncRegisteredServerResource storageSyncRegisteredServer = client.GetStorageSyncRegisteredServerResource(storageSyncRegisteredServerResourceId);

// invoke the operation
StorageSyncRegisteredServerResource result = await storageSyncRegisteredServer.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageSyncRegisteredServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
