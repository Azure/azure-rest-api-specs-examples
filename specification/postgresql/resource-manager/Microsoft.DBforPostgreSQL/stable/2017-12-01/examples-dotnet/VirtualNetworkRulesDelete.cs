using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PostgreSql;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/VirtualNetworkRulesDelete.json
// this example is just showing the usage of "VirtualNetworkRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlVirtualNetworkRuleResource created on azure
// for more information of creating PostgreSqlVirtualNetworkRuleResource, please refer to the document of PostgreSqlVirtualNetworkRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "vnet-test-svr";
string virtualNetworkRuleName = "vnet-firewall-rule";
ResourceIdentifier postgreSqlVirtualNetworkRuleResourceId = PostgreSqlVirtualNetworkRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, virtualNetworkRuleName);
PostgreSqlVirtualNetworkRuleResource postgreSqlVirtualNetworkRule = client.GetPostgreSqlVirtualNetworkRuleResource(postgreSqlVirtualNetworkRuleResourceId);

// invoke the operation
await postgreSqlVirtualNetworkRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
