using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/PutAlertsSuppressionRule_example.json
// this example is just showing the usage of "AlertsSuppressionRules_Update" operation, for the dependent resources, they will have to be created separately.

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
SecurityAlertsSuppressionRuleData data = new SecurityAlertsSuppressionRuleData
{
    AlertType = "IpAnomaly",
    ExpireOn = DateTimeOffset.Parse("2019-12-01T19:50:47.083633Z"),
    Reason = "FalsePositive",
    State = SecurityAlertsSuppressionRuleState.Enabled,
    Comment = "Test VM",
    SuppressionAlertsScopeAllOf = {new SuppressionAlertsScopeElement
    {
    Field = "entities.ip.address",
    }, new SuppressionAlertsScopeElement
    {
    Field = "entities.process.commandline",
    }},
};
ArmOperation<SecurityAlertsSuppressionRuleResource> lro = await securityAlertsSuppressionRule.UpdateAsync(WaitUntil.Completed, data);
SecurityAlertsSuppressionRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityAlertsSuppressionRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
