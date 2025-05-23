using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/extension/Extension_Update.json
// this example is just showing the usage of "MachineExtensions_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeMachineExtensionResource created on azure
// for more information of creating HybridComputeMachineExtensionResource, please refer to the document of HybridComputeMachineExtensionResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string machineName = "myMachine";
string extensionName = "CustomScriptExtension";
ResourceIdentifier hybridComputeMachineExtensionResourceId = HybridComputeMachineExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName, extensionName);
HybridComputeMachineExtensionResource hybridComputeMachineExtension = client.GetHybridComputeMachineExtensionResource(hybridComputeMachineExtensionResourceId);

// invoke the operation
HybridComputeMachineExtensionPatch patch = new HybridComputeMachineExtensionPatch()
{
    Publisher = "Microsoft.Compute",
    MachineExtensionUpdatePropertiesType = "CustomScriptExtension",
    TypeHandlerVersion = "1.10",
    EnableAutomaticUpgrade = true,
    Settings =
    {
    ["commandToExecute"] = BinaryData.FromString("\"powershell.exe -c \"Get-Process | Where-Object { $_.CPU -lt 100 }\"\""),
    },
};
ArmOperation<HybridComputeMachineExtensionResource> lro = await hybridComputeMachineExtension.UpdateAsync(WaitUntil.Completed, patch);
HybridComputeMachineExtensionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridComputeMachineExtensionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
