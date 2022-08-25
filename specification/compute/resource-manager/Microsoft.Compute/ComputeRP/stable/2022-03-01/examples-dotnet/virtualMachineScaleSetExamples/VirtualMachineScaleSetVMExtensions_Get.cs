using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtensions_Get.json
// this example is just showing the usage of "VirtualMachineScaleSetVMExtensions_Get" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this VirtualMachineScaleSetVmResource created on azure
// for more information of creating VirtualMachineScaleSetVmResource, please refer to the document of VirtualMachineScaleSetVmResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string virtualMachineScaleSetName = "myvmScaleSet";
string instanceId = "0";
ResourceIdentifier virtualMachineScaleSetVmResourceId = VirtualMachineScaleSetVmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName, instanceId);
VirtualMachineScaleSetVmResource virtualMachineScaleSetVm = client.GetVirtualMachineScaleSetVmResource(virtualMachineScaleSetVmResourceId);
            
// get the collection of this VirtualMachineScaleSetVmExtensionResource
VirtualMachineScaleSetVmExtensionCollection collection = virtualMachineScaleSetVm.GetVirtualMachineScaleSetVmExtensions();
            
// invoke the operation
string vmExtensionName = "myVMExtension";
bool result = await collection.ExistsAsync(vmExtensionName);
            
Console.WriteLine($"Succeeded: {result}");
