using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleRegenerateKey.json
// this example is just showing the usage of "WCFRelays_RegenerateKeys" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this WcfRelayAuthorizationRuleResource created on azure
// for more information of creating WcfRelayAuthorizationRuleResource, please refer to the document of WcfRelayAuthorizationRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "resourcegroup";
string namespaceName = "example-RelayNamespace-01";
string relayName = "example-Relay-wcf-01";
string authorizationRuleName = "example-RelayAuthRules-01";
ResourceIdentifier wcfRelayAuthorizationRuleResourceId = WcfRelayAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, relayName, authorizationRuleName);
WcfRelayAuthorizationRuleResource wcfRelayAuthorizationRule = client.GetWcfRelayAuthorizationRuleResource(wcfRelayAuthorizationRuleResourceId);

// invoke the operation
RelayRegenerateAccessKeyContent content = new RelayRegenerateAccessKeyContent(RelayAccessKeyType.PrimaryKey);
RelayAccessKeys result = await wcfRelayAuthorizationRule.RegenerateKeysAsync(content);

Console.WriteLine($"Succeeded: {result}");
