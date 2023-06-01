using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FirewallRuleDelete.json
// this example is just showing the usage of "FirewallRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlFirewallRuleResource created on azure
// for more information of creating SqlFirewallRuleResource, please refer to the document of SqlFirewallRuleResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "firewallrulecrudtest-9886";
string serverName = "firewallrulecrudtest-2368";
string firewallRuleName = "firewallrulecrudtest-7011";
ResourceIdentifier sqlFirewallRuleResourceId = SqlFirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, firewallRuleName);
SqlFirewallRuleResource sqlFirewallRule = client.GetSqlFirewallRuleResource(sqlFirewallRuleResourceId);

// invoke the operation
await sqlFirewallRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
