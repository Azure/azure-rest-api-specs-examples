using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/InboundSecurityRulePut.json
// this example is just showing the usage of "InboundSecurityRule_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InboundSecurityRuleResource created on azure
// for more information of creating InboundSecurityRuleResource, please refer to the document of InboundSecurityRuleResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkVirtualApplianceName = "nva";
string ruleCollectionName = "rule1";
ResourceIdentifier inboundSecurityRuleResourceId = InboundSecurityRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkVirtualApplianceName, ruleCollectionName);
InboundSecurityRuleResource inboundSecurityRule = client.GetInboundSecurityRuleResource(inboundSecurityRuleResourceId);

// invoke the operation
InboundSecurityRuleData data = new InboundSecurityRuleData
{
    RuleType = InboundSecurityRuleType.Permanent,
    Rules = {new InboundSecurityRules
    {
    Name = "inboundRule1",
    Protocol = InboundSecurityRulesProtocol.Tcp,
    SourceAddressPrefix = "50.20.121.5/32",
    DestinationPortRange = 22,
    DestinationPortRanges = {"80-100"},
    AppliesOn = {"slbip1"},
    }},
};
ArmOperation<InboundSecurityRuleResource> lro = await inboundSecurityRule.UpdateAsync(WaitUntil.Completed, data);
InboundSecurityRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
InboundSecurityRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
