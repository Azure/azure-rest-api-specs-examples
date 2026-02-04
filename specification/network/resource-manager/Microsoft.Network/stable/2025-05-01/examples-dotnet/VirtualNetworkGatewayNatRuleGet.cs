using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkGatewayNatRuleGet.json
// this example is just showing the usage of "VirtualNetworkGatewayNatRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkGatewayNatRuleResource created on azure
// for more information of creating VirtualNetworkGatewayNatRuleResource, please refer to the document of VirtualNetworkGatewayNatRuleResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualNetworkGatewayName = "gateway1";
string natRuleName = "natRule1";
ResourceIdentifier virtualNetworkGatewayNatRuleResourceId = VirtualNetworkGatewayNatRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkGatewayName, natRuleName);
VirtualNetworkGatewayNatRuleResource virtualNetworkGatewayNatRule = client.GetVirtualNetworkGatewayNatRuleResource(virtualNetworkGatewayNatRuleResourceId);

// invoke the operation
VirtualNetworkGatewayNatRuleResource result = await virtualNetworkGatewayNatRule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualNetworkGatewayNatRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
