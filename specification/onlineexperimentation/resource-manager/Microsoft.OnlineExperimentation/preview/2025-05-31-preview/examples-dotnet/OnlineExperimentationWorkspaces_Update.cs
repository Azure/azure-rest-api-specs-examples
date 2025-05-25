using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.OnlineExperimentation.Models;
using Azure.ResourceManager.OnlineExperimentation;

// Generated from example definition: 2025-05-31-preview/OnlineExperimentationWorkspaces_Update.json
// this example is just showing the usage of "OnlineExperimentationWorkspace_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OnlineExperimentationWorkspaceResource created on azure
// for more information of creating OnlineExperimentationWorkspaceResource, please refer to the document of OnlineExperimentationWorkspaceResource
string subscriptionId = "fa5fc227-a624-475e-b696-cdd604c735bc";
string resourceGroupName = "res9871";
string workspaceName = "expworkspace3";
ResourceIdentifier onlineExperimentationWorkspaceResourceId = OnlineExperimentationWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
OnlineExperimentationWorkspaceResource onlineExperimentationWorkspace = client.GetOnlineExperimentationWorkspaceResource(onlineExperimentationWorkspaceResourceId);

// invoke the operation
OnlineExperimentationWorkspacePatch patch = new OnlineExperimentationWorkspacePatch
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity(),
        [new ResourceIdentifier("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["newKey"] = "newVal"
    },
};
ArmOperation<OnlineExperimentationWorkspaceResource> lro = await onlineExperimentationWorkspace.UpdateAsync(WaitUntil.Completed, patch);
OnlineExperimentationWorkspaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OnlineExperimentationWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
