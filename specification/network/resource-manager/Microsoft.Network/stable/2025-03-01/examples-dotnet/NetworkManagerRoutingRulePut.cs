using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerRoutingRulePut.json
// this example is just showing the usage of "RoutingRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerRoutingRulesResource created on azure
// for more information of creating NetworkManagerRoutingRulesResource, please refer to the document of NetworkManagerRoutingRulesResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestRoutingConfig";
string ruleCollectionName = "testRuleCollection";
ResourceIdentifier networkManagerRoutingRulesResourceId = NetworkManagerRoutingRulesResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName, ruleCollectionName);
NetworkManagerRoutingRulesResource networkManagerRoutingRules = client.GetNetworkManagerRoutingRulesResource(networkManagerRoutingRulesResourceId);

// get the collection of this NetworkManagerRoutingRuleResource
NetworkManagerRoutingRuleCollection collection = networkManagerRoutingRules.GetNetworkManagerRoutingRules();

// invoke the operation
string ruleName = "SampleRoutingRule";
NetworkManagerRoutingRuleData data = new NetworkManagerRoutingRuleData
{
    Description = "This is Sample Routing Rule",
    Destination = new RoutingRuleRouteDestination(RoutingRuleDestinationType.AddressPrefix, "10.0.0.0/16"),
    NextHop = new RoutingRuleNextHop(RoutingRuleNextHopType.VirtualNetworkGateway),
};
ArmOperation<NetworkManagerRoutingRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleName, data);
NetworkManagerRoutingRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkManagerRoutingRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
