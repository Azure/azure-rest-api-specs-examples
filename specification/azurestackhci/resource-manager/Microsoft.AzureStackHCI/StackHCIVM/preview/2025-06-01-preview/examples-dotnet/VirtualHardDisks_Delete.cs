using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Vm.Models;
using Azure.ResourceManager.Hci.Vm;

// Generated from example definition: 2025-06-01-preview/VirtualHardDisks_Delete.json
// this example is just showing the usage of "VirtualHardDisk_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciVmVirtualHardDiskResource created on azure
// for more information of creating HciVmVirtualHardDiskResource, please refer to the document of HciVmVirtualHardDiskResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string virtualHardDiskName = "test-vhd";
ResourceIdentifier hciVmVirtualHardDiskResourceId = HciVmVirtualHardDiskResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHardDiskName);
HciVmVirtualHardDiskResource hciVmVirtualHardDisk = client.GetHciVmVirtualHardDiskResource(hciVmVirtualHardDiskResourceId);

// invoke the operation
await hciVmVirtualHardDisk.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
