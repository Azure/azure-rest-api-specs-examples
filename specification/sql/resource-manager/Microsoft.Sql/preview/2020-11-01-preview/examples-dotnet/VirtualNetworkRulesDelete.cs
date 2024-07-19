using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualNetworkRulesDelete.json
// this example is just showing the usage of "VirtualNetworkRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerVirtualNetworkRuleResource created on azure
// for more information of creating SqlServerVirtualNetworkRuleResource, please refer to the document of SqlServerVirtualNetworkRuleResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "vnet-test-svr";
string virtualNetworkRuleName = "vnet-firewall-rule";
ResourceIdentifier sqlServerVirtualNetworkRuleResourceId = SqlServerVirtualNetworkRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, virtualNetworkRuleName);
SqlServerVirtualNetworkRuleResource sqlServerVirtualNetworkRule = client.GetSqlServerVirtualNetworkRuleResource(sqlServerVirtualNetworkRuleResourceId);

// invoke the operation
await sqlServerVirtualNetworkRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
