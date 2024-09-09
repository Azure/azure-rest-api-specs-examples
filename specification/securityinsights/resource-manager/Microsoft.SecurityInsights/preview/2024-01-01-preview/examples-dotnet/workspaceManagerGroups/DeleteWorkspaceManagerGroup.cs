using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/workspaceManagerGroups/DeleteWorkspaceManagerGroup.json
// this example is just showing the usage of "WorkspaceManagerGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkspaceManagerGroupResource created on azure
// for more information of creating WorkspaceManagerGroupResource, please refer to the document of WorkspaceManagerGroupResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string workspaceManagerGroupName = "37207a7a-3b8a-438f-a559-c7df400e1b96";
ResourceIdentifier workspaceManagerGroupResourceId = WorkspaceManagerGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, workspaceManagerGroupName);
WorkspaceManagerGroupResource workspaceManagerGroup = client.GetWorkspaceManagerGroupResource(workspaceManagerGroupResourceId);

// invoke the operation
await workspaceManagerGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
