using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrades_GetLatest_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineScaleSetRollingUpgrades_GetLatest" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this VirtualMachineScaleSetRollingUpgradeResource created on azure
// for more information of creating VirtualMachineScaleSetRollingUpgradeResource, please refer to the document of VirtualMachineScaleSetRollingUpgradeResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string virtualMachineScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier virtualMachineScaleSetRollingUpgradeResourceId = VirtualMachineScaleSetRollingUpgradeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName);
VirtualMachineScaleSetRollingUpgradeResource virtualMachineScaleSetRollingUpgrade = client.GetVirtualMachineScaleSetRollingUpgradeResource(virtualMachineScaleSetRollingUpgradeResourceId);
            
// invoke the operation
VirtualMachineScaleSetRollingUpgradeResource result = await virtualMachineScaleSetRollingUpgrade.GetAsync();
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineScaleSetRollingUpgradeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
