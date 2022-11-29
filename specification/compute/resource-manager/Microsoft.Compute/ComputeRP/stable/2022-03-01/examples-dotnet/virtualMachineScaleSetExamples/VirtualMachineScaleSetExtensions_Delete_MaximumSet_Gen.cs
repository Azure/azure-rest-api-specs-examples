using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineScaleSetExtensions_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this VirtualMachineScaleSetExtensionResource created on azure
// for more information of creating VirtualMachineScaleSetExtensionResource, please refer to the document of VirtualMachineScaleSetExtensionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string virtualMachineScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
string vmssExtensionName = "aaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier virtualMachineScaleSetExtensionResourceId = VirtualMachineScaleSetExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName, vmssExtensionName);
VirtualMachineScaleSetExtensionResource virtualMachineScaleSetExtension = client.GetVirtualMachineScaleSetExtensionResource(virtualMachineScaleSetExtensionResourceId);

// invoke the operation
await virtualMachineScaleSetExtension.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
