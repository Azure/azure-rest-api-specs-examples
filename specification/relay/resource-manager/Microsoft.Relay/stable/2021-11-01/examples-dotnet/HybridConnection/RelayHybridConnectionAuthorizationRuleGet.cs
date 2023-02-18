using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay;
using Azure.ResourceManager.Relay.Models;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleGet.json
// this example is just showing the usage of "HybridConnections_GetAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RelayHybridConnectionResource created on azure
// for more information of creating RelayHybridConnectionResource, please refer to the document of RelayHybridConnectionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "resourcegroup";
string namespaceName = "example-RelayNamespace-01";
string hybridConnectionName = "example-Relay-Hybrid-01";
ResourceIdentifier relayHybridConnectionResourceId = RelayHybridConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, hybridConnectionName);
RelayHybridConnectionResource relayHybridConnection = client.GetRelayHybridConnectionResource(relayHybridConnectionResourceId);

// get the collection of this RelayHybridConnectionAuthorizationRuleResource
RelayHybridConnectionAuthorizationRuleCollection collection = relayHybridConnection.GetRelayHybridConnectionAuthorizationRules();

// invoke the operation
string authorizationRuleName = "example-RelayAuthRules-01";
bool result = await collection.ExistsAsync(authorizationRuleName);

Console.WriteLine($"Succeeded: {result}");
