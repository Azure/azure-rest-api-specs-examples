using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateVirtualHardDisk.json
// this example is just showing the usage of "VirtualHardDisks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualHardDiskResource created on azure
// for more information of creating VirtualHardDiskResource, please refer to the document of VirtualHardDiskResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string virtualHardDiskName = "test-vhd";
ResourceIdentifier virtualHardDiskResourceId = VirtualHardDiskResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHardDiskName);
VirtualHardDiskResource virtualHardDisk = client.GetVirtualHardDiskResource(virtualHardDiskResourceId);

// invoke the operation
VirtualHardDiskPatch patch = new VirtualHardDiskPatch()
{
    Tags =
    {
    ["additionalProperties"] = "sample",
    },
};
ArmOperation<VirtualHardDiskResource> lro = await virtualHardDisk.UpdateAsync(WaitUntil.Completed, patch);
VirtualHardDiskResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualHardDiskData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
