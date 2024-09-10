using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/dataConnectorDefinitions/CreateCustomizableDataConnectorDefinition.json
// this example is just showing the usage of "DataConnectorDefinitions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this SecurityInsightsDataConnectorDefinitionResource
SecurityInsightsDataConnectorDefinitionCollection collection = operationalInsightsWorkspaceSecurityInsights.GetSecurityInsightsDataConnectorDefinitions();

// invoke the operation
string dataConnectorDefinitionName = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
SecurityInsightsDataConnectorDefinitionData data = new CustomizableConnectorDefinitionData()
{
    ConnectorUiConfig = new CustomizableConnectorUiConfig("GitHub Enterprise Audit Log", "GitHub", "The GitHub audit log connector provides the capability to ingest GitHub logs into Azure Sentinel. By connecting GitHub audit logs into Azure Sentinel, you can view this data in workbooks, use it to create custom alerts, and improve your investigation process.", new ConnectorGraphQuery[]
{
new ConnectorGraphQuery("Total events received","GitHub audit log events","GitHubAuditLogPolling_CL")
}, new ConnectorDataType[]
{
new ConnectorDataType("GitHubAuditLogPolling_CL","GitHubAuditLogPolling_CL \n            | summarize Time = max(TimeGenerated)\n            | where isnotempty(Time)")
}, new ConnectorConnectivityCriterion[]
{
new ConnectorConnectivityCriterion("IsConnectedQuery")
{
Value =
{
"GitHubAuditLogPolling_CL \n | summarize LastLogReceived = max(TimeGenerated)\n | project IsConnected = LastLogReceived > ago(30d)"
},
}
}, new ConnectorDefinitionsPermissions()
{
    ResourceProvider =
{
new ConnectorDefinitionsResourceProvider("Microsoft.OperationalInsights/workspaces","read and write permissions are required.","Workspace",ProviderPermissionsScope.Workspace,new ConnectorResourceProviderRequiredPermissions()
{
IsReadAction = false,
IsWriteAction = true,
IsDeleteAction = false,
IsCustomAction = false,
})
},
    Customs =
{
new CustomPermissionDetails("GitHub API personal token Key","You need access to GitHub personal token, the key should have 'admin:org' scope")
},
}, new InstructionStep[]
{
new InstructionStep()
{
Title = "Connect GitHub Enterprise Audit Log to Azure Sentinel",
Description = "Enable GitHub audit Logs. \n Follow [this](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) to create or find your personal key",
Instructions =
{
new InstructionStepDetails(BinaryData.FromObjectAsJson(new Dictionary<string, object>()
{
["clientIdLabel"] = "Client ID",
["clientSecretLabel"] = "Client Secret",
["connectButtonLabel"] = "Connect",
["disconnectButtonLabel"] = "Disconnect"}),"OAuthForm")
},
}
})
    {
        Availability = new ConnectorDefinitionsAvailability()
        {
            Status = 1,
            IsPreview = false,
        },
    },
    ETag = new ETag("\"0300bf09-0000-0000-0000-5c37296e0000\""),
};
ArmOperation<SecurityInsightsDataConnectorDefinitionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataConnectorDefinitionName, data);
SecurityInsightsDataConnectorDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsDataConnectorDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
