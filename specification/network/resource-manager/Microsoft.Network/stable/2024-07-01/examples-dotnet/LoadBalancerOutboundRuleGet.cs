using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/LoadBalancerOutboundRuleGet.json
// this example is just showing the usage of "LoadBalancerOutboundRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OutboundRuleResource created on azure
// for more information of creating OutboundRuleResource, please refer to the document of OutboundRuleResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string loadBalancerName = "lb1";
string outboundRuleName = "rule1";
ResourceIdentifier outboundRuleResourceId = OutboundRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadBalancerName, outboundRuleName);
OutboundRuleResource outboundRule = client.GetOutboundRuleResource(outboundRuleResourceId);

// invoke the operation
OutboundRuleResource result = await outboundRule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OutboundRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
