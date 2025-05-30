using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/machine/Machines_Update.json
// this example is just showing the usage of "Machines_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeMachineResource created on azure
// for more information of creating HybridComputeMachineResource, please refer to the document of HybridComputeMachineResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string machineName = "myMachine";
ResourceIdentifier hybridComputeMachineResourceId = HybridComputeMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName);
HybridComputeMachineResource hybridComputeMachine = client.GetHybridComputeMachineResource(hybridComputeMachineResourceId);

// invoke the operation
HybridComputeMachinePatch patch = new HybridComputeMachinePatch()
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    LocationData = new HybridComputeLocation("Redmond"),
    OSProfile = new HybridComputeOSProfile()
    {
        WindowsConfiguration = new HybridComputeWindowsConfiguration()
        {
            AssessmentMode = AssessmentModeType.ImageDefault,
            PatchMode = PatchModeType.AutomaticByPlatform,
            IsHotpatchingEnabled = true,
        },
        LinuxConfiguration = new HybridComputeLinuxConfiguration()
        {
            AssessmentMode = AssessmentModeType.ImageDefault,
            PatchMode = PatchModeType.Manual,
        },
    },
    ParentClusterResourceId = new ResourceIdentifier("{AzureStackHCIResourceId}"),
    PrivateLinkScopeResourceId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
};
HybridComputeMachineResource result = await hybridComputeMachine.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridComputeMachineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
