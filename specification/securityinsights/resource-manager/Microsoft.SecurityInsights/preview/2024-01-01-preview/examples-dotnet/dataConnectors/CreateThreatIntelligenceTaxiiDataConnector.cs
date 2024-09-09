using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/dataConnectors/CreateThreatIntelligenceTaxiiDataConnector.json
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
string dataConnectorId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
SecurityInsightsDataConnectorData data = new ThreatIntelligenceTaxiiDataConnector()
{
    TenantId = Guid.Parse("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
    WorkspaceId = "dd124572-4962-4495-9bd2-9dade12314b4",
    FriendlyName = "testTaxii",
    TaxiiServer = "https://limo.anomali.com/api/v1/taxii2/feeds",
    CollectionId = "135",
    UserName = "--",
    Password = "--",
    TaxiiLookbackPeriod = DateTimeOffset.Parse("2020-01-01T13:00:30.123Z"),
    PollingFrequency = PollingFrequency.OnceADay,
    TaxiiClientState = SecurityInsightsDataTypeConnectionState.Enabled,
    ETag = new ETag("d12423f6-a60b-4ca5-88c0-feb1a182d0f0"),
};
ArmOperation<SecurityInsightsDataConnectorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataConnectorId, data);
SecurityInsightsDataConnectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsDataConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
