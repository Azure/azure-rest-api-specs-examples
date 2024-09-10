using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/dataConnectors/CreateGoogleCloudPlatform.json
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
string dataConnectorId = "GCP_fce27b90-d6f5-4d30-991a-af509a2b50a1";
SecurityInsightsDataConnectorData data = new GcpDataConnector()
{
    ConnectorDefinitionName = "GcpConnector",
    Auth = new GcpAuthProperties("sentinel-service-account@project-id.iam.gserviceaccount.com", "123456789012", "sentinel-identity-provider"),
    Request = new GcpRequestProperties("project-id", new string[]
{
"sentinel-subscription"
}),
    DcrConfig = new DcrConfiguration("https://microsoft-sentinel-datacollectionendpoint-123m.westeurope-1.ingest.monitor.azure.com", "dcr-de21b053bd5a44beb99a256c9db85023", "SENTINEL_GCP_AUDIT_LOGS"),
};
ArmOperation<SecurityInsightsDataConnectorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataConnectorId, data);
SecurityInsightsDataConnectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsDataConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
