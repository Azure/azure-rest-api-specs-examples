using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/workspaceManagerAssignments/DeleteWorkspaceManagerAssignment.json
// this example is just showing the usage of "WorkspaceManagerAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkspaceManagerAssignmentResource created on azure
// for more information of creating WorkspaceManagerAssignmentResource, please refer to the document of WorkspaceManagerAssignmentResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string workspaceManagerAssignmentName = "47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58";
ResourceIdentifier workspaceManagerAssignmentResourceId = WorkspaceManagerAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, workspaceManagerAssignmentName);
WorkspaceManagerAssignmentResource workspaceManagerAssignment = client.GetWorkspaceManagerAssignmentResource(workspaceManagerAssignmentResourceId);

// invoke the operation
await workspaceManagerAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
