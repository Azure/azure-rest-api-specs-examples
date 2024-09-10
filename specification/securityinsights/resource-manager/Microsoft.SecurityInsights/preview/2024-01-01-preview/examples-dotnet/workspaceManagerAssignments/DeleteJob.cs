using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/workspaceManagerAssignments/DeleteJob.json
// this example is just showing the usage of "WorkspaceManagerAssignmentJobs_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkspaceManagerAssignmentJobResource created on azure
// for more information of creating WorkspaceManagerAssignmentJobResource, please refer to the document of WorkspaceManagerAssignmentJobResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string workspaceManagerAssignmentName = "47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58";
string jobName = "cfbe1338-8276-4d5d-8b96-931117f9fa0e";
ResourceIdentifier workspaceManagerAssignmentJobResourceId = WorkspaceManagerAssignmentJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, workspaceManagerAssignmentName, jobName);
WorkspaceManagerAssignmentJobResource workspaceManagerAssignmentJob = client.GetWorkspaceManagerAssignmentJobResource(workspaceManagerAssignmentJobResourceId);

// invoke the operation
await workspaceManagerAssignmentJob.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
