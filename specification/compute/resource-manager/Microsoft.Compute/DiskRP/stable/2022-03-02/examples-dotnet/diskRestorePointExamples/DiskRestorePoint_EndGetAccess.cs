using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskRestorePointExamples/DiskRestorePoint_EndGetAccess.json
// this example is just showing the usage of "DiskRestorePoint_RevokeAccess" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this DiskRestorePointResource created on azure
// for more information of creating DiskRestorePointResource, please refer to the document of DiskRestorePointResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string restorePointGroupName = "rpc";
string vmRestorePointName = "vmrp";
string diskRestorePointName = "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745";
ResourceIdentifier diskRestorePointResourceId = DiskRestorePointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, restorePointGroupName, vmRestorePointName, diskRestorePointName);
DiskRestorePointResource diskRestorePoint = client.GetDiskRestorePointResource(diskRestorePointResourceId);

// invoke the operation
await diskRestorePoint.RevokeAccessAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
