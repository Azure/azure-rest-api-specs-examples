using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/PrivateEndpointConnection_DeleteByWorkspace.json
// this example is just showing the usage of "PrivateEndpointConnections_DeleteByWorkspace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkspacePrivateEndpointConnectionResource created on azure
// for more information of creating WorkspacePrivateEndpointConnectionResource, please refer to the document of WorkspacePrivateEndpointConnectionResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string workspaceName = "workspace1";
string privateEndpointConnectionName = "workspace1.377103f1-5179-4bdf-8556-4cdd3207cc5b";
ResourceIdentifier workspacePrivateEndpointConnectionResourceId = WorkspacePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, privateEndpointConnectionName);
WorkspacePrivateEndpointConnectionResource workspacePrivateEndpointConnection = client.GetWorkspacePrivateEndpointConnectionResource(workspacePrivateEndpointConnectionResourceId);

// invoke the operation
await workspacePrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
