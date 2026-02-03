using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerRoutingRuleDelete.json
// this example is just showing the usage of "RoutingRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerRoutingRuleResource created on azure
// for more information of creating NetworkManagerRoutingRuleResource, please refer to the document of NetworkManagerRoutingRuleResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestRoutingConfig";
string ruleCollectionName = "testRuleCollection";
string ruleName = "sampleRule";
ResourceIdentifier networkManagerRoutingRuleResourceId = NetworkManagerRoutingRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName);
NetworkManagerRoutingRuleResource networkManagerRoutingRule = client.GetNetworkManagerRoutingRuleResource(networkManagerRoutingRuleResourceId);

// invoke the operation
bool? force = false;
await networkManagerRoutingRule.DeleteAsync(WaitUntil.Completed, force: force);

Console.WriteLine("Succeeded");
