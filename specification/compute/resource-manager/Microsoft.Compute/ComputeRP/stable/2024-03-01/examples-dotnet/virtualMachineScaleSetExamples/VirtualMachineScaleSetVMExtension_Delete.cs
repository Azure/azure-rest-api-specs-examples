using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Delete.json
// this example is just showing the usage of "VirtualMachineScaleSetVMExtensions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

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
await virtualMachineScaleSetVmExtension.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
