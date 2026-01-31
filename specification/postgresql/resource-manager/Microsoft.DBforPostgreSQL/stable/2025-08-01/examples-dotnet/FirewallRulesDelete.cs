using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/FirewallRulesDelete.json
// this example is just showing the usage of "FirewallRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerFirewallRuleResource created on azure
// for more information of creating PostgreSqlFlexibleServerFirewallRuleResource, please refer to the document of PostgreSqlFlexibleServerFirewallRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "exampleresourcegroup";
string serverName = "exampleserver";
string firewallRuleName = "examplefirewallrule";
ResourceIdentifier postgreSqlFlexibleServerFirewallRuleResourceId = PostgreSqlFlexibleServerFirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, firewallRuleName);
PostgreSqlFlexibleServerFirewallRuleResource postgreSqlFlexibleServerFirewallRule = client.GetPostgreSqlFlexibleServerFirewallRuleResource(postgreSqlFlexibleServerFirewallRuleResourceId);

// invoke the operation
await postgreSqlFlexibleServerFirewallRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
