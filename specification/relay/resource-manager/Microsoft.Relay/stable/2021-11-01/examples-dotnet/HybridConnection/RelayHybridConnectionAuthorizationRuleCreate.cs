using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleCreate.json
// this example is just showing the usage of "HybridConnections_CreateOrUpdateAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

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
RelayAuthorizationRuleData data = new RelayAuthorizationRuleData()
{
    Rights =
    {
    RelayAccessRight.Listen,RelayAccessRight.Send
    },
};
ArmOperation<RelayHybridConnectionAuthorizationRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, authorizationRuleName, data);
RelayHybridConnectionAuthorizationRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RelayAuthorizationRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
