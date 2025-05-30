using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/VirtualNetworkRulesCreateOrUpdate.json
// this example is just showing the usage of "VirtualNetworkRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlVirtualNetworkRuleResource created on azure
// for more information of creating MySqlVirtualNetworkRuleResource, please refer to the document of MySqlVirtualNetworkRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "vnet-test-svr";
string virtualNetworkRuleName = "vnet-firewall-rule";
ResourceIdentifier mySqlVirtualNetworkRuleResourceId = MySqlVirtualNetworkRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, virtualNetworkRuleName);
MySqlVirtualNetworkRuleResource mySqlVirtualNetworkRule = client.GetMySqlVirtualNetworkRuleResource(mySqlVirtualNetworkRuleResourceId);

// invoke the operation
MySqlVirtualNetworkRuleData data = new MySqlVirtualNetworkRuleData
{
    VirtualNetworkSubnetId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
    IgnoreMissingVnetServiceEndpoint = false,
};
ArmOperation<MySqlVirtualNetworkRuleResource> lro = await mySqlVirtualNetworkRule.UpdateAsync(WaitUntil.Completed, data);
MySqlVirtualNetworkRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlVirtualNetworkRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
