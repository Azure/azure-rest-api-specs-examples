using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/mongo-cluster/CosmosDBMongoClusterFirewallRuleGet.json
// this example is just showing the usage of "MongoClusters_GetFirewallRule" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoClusterResource created on azure
// for more information of creating MongoClusterResource, please refer to the document of MongoClusterResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string mongoClusterName = "myMongoCluster";
ResourceIdentifier mongoClusterResourceId = MongoClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mongoClusterName);
MongoClusterResource mongoCluster = client.GetMongoClusterResource(mongoClusterResourceId);

// get the collection of this CosmosDBFirewallRuleResource
CosmosDBFirewallRuleCollection collection = mongoCluster.GetCosmosDBFirewallRules();

// invoke the operation
string firewallRuleName = "rule1";
NullableResponse<CosmosDBFirewallRuleResource> response = await collection.GetIfExistsAsync(firewallRuleName);
CosmosDBFirewallRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CosmosDBFirewallRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
