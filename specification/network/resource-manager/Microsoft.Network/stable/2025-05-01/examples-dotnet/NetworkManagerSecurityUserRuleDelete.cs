using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerSecurityUserRuleDelete.json
// this example is just showing the usage of "SecurityUserRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerSecurityUserRuleResource created on azure
// for more information of creating NetworkManagerSecurityUserRuleResource, please refer to the document of NetworkManagerSecurityUserRuleResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestSecurityConfig";
string ruleCollectionName = "testRuleCollection";
string ruleName = "SampleUserRule";
ResourceIdentifier networkManagerSecurityUserRuleResourceId = NetworkManagerSecurityUserRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName);
NetworkManagerSecurityUserRuleResource networkManagerSecurityUserRule = client.GetNetworkManagerSecurityUserRuleResource(networkManagerSecurityUserRuleResourceId);

// invoke the operation
bool? force = false;
await networkManagerSecurityUserRule.DeleteAsync(WaitUntil.Completed, force: force);

Console.WriteLine("Succeeded");
