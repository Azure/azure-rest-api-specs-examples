using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/FirewallRuleDelete.json
// this example is just showing the usage of "FirewallRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFirewallRuleResource created on azure
// for more information of creating PostgreSqlFirewallRuleResource, please refer to the document of PostgreSqlFirewallRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "testserver";
string firewallRuleName = "rule1";
ResourceIdentifier postgreSqlFirewallRuleResourceId = PostgreSqlFirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, firewallRuleName);
PostgreSqlFirewallRuleResource postgreSqlFirewallRule = client.GetPostgreSqlFirewallRuleResource(postgreSqlFirewallRuleResourceId);

// invoke the operation
await postgreSqlFirewallRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
