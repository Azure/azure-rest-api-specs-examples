using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/FirewallRuleGet.json
// this example is just showing the usage of "FirewallRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerResource created on azure
// for more information of creating SqlServerResource, please refer to the document of SqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "firewallrulecrudtest-12";
string serverName = "firewallrulecrudtest-6285";
ResourceIdentifier sqlServerResourceId = SqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
SqlServerResource sqlServer = client.GetSqlServerResource(sqlServerResourceId);

// get the collection of this SqlFirewallRuleResource
SqlFirewallRuleCollection collection = sqlServer.GetSqlFirewallRules();

// invoke the operation
string firewallRuleName = "firewallrulecrudtest-2304";
NullableResponse<SqlFirewallRuleResource> response = await collection.GetIfExistsAsync(firewallRuleName);
SqlFirewallRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SqlFirewallRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
