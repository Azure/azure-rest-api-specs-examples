using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualNetworkRulesGet.json
// this example is just showing the usage of "VirtualNetworkRules_Get" operation, for the dependent resources, they will have to be created separately.

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
SqlServerVirtualNetworkRuleResource result = await sqlServerVirtualNetworkRule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlServerVirtualNetworkRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
