using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/workspaceManagerGroups/CreateOrUpdateWorkspaceManagerGroup.json
// this example is just showing the usage of "WorkspaceManagerGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsWorkspaceSecurityInsightsResource created on azure
// for more information of creating OperationalInsightsWorkspaceSecurityInsightsResource, please refer to the document of OperationalInsightsWorkspaceSecurityInsightsResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
ResourceIdentifier operationalInsightsWorkspaceSecurityInsightsResourceId = OperationalInsightsWorkspaceSecurityInsightsResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
OperationalInsightsWorkspaceSecurityInsightsResource operationalInsightsWorkspaceSecurityInsights = client.GetOperationalInsightsWorkspaceSecurityInsightsResource(operationalInsightsWorkspaceSecurityInsightsResourceId);

// get the collection of this WorkspaceManagerGroupResource
WorkspaceManagerGroupCollection collection = operationalInsightsWorkspaceSecurityInsights.GetWorkspaceManagerGroups();

// invoke the operation
string workspaceManagerGroupName = "37207a7a-3b8a-438f-a559-c7df400e1b96";
WorkspaceManagerGroupData data = new WorkspaceManagerGroupData()
{
    Description = "Group of all financial and banking institutions",
    DisplayName = "Banks",
    MemberResourceNames =
    {
    "afbd324f-6c48-459c-8710-8d1e1cd03812","f5fa104e-c0e3-4747-9182-d342dc048a9e"
    },
};
ArmOperation<WorkspaceManagerGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workspaceManagerGroupName, data);
WorkspaceManagerGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkspaceManagerGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
