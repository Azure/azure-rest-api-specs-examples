using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/VirtualNetworkRulesGet.json
// this example is just showing the usage of "VirtualNetworkRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlServerResource created on azure
// for more information of creating PostgreSqlServerResource, please refer to the document of PostgreSqlServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "vnet-test-svr";
ResourceIdentifier postgreSqlServerResourceId = PostgreSqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
PostgreSqlServerResource postgreSqlServer = client.GetPostgreSqlServerResource(postgreSqlServerResourceId);

// get the collection of this PostgreSqlVirtualNetworkRuleResource
PostgreSqlVirtualNetworkRuleCollection collection = postgreSqlServer.GetPostgreSqlVirtualNetworkRules();

// invoke the operation
string virtualNetworkRuleName = "vnet-firewall-rule";
NullableResponse<PostgreSqlVirtualNetworkRuleResource> response = await collection.GetIfExistsAsync(virtualNetworkRuleName);
PostgreSqlVirtualNetworkRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PostgreSqlVirtualNetworkRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
