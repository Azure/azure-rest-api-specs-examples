using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/extension/Extensions_Upgrade.json
// this example is just showing the usage of "UpgradeExtensions" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeMachineResource created on azure
// for more information of creating HybridComputeMachineResource, please refer to the document of HybridComputeMachineResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string machineName = "myMachine";
ResourceIdentifier hybridComputeMachineResourceId = HybridComputeMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName);
HybridComputeMachineResource hybridComputeMachine = client.GetHybridComputeMachineResource(hybridComputeMachineResourceId);

// invoke the operation
MachineExtensionUpgrade extensionUpgradeParameters = new MachineExtensionUpgrade()
{
    ExtensionTargets =
    {
    ["Microsoft.Azure.Monitoring"] = new ExtensionTargetProperties()
    {
    TargetVersion = "2.0",
    },
    ["Microsoft.Compute.CustomScriptExtension"] = new ExtensionTargetProperties()
    {
    TargetVersion = "1.10",
    },
    },
};
await hybridComputeMachine.UpgradeExtensionsAsync(WaitUntil.Completed, extensionUpgradeParameters);

Console.WriteLine($"Succeeded");
