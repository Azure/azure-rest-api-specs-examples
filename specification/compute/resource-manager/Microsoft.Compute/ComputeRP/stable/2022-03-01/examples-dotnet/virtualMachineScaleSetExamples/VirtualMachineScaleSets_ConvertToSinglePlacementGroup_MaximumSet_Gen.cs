using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ConvertToSinglePlacementGroup_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineScaleSets_ConvertToSinglePlacementGroup" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this VirtualMachineScaleSetResource created on azure
// for more information of creating VirtualMachineScaleSetResource, please refer to the document of VirtualMachineScaleSetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string virtualMachineScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier virtualMachineScaleSetResourceId = VirtualMachineScaleSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName);
VirtualMachineScaleSetResource virtualMachineScaleSet = client.GetVirtualMachineScaleSetResource(virtualMachineScaleSetResourceId);

// invoke the operation
VirtualMachineScaleSetConvertToSinglePlacementGroupContent content = new VirtualMachineScaleSetConvertToSinglePlacementGroupContent()
{
    ActivePlacementGroupId = "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
};
await virtualMachineScaleSet.ConvertToSinglePlacementGroupAsync(content);

Console.WriteLine($"Succeeded");
