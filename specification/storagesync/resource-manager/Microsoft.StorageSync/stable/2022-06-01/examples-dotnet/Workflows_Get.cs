using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/Workflows_Get.json
// this example is just showing the usage of "Workflows_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this StorageSyncWorkflowResource
StorageSyncWorkflowCollection collection = storageSyncService.GetStorageSyncWorkflows();

// invoke the operation
string workflowId = "828219ea-083e-48b5-89ea-8fd9991b2e75";
bool result = await collection.ExistsAsync(workflowId);

Console.WriteLine($"Succeeded: {result}");
