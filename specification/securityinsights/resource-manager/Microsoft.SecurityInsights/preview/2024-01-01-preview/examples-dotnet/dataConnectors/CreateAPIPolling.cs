using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/dataConnectors/CreateAPIPolling.json
// this example is just showing the usage of "DataConnectors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this SecurityInsightsDataConnectorResource
SecurityInsightsDataConnectorCollection collection = operationalInsightsWorkspaceSecurityInsights.GetSecurityInsightsDataConnectors();

// invoke the operation
string dataConnectorId = "316ec55e-7138-4d63-ab18-90c8a60fd1c8";
SecurityInsightsDataConnectorData data = new CodelessApiPollingDataConnector()
{
    ConnectorUiConfig = new CodelessUiConnectorConfigProperties("GitHub Enterprise Audit Log", "GitHub", "The GitHub audit log connector provides the capability to ingest GitHub logs into Azure Sentinel. By connecting GitHub audit logs into Azure Sentinel, you can view this data in workbooks, use it to create custom alerts, and improve your investigation process.", "GitHubAuditLogPolling_CL", new ConnectorGraphQueries[]
{
new ConnectorGraphQueries()
{
MetricName = "Total events received",
Legend = "GitHub audit log events",
BaseQuery = "{{graphQueriesTableName}}",
}
}, new SourceControlSampleQueries[]
{
new SourceControlSampleQueries()
{
Description = "All logs",
Query = "{{graphQueriesTableName}}\n | take 10 <change>",
}
}, new LastDataReceivedDataType[]
{
new LastDataReceivedDataType()
{
Name = "{{graphQueriesTableName}}",
LastDataReceivedQuery = "{{graphQueriesTableName}}\n            | summarize Time = max(TimeGenerated)\n            | where isnotempty(Time)",
}
}, new ConnectorConnectivityCriteria[]
{
new ConnectorConnectivityCriteria()
{
ConnectivityType = new ConnectorConnectivityType("SentinelKindsV2"),
Value =
{
},
}
}, new ConnectorAvailability()
{
    Status = ConnectorAvailabilityStatus._1,
    IsPreview = true,
}, new ConnectorPermissions()
{
    ResourceProvider =
{
new ConnectorResourceProvider()
{
Provider = ConnectorProviderName.MicrosoftOperationalInsightsWorkspaces,
PermissionsDisplayText = "read and write permissions are required.",
ProviderDisplayName = "Workspace",
Scope = PermissionProviderScope.Workspace,
RequiredPermissions = new ConnectorRequiredPermissions()
{
IsWriteAction = true,
IsReadAction = true,
IsDeleteAction = true,
},
}
},
    Customs =
{
new ConnectorCustoms()
{
Name = "GitHub API personal token Key",
Description = "You need access to GitHub personal token, the key should have 'admin:org' scope",
}
},
}, new InstructionSteps[]
{
new InstructionSteps()
{
Title = "Connect GitHub Enterprise Audit Log to Azure Sentinel",
Description = "Enable GitHub audit Logs. \n Follow [this](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) to create or find your personal key",
Instructions =
{
new ConnectorInstructionModelBase(new ConnectorSettingType("APIKey"))
{
Parameters = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
{
["enable"] = "true",
["userRequestPlaceHoldersInput"] = new object[] { new Dictionary<string, object>()
{
["displayText"] = "Organization Name",
["placeHolderName"] = "{{placeHolder1}}",
["placeHolderValue"] = "",
["requestObjectKey"] = "apiEndpoint"} }}),
}
},
}
}),
    PollingConfig = new CodelessConnectorPollingConfigProperties(new CodelessConnectorPollingAuthProperties("APIKey")
    {
        ApiKeyName = "Authorization",
        ApiKeyIdentifier = "token",
    }, new CodelessConnectorPollingRequestProperties("https://api.github.com/organizations/{{placeHolder1}}/audit-log", 15, "Get", "yyyy-MM-ddTHH:mm:ssZ")
    {
        RateLimitQps = 50,
        RetryCount = 2,
        TimeoutInSeconds = 60,
        Headers = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
        {
            ["Accept"] = "application/json",
            ["User-Agent"] = "Scuba"
        }),
        QueryParameters = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
        {
            ["phrase"] = "created:{_QueryWindowStartTime}..{_QueryWindowEndTime}"
        }),
    })
    {
        Paging = new CodelessConnectorPollingPagingProperties("LinkHeader")
        {
            PageSizeParaName = "per_page",
        },
        Response = new CodelessConnectorPollingResponseProperties(new string[]
{
"$"
}),
    },
};
ArmOperation<SecurityInsightsDataConnectorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataConnectorId, data);
SecurityInsightsDataConnectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsDataConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
