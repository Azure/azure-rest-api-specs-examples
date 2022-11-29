using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtensions_Get.json
// this example is just showing the usage of "VirtualMachineScaleSetVMExtensions_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this VirtualMachineScaleSetVmExtensionResource created on azure
// for more information of creating VirtualMachineScaleSetVmExtensionResource, please refer to the document of VirtualMachineScaleSetVmExtensionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string virtualMachineScaleSetName = "myvmScaleSet";
string instanceId = "0";
string vmExtensionName = "myVMExtension";
ResourceIdentifier virtualMachineScaleSetVmExtensionResourceId = VirtualMachineScaleSetVmExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName, instanceId, vmExtensionName);
VirtualMachineScaleSetVmExtensionResource virtualMachineScaleSetVmExtension = client.GetVirtualMachineScaleSetVmExtensionResource(virtualMachineScaleSetVmExtensionResourceId);

// invoke the operation
VirtualMachineScaleSetVmExtensionResource result = await virtualMachineScaleSetVmExtension.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineScaleSetVmExtensionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
