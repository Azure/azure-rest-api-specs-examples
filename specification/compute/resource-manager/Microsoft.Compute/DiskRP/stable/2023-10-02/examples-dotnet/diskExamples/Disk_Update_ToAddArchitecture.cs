using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskExamples/Disk_Update_ToAddArchitecture.json
// this example is just showing the usage of "Disks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDiskResource created on azure
// for more information of creating ManagedDiskResource, please refer to the document of ManagedDiskResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string diskName = "myDisk";
ResourceIdentifier managedDiskResourceId = ManagedDiskResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskName);
ManagedDiskResource managedDisk = client.GetManagedDiskResource(managedDiskResourceId);

// invoke the operation
ManagedDiskPatch patch = new ManagedDiskPatch()
{
    SupportedCapabilities = new SupportedCapabilities()
    {
        Architecture = ArchitectureType.Arm64,
    },
};
ArmOperation<ManagedDiskResource> lro = await managedDisk.UpdateAsync(WaitUntil.Completed, patch);
ManagedDiskResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedDiskData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
