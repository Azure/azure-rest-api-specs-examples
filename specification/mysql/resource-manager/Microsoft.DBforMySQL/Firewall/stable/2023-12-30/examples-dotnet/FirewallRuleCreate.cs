using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/Firewall/stable/2023-12-30/examples/FirewallRuleCreate.json
// this example is just showing the usage of "FirewallRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerFirewallRuleResource created on azure
// for more information of creating MySqlFlexibleServerFirewallRuleResource, please refer to the document of MySqlFlexibleServerFirewallRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "testserver";
string firewallRuleName = "rule1";
ResourceIdentifier mySqlFlexibleServerFirewallRuleResourceId = MySqlFlexibleServerFirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, firewallRuleName);
MySqlFlexibleServerFirewallRuleResource mySqlFlexibleServerFirewallRule = client.GetMySqlFlexibleServerFirewallRuleResource(mySqlFlexibleServerFirewallRuleResourceId);

// invoke the operation
MySqlFlexibleServerFirewallRuleData data = new MySqlFlexibleServerFirewallRuleData(IPAddress.Parse("0.0.0.0"), IPAddress.Parse("255.255.255.255"));
ArmOperation<MySqlFlexibleServerFirewallRuleResource> lro = await mySqlFlexibleServerFirewallRule.UpdateAsync(WaitUntil.Completed, data);
MySqlFlexibleServerFirewallRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlFlexibleServerFirewallRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
