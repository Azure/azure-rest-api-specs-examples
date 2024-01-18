using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/InboundNatRuleCreate.json
// this example is just showing the usage of "InboundNatRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InboundNatRuleResource created on azure
// for more information of creating InboundNatRuleResource, please refer to the document of InboundNatRuleResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string loadBalancerName = "lb1";
string inboundNatRuleName = "natRule1.1";
ResourceIdentifier inboundNatRuleResourceId = InboundNatRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadBalancerName, inboundNatRuleName);
InboundNatRuleResource inboundNatRule = client.GetInboundNatRuleResource(inboundNatRuleResourceId);

// invoke the operation
InboundNatRuleData data = new InboundNatRuleData()
{
    FrontendIPConfigurationId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
    Protocol = LoadBalancingTransportProtocol.Tcp,
    FrontendPort = 3390,
    BackendPort = 3389,
    IdleTimeoutInMinutes = 4,
    EnableFloatingIP = false,
    EnableTcpReset = false,
};
ArmOperation<InboundNatRuleResource> lro = await inboundNatRule.UpdateAsync(WaitUntil.Completed, data);
InboundNatRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
InboundNatRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
