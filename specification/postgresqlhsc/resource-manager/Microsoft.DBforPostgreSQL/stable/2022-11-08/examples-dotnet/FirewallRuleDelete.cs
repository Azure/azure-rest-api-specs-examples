using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDBForPostgreSql;

// Generated from example definition: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/FirewallRuleDelete.json
// this example is just showing the usage of "FirewallRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBForPostgreSqlFirewallRuleResource created on azure
// for more information of creating CosmosDBForPostgreSqlFirewallRuleResource, please refer to the document of CosmosDBForPostgreSqlFirewallRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string clusterName = "pgtestsvc4";
string firewallRuleName = "rule1";
ResourceIdentifier cosmosDBForPostgreSqlFirewallRuleResourceId = CosmosDBForPostgreSqlFirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, firewallRuleName);
CosmosDBForPostgreSqlFirewallRuleResource cosmosDBForPostgreSqlFirewallRule = client.GetCosmosDBForPostgreSqlFirewallRuleResource(cosmosDBForPostgreSqlFirewallRuleResourceId);

// invoke the operation
await cosmosDBForPostgreSqlFirewallRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
