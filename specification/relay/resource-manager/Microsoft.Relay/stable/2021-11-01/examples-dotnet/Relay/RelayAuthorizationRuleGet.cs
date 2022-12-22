using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleGet.json
// this example is just showing the usage of "WCFRelays_GetAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this WcfRelayResource created on azure
// for more information of creating WcfRelayResource, please refer to the document of WcfRelayResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "resourcegroup";
string namespaceName = "example-RelayNamespace-01";
string relayName = "example-Relay-wcf-01";
ResourceIdentifier wcfRelayResourceId = WcfRelayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, relayName);
WcfRelayResource wcfRelay = client.GetWcfRelayResource(wcfRelayResourceId);

// get the collection of this WcfRelayAuthorizationRuleResource
WcfRelayAuthorizationRuleCollection collection = wcfRelay.GetWcfRelayAuthorizationRules();

// invoke the operation
string authorizationRuleName = "example-RelayAuthRules-01";
bool result = await collection.ExistsAsync(authorizationRuleName);

Console.WriteLine($"Succeeded: {result}");
