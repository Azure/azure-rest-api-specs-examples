using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_CreateOrUpdate_MinimumSet_Gen.json
// this example is just showing the usage of "VirtualMachineScaleSetExtensions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this VirtualMachineScaleSetResource created on azure
// for more information of creating VirtualMachineScaleSetResource, please refer to the document of VirtualMachineScaleSetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string virtualMachineScaleSetName = "aaaaaaaaaaa";
ResourceIdentifier virtualMachineScaleSetResourceId = VirtualMachineScaleSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName);
VirtualMachineScaleSetResource virtualMachineScaleSet = client.GetVirtualMachineScaleSetResource(virtualMachineScaleSetResourceId);
            
// get the collection of this VirtualMachineScaleSetExtensionResource
VirtualMachineScaleSetExtensionCollection collection = virtualMachineScaleSet.GetVirtualMachineScaleSetExtensions();
            
// invoke the operation
string vmssExtensionName = "aaaaaaaaaaa";
VirtualMachineScaleSetExtensionData data = new VirtualMachineScaleSetExtensionData();
ArmOperation<VirtualMachineScaleSetExtensionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vmssExtensionName, data);
VirtualMachineScaleSetExtensionResource result = lro.Value;
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineScaleSetExtensionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
