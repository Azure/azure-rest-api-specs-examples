using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/Workspace_Update.json
// this example is just showing the usage of "Workspaces_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualWorkspaceResource created on azure
// for more information of creating VirtualWorkspaceResource, please refer to the document of VirtualWorkspaceResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string workspaceName = "workspace1";
ResourceIdentifier virtualWorkspaceResourceId = VirtualWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
VirtualWorkspaceResource virtualWorkspace = client.GetVirtualWorkspaceResource(virtualWorkspaceResourceId);

// invoke the operation
VirtualWorkspacePatch patch = new VirtualWorkspacePatch
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
    Description = "des1",
    FriendlyName = "friendly",
};
VirtualWorkspaceResource result = await virtualWorkspace.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
