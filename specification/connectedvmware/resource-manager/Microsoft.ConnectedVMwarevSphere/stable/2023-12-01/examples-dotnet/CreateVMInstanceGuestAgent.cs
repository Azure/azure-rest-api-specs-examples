using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.ConnectedVMwarevSphere;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-12-01/examples/CreateVMInstanceGuestAgent.json
// this example is just showing the usage of "VMInstanceGuestAgents_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VmInstanceGuestAgentResource created on azure
// for more information of creating VmInstanceGuestAgentResource, please refer to the document of VmInstanceGuestAgentResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
ResourceIdentifier vmInstanceGuestAgentResourceId = VmInstanceGuestAgentResource.CreateResourceIdentifier(resourceUri);
VmInstanceGuestAgentResource vmInstanceGuestAgent = client.GetVmInstanceGuestAgentResource(vmInstanceGuestAgentResourceId);

// invoke the operation
VmInstanceGuestAgentData data = new VmInstanceGuestAgentData
{
    Credentials = new VmInstanceGuestCredential
    {
        Username = "tempuser",
        Password = "<password>",
    },
    PrivateLinkScopeResourceId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
    HttpsProxy = "http://192.1.2.3:8080",
    ProvisioningAction = GuestAgentProvisioningAction.Install,
};
ArmOperation<VmInstanceGuestAgentResource> lro = await vmInstanceGuestAgent.CreateOrUpdateAsync(WaitUntil.Completed, data);
VmInstanceGuestAgentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VmInstanceGuestAgentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
