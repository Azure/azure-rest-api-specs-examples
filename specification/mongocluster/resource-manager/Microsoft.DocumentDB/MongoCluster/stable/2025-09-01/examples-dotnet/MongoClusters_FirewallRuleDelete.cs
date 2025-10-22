using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MongoCluster.Models;
using Azure.ResourceManager.MongoCluster;

// Generated from example definition: 2025-09-01/MongoClusters_FirewallRuleDelete.json
// this example is just showing the usage of "FirewallRule_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoClusterFirewallRuleResource created on azure
// for more information of creating MongoClusterFirewallRuleResource, please refer to the document of MongoClusterFirewallRuleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string mongoClusterName = "myMongoCluster";
string firewallRuleName = "rule1";
ResourceIdentifier mongoClusterFirewallRuleResourceId = MongoClusterFirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mongoClusterName, firewallRuleName);
MongoClusterFirewallRuleResource mongoClusterFirewallRule = client.GetMongoClusterFirewallRuleResource(mongoClusterFirewallRuleResourceId);

// invoke the operation
await mongoClusterFirewallRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
