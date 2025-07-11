using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PureStorageBlock.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PureStorageBlock;

// Generated from example definition: 2024-11-01/StoragePools_Create_MaximumSet_Gen.json
// this example is just showing the usage of "StoragePool_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
string resourceGroupName = "rgpurestorage";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PureStoragePoolResource
PureStoragePoolCollection collection = resourceGroupResource.GetPureStoragePools();

// invoke the operation
string storagePoolName = "storagePoolname";
PureStoragePoolData data = new PureStoragePoolData(new AzureLocation("lonlc"))
{
    Properties = new PureStoragePoolProperties("vknyl", new PureStoragePoolVnetInjection(new ResourceIdentifier("tnlctolrxdvnkjiphlrdxq"), new ResourceIdentifier("zbumtytyqwewjcyckwqchiypshv")), 17L, new ResourceIdentifier("xiowoxnbtcotutcmmrofvgdi")),
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key4211")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["key7593"] = "vsyiygyurvwlfaezpuqu"
    },
};
ArmOperation<PureStoragePoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, storagePoolName, data);
PureStoragePoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PureStoragePoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
