using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StorageSync;
using Azure.ResourceManager.StorageSync.Models;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/CloudEndpoints_PreRestore.json
// this example is just showing the usage of "CloudEndpoints_PreRestore" operation, for the dependent resources, they will have to be created separately.

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
PreRestoreContent content = new PreRestoreContent()
{
    AzureFileShareUri = new Uri("https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare"),
    RestoreFileSpec =
    {
    new RestoreFileSpec()
    {
    Path = "text1.txt",
    IsDirectory = false,
    },new RestoreFileSpec()
    {
    Path = "MyDir",
    IsDirectory = true,
    },new RestoreFileSpec()
    {
    Path = "MyDir/SubDir",
    IsDirectory = false,
    },new RestoreFileSpec()
    {
    Path = "MyDir/SubDir/File1.pdf",
    IsDirectory = false,
    }
    },
};
await cloudEndpoint.PreRestoreAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
