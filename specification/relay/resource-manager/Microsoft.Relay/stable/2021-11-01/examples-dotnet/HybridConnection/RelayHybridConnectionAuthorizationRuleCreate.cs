using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleCreate.json
// this example is just showing the usage of "HybridConnections_CreateOrUpdateAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

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
RelayAuthorizationRuleData data = new RelayAuthorizationRuleData
{
    Rights = { RelayAccessRight.Listen, RelayAccessRight.Send },
};
ArmOperation<RelayHybridConnectionAuthorizationRuleResource> lro = await relayHybridConnectionAuthorizationRule.UpdateAsync(WaitUntil.Completed, data);
RelayHybridConnectionAuthorizationRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RelayAuthorizationRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
