using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/incidents/IncidentAlerts/Incidents_ListAlerts.json
// this example is just showing the usage of "Incidents_ListAlerts" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsIncidentResource created on azure
// for more information of creating SecurityInsightsIncidentResource, please refer to the document of SecurityInsightsIncidentResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string incidentId = "69a30280-6a4c-4aa7-9af0-5d63f335d600";
ResourceIdentifier securityInsightsIncidentResourceId = SecurityInsightsIncidentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, incidentId);
SecurityInsightsIncidentResource securityInsightsIncident = client.GetSecurityInsightsIncidentResource(securityInsightsIncidentResourceId);

// invoke the operation and iterate over the result
await foreach (SecurityInsightsAlert item in securityInsightsIncident.GetAlertsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
