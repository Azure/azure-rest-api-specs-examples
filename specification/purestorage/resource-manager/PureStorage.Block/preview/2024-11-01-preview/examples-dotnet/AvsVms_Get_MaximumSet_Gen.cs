using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PureStorageBlock;

// Generated from example definition: 2024-11-01-preview/AvsVms_Get_MaximumSet_Gen.json
// this example is just showing the usage of "AvsVm_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PureStoragePoolResource created on azure
// for more information of creating PureStoragePoolResource, please refer to the document of PureStoragePoolResource
string subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
string resourceGroupName = "rgpurestorage";
string storagePoolName = "storagePoolname";
ResourceIdentifier pureStoragePoolResourceId = PureStoragePoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storagePoolName);
PureStoragePoolResource pureStoragePool = client.GetPureStoragePoolResource(pureStoragePoolResourceId);

// get the collection of this PureStorageAvsVmResource
PureStorageAvsVmCollection collection = pureStoragePool.GetPureStorageAvsVms();

// invoke the operation
string avsVmId = "cbdec-ddbb";
NullableResponse<PureStorageAvsVmResource> response = await collection.GetIfExistsAsync(avsVmId);
PureStorageAvsVmResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PureStorageAvsVmData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
