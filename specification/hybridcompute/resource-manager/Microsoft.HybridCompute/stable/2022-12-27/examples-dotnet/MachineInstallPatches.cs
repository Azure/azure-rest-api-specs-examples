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

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-12-27/examples/MachineInstallPatches.json
// this example is just showing the usage of "Machines_InstallPatches" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeMachineResource created on azure
// for more information of creating HybridComputeMachineResource, please refer to the document of HybridComputeMachineResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroupName";
string name = "myMachineName";
ResourceIdentifier hybridComputeMachineResourceId = HybridComputeMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
HybridComputeMachineResource hybridComputeMachine = client.GetHybridComputeMachineResource(hybridComputeMachineResourceId);

// invoke the operation
MachineInstallPatchesContent content = new MachineInstallPatchesContent(XmlConvert.ToTimeSpan("PT4H"), VmGuestPatchRebootSetting.IfRequired)
{
    WindowsParameters = new WindowsParameters()
    {
        ClassificationsToInclude =
        {
        VmGuestPatchClassificationWindow.Critical,VmGuestPatchClassificationWindow.Security
        },
        MaxPatchPublishOn = DateTimeOffset.Parse("2021-08-19T02:36:43.0539904+00:00"),
    },
};
ArmOperation<MachineInstallPatchesResult> lro = await hybridComputeMachine.InstallPatchesAsync(WaitUntil.Completed, content);
MachineInstallPatchesResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
