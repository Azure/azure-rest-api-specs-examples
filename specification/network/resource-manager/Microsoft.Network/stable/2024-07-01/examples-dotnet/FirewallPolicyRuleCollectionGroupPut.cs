using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/FirewallPolicyRuleCollectionGroupPut.json
// this example is just showing the usage of "FirewallPolicyRuleCollectionGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirewallPolicyResource created on azure
// for more information of creating FirewallPolicyResource, please refer to the document of FirewallPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string firewallPolicyName = "firewallPolicy";
ResourceIdentifier firewallPolicyResourceId = FirewallPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallPolicyName);
FirewallPolicyResource firewallPolicy = client.GetFirewallPolicyResource(firewallPolicyResourceId);

// get the collection of this FirewallPolicyRuleCollectionGroupResource
FirewallPolicyRuleCollectionGroupCollection collection = firewallPolicy.GetFirewallPolicyRuleCollectionGroups();

// invoke the operation
string ruleCollectionGroupName = "ruleCollectionGroup1";
FirewallPolicyRuleCollectionGroupData data = new FirewallPolicyRuleCollectionGroupData
{
    Priority = 100,
    RuleCollections = {new FirewallPolicyFilterRuleCollectionInfo
    {
    ActionType = FirewallPolicyFilterRuleCollectionActionType.Deny,
    Rules = {new NetworkRule
    {
    IPProtocols = {FirewallPolicyRuleNetworkProtocol.Tcp},
    SourceAddresses = {"10.1.25.0/24"},
    DestinationAddresses = {"*"},
    DestinationPorts = {"*"},
    Name = "network-rule1",
    }},
    Name = "Example-Filter-Rule-Collection",
    Priority = 100,
    }},
};
ArmOperation<FirewallPolicyRuleCollectionGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleCollectionGroupName, data);
FirewallPolicyRuleCollectionGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FirewallPolicyRuleCollectionGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
