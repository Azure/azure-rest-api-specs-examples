using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/FirewallRuleGet.json
// this example is just showing the usage of "FirewallRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerResource created on azure
// for more information of creating PostgreSqlFlexibleServerResource, please refer to the document of PostgreSqlFlexibleServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "testserver";
ResourceIdentifier postgreSqlFlexibleServerResourceId = PostgreSqlFlexibleServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
PostgreSqlFlexibleServerResource postgreSqlFlexibleServer = client.GetPostgreSqlFlexibleServerResource(postgreSqlFlexibleServerResourceId);

// get the collection of this PostgreSqlFlexibleServerFirewallRuleResource
PostgreSqlFlexibleServerFirewallRuleCollection collection = postgreSqlFlexibleServer.GetPostgreSqlFlexibleServerFirewallRules();

// invoke the operation
string firewallRuleName = "rule1";
NullableResponse<PostgreSqlFlexibleServerFirewallRuleResource> response = await collection.GetIfExistsAsync(firewallRuleName);
PostgreSqlFlexibleServerFirewallRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PostgreSqlFlexibleServerFirewallRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
