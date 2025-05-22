using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PureStorageBlock.Models;
using Azure.ResourceManager.PureStorageBlock;

// Generated from example definition: 2024-11-01-preview/AvsVmVolumes_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "AvsVmVolume_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PureStorageAvsVmVolumeResource created on azure
// for more information of creating PureStorageAvsVmVolumeResource, please refer to the document of PureStorageAvsVmVolumeResource
string subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
string resourceGroupName = "rgpurestorage";
string storagePoolName = "storagePoolname";
string avsVmId = "cbdec-ddbb";
string volumeId = "cbdec-ddbb";
ResourceIdentifier pureStorageAvsVmVolumeResourceId = PureStorageAvsVmVolumeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storagePoolName, avsVmId, volumeId);
PureStorageAvsVmVolumeResource pureStorageAvsVmVolume = client.GetPureStorageAvsVmVolumeResource(pureStorageAvsVmVolumeResourceId);

// invoke the operation
await pureStorageAvsVmVolume.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
