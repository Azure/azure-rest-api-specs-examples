using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerRoutingRuleCollectionList.json
// this example is just showing the usage of "RoutingRuleCollections_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerRoutingConfigurationResource created on azure
// for more information of creating NetworkManagerRoutingConfigurationResource, please refer to the document of NetworkManagerRoutingConfigurationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestRoutingConfig";
ResourceIdentifier networkManagerRoutingConfigurationResourceId = NetworkManagerRoutingConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName);
NetworkManagerRoutingConfigurationResource networkManagerRoutingConfiguration = client.GetNetworkManagerRoutingConfigurationResource(networkManagerRoutingConfigurationResourceId);

// get the collection of this NetworkManagerRoutingRulesResource
NetworkManagerRoutingRulesCollection collection = networkManagerRoutingConfiguration.GetAllNetworkManagerRoutingRules();

// invoke the operation and iterate over the result
await foreach (NetworkManagerRoutingRulesResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkManagerRoutingRulesData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
