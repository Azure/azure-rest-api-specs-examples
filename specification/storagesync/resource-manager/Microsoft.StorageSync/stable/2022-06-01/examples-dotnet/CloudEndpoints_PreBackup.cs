using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageSync.Models;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/CloudEndpoints_PreBackup.json
// this example is just showing the usage of "CloudEndpoints_PreBackup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudEndpointResource created on azure
// for more information of creating CloudEndpointResource, please refer to the document of CloudEndpointResource
string subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
string resourceGroupName = "SampleResourceGroup_1";
string storageSyncServiceName = "SampleStorageSyncService_1";
string syncGroupName = "SampleSyncGroup_1";
string cloudEndpointName = "SampleCloudEndpoint_1";
ResourceIdentifier cloudEndpointResourceId = CloudEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName);
CloudEndpointResource cloudEndpoint = client.GetCloudEndpointResource(cloudEndpointResourceId);

// invoke the operation
CloudEndpointBackupContent content = new CloudEndpointBackupContent
{
    AzureFileShare = "https://sampleserver.file.core.test-cint.azure-test.net/sampleFileShare",
};
await cloudEndpoint.PreBackupAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
