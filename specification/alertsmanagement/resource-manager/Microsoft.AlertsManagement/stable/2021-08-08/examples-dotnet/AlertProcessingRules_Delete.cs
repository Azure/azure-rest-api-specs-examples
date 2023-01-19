using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AlertsManagement;
using Azure.ResourceManager.AlertsManagement.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Delete.json
// this example is just showing the usage of "AlertProcessingRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AlertProcessingRuleResource created on azure
// for more information of creating AlertProcessingRuleResource, please refer to the document of AlertProcessingRuleResource
string subscriptionId = "1e3ff1c0-771a-4119-a03b-be82a51e232d";
string resourceGroupName = "alertscorrelationrg";
string alertProcessingRuleName = "DailySuppression";
ResourceIdentifier alertProcessingRuleResourceId = AlertProcessingRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, alertProcessingRuleName);
AlertProcessingRuleResource alertProcessingRule = client.GetAlertProcessingRuleResource(alertProcessingRuleResourceId);

// invoke the operation
await alertProcessingRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
