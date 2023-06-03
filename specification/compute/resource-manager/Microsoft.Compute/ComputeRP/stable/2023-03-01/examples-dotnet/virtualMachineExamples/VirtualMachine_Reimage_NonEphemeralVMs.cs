using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExamples/VirtualMachine_Reimage_NonEphemeralVMs.json
// this example is just showing the usage of "VirtualMachines_Reimage" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineResource created on azure
// for more information of creating VirtualMachineResource, please refer to the document of VirtualMachineResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string vmName = "myVMName";
ResourceIdentifier virtualMachineResourceId = VirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName);
VirtualMachineResource virtualMachine = client.GetVirtualMachineResource(virtualMachineResourceId);

// invoke the operation
VirtualMachineReimageContent content = new VirtualMachineReimageContent()
{
    TempDisk = true,
    ExactVersion = "aaaaaa",
    OSProfile = new OSProfileProvisioningData()
    {
        AdminPassword = "{your-password}",
        CustomData = "{your-custom-data}",
    },
};
await virtualMachine.ReimageAsync(WaitUntil.Completed, content: content);

Console.WriteLine($"Succeeded");
