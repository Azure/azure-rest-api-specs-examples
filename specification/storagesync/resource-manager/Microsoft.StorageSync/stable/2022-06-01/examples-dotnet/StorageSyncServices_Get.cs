using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StorageSync;
using Azure.ResourceManager.StorageSync.Models;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/StorageSyncServices_Get.json
// this example is just showing the usage of "StorageSyncServices_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncServiceResource created on azure
// for more information of creating StorageSyncServiceResource, please refer to the document of StorageSyncServiceResource
string subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
string resourceGroupName = "SampleResourceGroup_1";
string storageSyncServiceName = "SampleStorageSyncService_1";
ResourceIdentifier storageSyncServiceResourceId = StorageSyncServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName);
StorageSyncServiceResource storageSyncService = client.GetStorageSyncServiceResource(storageSyncServiceResourceId);

// invoke the operation
StorageSyncServiceResource result = await storageSyncService.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageSyncServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
