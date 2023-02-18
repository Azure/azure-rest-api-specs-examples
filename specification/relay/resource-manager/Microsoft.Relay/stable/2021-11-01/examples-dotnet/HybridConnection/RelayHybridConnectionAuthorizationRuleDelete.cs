using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay;
using Azure.ResourceManager.Relay.Models;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleDelete.json
// this example is just showing the usage of "HybridConnections_DeleteAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RelayHybridConnectionAuthorizationRuleResource created on azure
// for more information of creating RelayHybridConnectionAuthorizationRuleResource, please refer to the document of RelayHybridConnectionAuthorizationRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "resourcegroup";
string namespaceName = "example-RelayNamespace-01";
string hybridConnectionName = "example-Relay-Hybrid-01";
string authorizationRuleName = "example-RelayAuthRules-01";
ResourceIdentifier relayHybridConnectionAuthorizationRuleResourceId = RelayHybridConnectionAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, hybridConnectionName, authorizationRuleName);
RelayHybridConnectionAuthorizationRuleResource relayHybridConnectionAuthorizationRule = client.GetRelayHybridConnectionAuthorizationRuleResource(relayHybridConnectionAuthorizationRuleResourceId);

// invoke the operation
await relayHybridConnectionAuthorizationRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
