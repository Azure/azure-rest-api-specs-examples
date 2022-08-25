using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_Update_MinimumSet_Gen.json
// this example is just showing the usage of "VirtualMachineScaleSetVMs_Update" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this VirtualMachineScaleSetVmResource created on azure
// for more information of creating VirtualMachineScaleSetVmResource, please refer to the document of VirtualMachineScaleSetVmResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string virtualMachineScaleSetName = "aaaaaaaaaaaaaaaaaa";
string instanceId = "aaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier virtualMachineScaleSetVmResourceId = VirtualMachineScaleSetVmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName, instanceId);
VirtualMachineScaleSetVmResource virtualMachineScaleSetVm = client.GetVirtualMachineScaleSetVmResource(virtualMachineScaleSetVmResourceId);

// invoke the operation
VirtualMachineScaleSetVmData data = new VirtualMachineScaleSetVmData(new AzureLocation("westus"));
ArmOperation<VirtualMachineScaleSetVmResource> lro = await virtualMachineScaleSetVm.UpdateAsync(WaitUntil.Completed, data);
VirtualMachineScaleSetVmResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineScaleSetVmData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
