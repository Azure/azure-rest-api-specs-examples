using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PureStorageBlock;

// Generated from example definition: 2024-11-01-preview/AvsVmVolumes_ListByAvsVm_MaximumSet_Gen.json
// this example is just showing the usage of "AvsVmVolume_ListByAvsVm" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PureStorageAvsVmResource created on azure
// for more information of creating PureStorageAvsVmResource, please refer to the document of PureStorageAvsVmResource
string subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
string resourceGroupName = "rgpurestorage";
string storagePoolName = "storagePoolname";
string avsVmId = "cbdec-ddbb";
ResourceIdentifier pureStorageAvsVmResourceId = PureStorageAvsVmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storagePoolName, avsVmId);
PureStorageAvsVmResource pureStorageAvsVm = client.GetPureStorageAvsVmResource(pureStorageAvsVmResourceId);

// get the collection of this PureStorageAvsVmVolumeResource
PureStorageAvsVmVolumeCollection collection = pureStorageAvsVm.GetPureStorageAvsVmVolumes();

// invoke the operation and iterate over the result
await foreach (PureStorageAvsVmVolumeResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PureStorageAvsVmVolumeData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
