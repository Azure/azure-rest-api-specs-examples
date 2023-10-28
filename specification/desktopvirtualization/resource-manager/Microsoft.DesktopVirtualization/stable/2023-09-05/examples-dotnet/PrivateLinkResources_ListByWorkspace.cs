using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/PrivateLinkResources_ListByWorkspace.json
// this example is just showing the usage of "PrivateLinkResources_ListByWorkspace" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (DesktopVirtualizationPrivateLinkResourceData item in virtualWorkspace.GetPrivateLinkResourcesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
