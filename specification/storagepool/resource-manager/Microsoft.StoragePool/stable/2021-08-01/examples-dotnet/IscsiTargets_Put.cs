using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StoragePool.Models;
using Azure.ResourceManager.StoragePool;

// Generated from example definition: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Put.json
// this example is just showing the usage of "IscsiTargets_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiskPoolResource created on azure
// for more information of creating DiskPoolResource, please refer to the document of DiskPoolResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string diskPoolName = "myDiskPool";
ResourceIdentifier diskPoolResourceId = DiskPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskPoolName);
DiskPoolResource diskPool = client.GetDiskPoolResource(diskPoolResourceId);

// get the collection of this DiskPoolIscsiTargetResource
DiskPoolIscsiTargetCollection collection = diskPool.GetDiskPoolIscsiTargets();

// invoke the operation
string iscsiTargetName = "myIscsiTarget";
DiskPoolIscsiTargetCreateOrUpdateContent content = new DiskPoolIscsiTargetCreateOrUpdateContent(DiskPoolIscsiTargetAclMode.Dynamic)
{
    TargetIqn = "iqn.2005-03.org.iscsi:server1",
    Luns = { new ManagedDiskIscsiLun("lun0", new ResourceIdentifier("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1")) },
};
ArmOperation<DiskPoolIscsiTargetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, iscsiTargetName, content);
DiskPoolIscsiTargetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiskPoolIscsiTargetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
