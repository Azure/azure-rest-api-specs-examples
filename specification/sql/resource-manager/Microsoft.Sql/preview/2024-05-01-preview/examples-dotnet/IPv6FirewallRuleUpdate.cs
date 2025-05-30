using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/IPv6FirewallRuleUpdate.json
// this example is just showing the usage of "IPv6FirewallRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IPv6FirewallRuleResource created on azure
// for more information of creating IPv6FirewallRuleResource, please refer to the document of IPv6FirewallRuleResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "firewallrulecrudtest-12";
string serverName = "firewallrulecrudtest-6285";
string firewallRuleName = "firewallrulecrudtest-3927";
ResourceIdentifier iPv6FirewallRuleResourceId = IPv6FirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, firewallRuleName);
IPv6FirewallRuleResource iPv6FirewallRule = client.GetIPv6FirewallRuleResource(iPv6FirewallRuleResourceId);

// invoke the operation
IPv6FirewallRuleData data = new IPv6FirewallRuleData
{
    StartIPv6Address = "0000:0000:0000:0000:0000:ffff:0000:0001",
    EndIPv6Address = "0000:0000:0000:0000:0000:ffff:0000:0001",
};
ArmOperation<IPv6FirewallRuleResource> lro = await iPv6FirewallRule.UpdateAsync(WaitUntil.Completed, data);
IPv6FirewallRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IPv6FirewallRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
