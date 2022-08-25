using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExamples/VirtualMachineExtensions_Update.json
// this example is just showing the usage of "VirtualMachineExtensions_Update" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this VirtualMachineExtensionResource created on azure
// for more information of creating VirtualMachineExtensionResource, please refer to the document of VirtualMachineExtensionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string vmName = "myVM";
string vmExtensionName = "myVMExtension";
ResourceIdentifier virtualMachineExtensionResourceId = VirtualMachineExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, vmExtensionName);
VirtualMachineExtensionResource virtualMachineExtension = client.GetVirtualMachineExtensionResource(virtualMachineExtensionResourceId);
            
// invoke the operation
VirtualMachineExtensionPatch patch = new VirtualMachineExtensionPatch()
{
    Publisher = "extPublisher",
    ExtensionType = "extType",
    TypeHandlerVersion = "1.2",
    AutoUpgradeMinorVersion = true,
    Settings = BinaryData.FromObjectAsJson(new
    {
        UserName = "xyz@microsoft.com"
    }),
    SuppressFailures = true,
    ProtectedSettingsFromKeyVault = BinaryData.FromObjectAsJson(new
    {
        secretUrl = "https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e",
        sourceVault = new
        {
            id = "/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"
        }
    }),
};
ArmOperation<VirtualMachineExtensionResource> lro = await virtualMachineExtension.UpdateAsync(WaitUntil.Completed, patch);
VirtualMachineExtensionResource result = lro.Value;
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineExtensionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
