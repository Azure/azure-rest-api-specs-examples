using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/DeleteAlertsSuppressionRule_example.json
// this example is just showing the usage of "AlertsSuppressionRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityAlertsSuppressionRuleResource created on azure
// for more information of creating SecurityAlertsSuppressionRuleResource, please refer to the document of SecurityAlertsSuppressionRuleResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string alertsSuppressionRuleName = "dismissIpAnomalyAlerts";
ResourceIdentifier securityAlertsSuppressionRuleResourceId = SecurityAlertsSuppressionRuleResource.CreateResourceIdentifier(subscriptionId, alertsSuppressionRuleName);
SecurityAlertsSuppressionRuleResource securityAlertsSuppressionRule = client.GetSecurityAlertsSuppressionRuleResource(securityAlertsSuppressionRuleResourceId);

// invoke the operation
await securityAlertsSuppressionRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
