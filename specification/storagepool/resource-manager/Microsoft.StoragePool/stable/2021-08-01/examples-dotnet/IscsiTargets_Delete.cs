using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StoragePool.Models;
using Azure.ResourceManager.StoragePool;

// Generated from example definition: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Delete.json
// this example is just showing the usage of "IscsiTargets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiskPoolIscsiTargetResource created on azure
// for more information of creating DiskPoolIscsiTargetResource, please refer to the document of DiskPoolIscsiTargetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string diskPoolName = "myDiskPool";
string iscsiTargetName = "myIscsiTarget";
ResourceIdentifier diskPoolIscsiTargetResourceId = DiskPoolIscsiTargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskPoolName, iscsiTargetName);
DiskPoolIscsiTargetResource diskPoolIscsiTarget = client.GetDiskPoolIscsiTargetResource(diskPoolIscsiTargetResourceId);

// invoke the operation
await diskPoolIscsiTarget.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
