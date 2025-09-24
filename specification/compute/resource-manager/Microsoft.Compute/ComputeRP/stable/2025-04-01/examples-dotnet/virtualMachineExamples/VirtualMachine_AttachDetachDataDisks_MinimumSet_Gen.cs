using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachine_AttachDetachDataDisks_MinimumSet_Gen.json
// this example is just showing the usage of "VirtualMachines_AttachDetachDataDisks" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineResource created on azure
// for more information of creating VirtualMachineResource, please refer to the document of VirtualMachineResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string vmName = "azure-vm";
ResourceIdentifier virtualMachineResourceId = VirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName);
VirtualMachineResource virtualMachine = client.GetVirtualMachineResource(virtualMachineResourceId);

// invoke the operation
AttachDetachDataDisksRequest attachDetachDataDisksRequest = new AttachDetachDataDisksRequest
{
    DataDisksToAttach = { new DataDisksToAttach("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d") },
    DataDisksToDetach = { new DataDisksToDetach("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_1_disk1_1a4e784bdafa49baa780eb2d128ff65x") },
};
ArmOperation<VirtualMachineStorageProfile> lro = await virtualMachine.AttachDetachDataDisksAsync(WaitUntil.Completed, attachDetachDataDisksRequest);
VirtualMachineStorageProfile result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
