using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/workspaceManagerAssignments/CreateOrUpdateWorkspaceManagerAssignment.json
// this example is just showing the usage of "WorkspaceManagerAssignments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
WorkspaceManagerAssignmentData data = new WorkspaceManagerAssignmentData()
{
    TargetResourceName = "37207a7a-3b8a-438f-a559-c7df400e1b96",
    Items =
    {
    new WorkspaceManagerAssignmentItem()
    {
    ResourceId = new ResourceIdentifier("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspac-es/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExampleOne"),
    },new WorkspaceManagerAssignmentItem()
    {
    ResourceId = new ResourceIdentifier("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspac-es/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExampleTwo"),
    }
    },
};
ArmOperation<WorkspaceManagerAssignmentResource> lro = await workspaceManagerAssignment.UpdateAsync(WaitUntil.Completed, data);
WorkspaceManagerAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkspaceManagerAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
