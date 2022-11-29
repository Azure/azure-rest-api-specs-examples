using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskExamples/Disk_BeginGetAccess_WithVMGuestState.json
// this example is just showing the usage of "Disks_GrantAccess" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ManagedDiskResource created on azure
// for more information of creating ManagedDiskResource, please refer to the document of ManagedDiskResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string diskName = "myDisk";
ResourceIdentifier managedDiskResourceId = ManagedDiskResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskName);
ManagedDiskResource managedDisk = client.GetManagedDiskResource(managedDiskResourceId);

// invoke the operation
GrantAccessData data = new GrantAccessData(AccessLevel.Read, 300)
{
    GetSecureVmGuestStateSas = true,
};
ArmOperation<AccessUri> lro = await managedDisk.GrantAccessAsync(WaitUntil.Completed, data);
AccessUri result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
