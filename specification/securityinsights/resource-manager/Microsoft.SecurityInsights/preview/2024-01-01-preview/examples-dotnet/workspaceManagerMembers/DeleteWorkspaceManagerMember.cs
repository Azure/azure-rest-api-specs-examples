using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/workspaceManagerMembers/DeleteWorkspaceManagerMember.json
// this example is just showing the usage of "WorkspaceManagerMembers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkspaceManagerMemberResource created on azure
// for more information of creating WorkspaceManagerMemberResource, please refer to the document of WorkspaceManagerMemberResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string workspaceManagerMemberName = "afbd324f-6c48-459c-8710-8d1e1cd03812";
ResourceIdentifier workspaceManagerMemberResourceId = WorkspaceManagerMemberResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, workspaceManagerMemberName);
WorkspaceManagerMemberResource workspaceManagerMember = client.GetWorkspaceManagerMemberResource(workspaceManagerMemberResourceId);

// invoke the operation
await workspaceManagerMember.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
