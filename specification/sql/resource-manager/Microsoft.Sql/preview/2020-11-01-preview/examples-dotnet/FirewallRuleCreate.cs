using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FirewallRuleCreate.json
// this example is just showing the usage of "FirewallRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlFirewallRuleResource created on azure
// for more information of creating SqlFirewallRuleResource, please refer to the document of SqlFirewallRuleResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "firewallrulecrudtest-12";
string serverName = "firewallrulecrudtest-6285";
string firewallRuleName = "firewallrulecrudtest-5370";
ResourceIdentifier sqlFirewallRuleResourceId = SqlFirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, firewallRuleName);
SqlFirewallRuleResource sqlFirewallRule = client.GetSqlFirewallRuleResource(sqlFirewallRuleResourceId);

// invoke the operation
SqlFirewallRuleData data = new SqlFirewallRuleData()
{
    StartIPAddress = "0.0.0.3",
    EndIPAddress = "0.0.0.3",
};
ArmOperation<SqlFirewallRuleResource> lro = await sqlFirewallRule.UpdateAsync(WaitUntil.Completed, data);
SqlFirewallRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlFirewallRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
