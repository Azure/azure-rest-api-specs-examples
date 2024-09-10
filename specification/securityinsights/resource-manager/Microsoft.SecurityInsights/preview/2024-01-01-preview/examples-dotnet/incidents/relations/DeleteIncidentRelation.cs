using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/incidents/relations/DeleteIncidentRelation.json
// this example is just showing the usage of "IncidentRelations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsIncidentRelationResource created on azure
// for more information of creating SecurityInsightsIncidentRelationResource, please refer to the document of SecurityInsightsIncidentRelationResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string incidentId = "afbd324f-6c48-459c-8710-8d1e1cd03812";
string relationName = "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014";
ResourceIdentifier securityInsightsIncidentRelationResourceId = SecurityInsightsIncidentRelationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, incidentId, relationName);
SecurityInsightsIncidentRelationResource securityInsightsIncidentRelation = client.GetSecurityInsightsIncidentRelationResource(securityInsightsIncidentRelationResourceId);

// invoke the operation
await securityInsightsIncidentRelation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
