using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/LoadBalancerHealth.json
// this example is just showing the usage of "LoadBalancerLoadBalancingRules_Health" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LoadBalancingRuleResource created on azure
// for more information of creating LoadBalancingRuleResource, please refer to the document of LoadBalancingRuleResource
string subscriptionId = "subid";
string groupName = "rg1";
string loadBalancerName = "lb1";
string loadBalancingRuleName = "rulelb";
ResourceIdentifier loadBalancingRuleResourceId = LoadBalancingRuleResource.CreateResourceIdentifier(subscriptionId, groupName, loadBalancerName, loadBalancingRuleName);
LoadBalancingRuleResource loadBalancingRule = client.GetLoadBalancingRuleResource(loadBalancingRuleResourceId);

// invoke the operation
ArmOperation<LoadBalancerHealthPerRule> lro = await loadBalancingRule.HealthAsync(WaitUntil.Completed);
LoadBalancerHealthPerRule result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
