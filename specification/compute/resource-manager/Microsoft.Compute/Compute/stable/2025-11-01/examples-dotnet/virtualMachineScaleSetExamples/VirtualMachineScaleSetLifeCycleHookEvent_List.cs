using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/Compute/stable/2025-11-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetLifeCycleHookEvent_List.json
// this example is just showing the usage of "VirtualMachineScaleSetLifeCycleHookEvents_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineScaleSetResource created on azure
// for more information of creating VirtualMachineScaleSetResource, please refer to the document of VirtualMachineScaleSetResource
string subscriptionId = "2e2e3046-f85f-4966-8fd2-5fd7bf6ea717";
string resourceGroupName = "RG01";
string virtualMachineScaleSetName = "VMSS01";
ResourceIdentifier virtualMachineScaleSetResourceId = VirtualMachineScaleSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName);
VirtualMachineScaleSetResource virtualMachineScaleSet = client.GetVirtualMachineScaleSetResource(virtualMachineScaleSetResourceId);

// get the collection of this VmScaleSetLifecycleHookEventResource
VmScaleSetLifecycleHookEventCollection collection = virtualMachineScaleSet.GetVmScaleSetLifecycleHookEvents();

// invoke the operation and iterate over the result
await foreach (VmScaleSetLifecycleHookEventResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    VmScaleSetLifecycleHookEventData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
