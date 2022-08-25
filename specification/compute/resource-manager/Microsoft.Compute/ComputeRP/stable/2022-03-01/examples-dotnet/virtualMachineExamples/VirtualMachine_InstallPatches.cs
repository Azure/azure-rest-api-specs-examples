using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExamples/VirtualMachine_InstallPatches.json
// this example is just showing the usage of "VirtualMachines_InstallPatches" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this VirtualMachineResource created on azure
// for more information of creating VirtualMachineResource, please refer to the document of VirtualMachineResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroupName";
string vmName = "myVMName";
ResourceIdentifier virtualMachineResourceId = VirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName);
VirtualMachineResource virtualMachine = client.GetVirtualMachineResource(virtualMachineResourceId);
            
// invoke the operation
VirtualMachineInstallPatchesContent content = new VirtualMachineInstallPatchesContent(VmGuestPatchRebootSetting.IfRequired)
{
    MaximumDuration = XmlConvert.ToTimeSpan("PT4H"),
    WindowsParameters = new WindowsParameters()
    {
        ClassificationsToInclude =
                    {
                    VmGuestPatchClassificationForWindows.Critical,VmGuestPatchClassificationForWindows.Security
                    },
        MaxPatchPublishOn = DateTimeOffset.Parse("2020-11-19T02:36:43.0539904+00:00"),
    },
};
ArmOperation<VirtualMachineInstallPatchesResult> lro = await virtualMachine.InstallPatchesAsync(WaitUntil.Completed, content);
VirtualMachineInstallPatchesResult result = lro.Value;
            
Console.WriteLine($"Succeeded: {result}");
