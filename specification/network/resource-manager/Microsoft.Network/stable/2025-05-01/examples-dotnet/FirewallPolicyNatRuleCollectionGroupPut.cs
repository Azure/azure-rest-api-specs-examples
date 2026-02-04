using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/FirewallPolicyNatRuleCollectionGroupPut.json
// this example is just showing the usage of "FirewallPolicyRuleCollectionGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirewallPolicyRuleCollectionGroupResource created on azure
// for more information of creating FirewallPolicyRuleCollectionGroupResource, please refer to the document of FirewallPolicyRuleCollectionGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string firewallPolicyName = "firewallPolicy";
string ruleCollectionGroupName = "ruleCollectionGroup1";
ResourceIdentifier firewallPolicyRuleCollectionGroupResourceId = FirewallPolicyRuleCollectionGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallPolicyName, ruleCollectionGroupName);
FirewallPolicyRuleCollectionGroupResource firewallPolicyRuleCollectionGroup = client.GetFirewallPolicyRuleCollectionGroupResource(firewallPolicyRuleCollectionGroupResourceId);

// invoke the operation
FirewallPolicyRuleCollectionGroupData data = new FirewallPolicyRuleCollectionGroupData
{
    Priority = 100,
    RuleCollections = {new FirewallPolicyNatRuleCollectionInfo
    {
    ActionType = FirewallPolicyNatRuleCollectionActionType.Dnat,
    Rules = {new NatRule
    {
    IPProtocols = {FirewallPolicyRuleNetworkProtocol.Tcp, FirewallPolicyRuleNetworkProtocol.Udp},
    SourceAddresses = {"2.2.2.2"},
    DestinationAddresses = {"152.23.32.23"},
    DestinationPorts = {"8080"},
    TranslatedPort = "8080",
    SourceIPGroups = {},
    TranslatedFqdn = "internalhttp.server.net",
    Name = "nat-rule1",
    }},
    Name = "Example-Nat-Rule-Collection",
    Priority = 100,
    }},
};
ArmOperation<FirewallPolicyRuleCollectionGroupResource> lro = await firewallPolicyRuleCollectionGroup.UpdateAsync(WaitUntil.Completed, data);
FirewallPolicyRuleCollectionGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FirewallPolicyRuleCollectionGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
