using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/NetworkManagerDefaultAdminRuleGet.json
// this example is just showing the usage of "AdminRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AdminRuleGroupResource created on azure
// for more information of creating AdminRuleGroupResource, please refer to the document of AdminRuleGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestSecurityConfig";
string ruleCollectionName = "testRuleCollection";
ResourceIdentifier adminRuleGroupResourceId = AdminRuleGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName, ruleCollectionName);
AdminRuleGroupResource adminRuleGroup = client.GetAdminRuleGroupResource(adminRuleGroupResourceId);

// get the collection of this BaseAdminRuleResource
BaseAdminRuleCollection collection = adminRuleGroup.GetBaseAdminRules();

// invoke the operation
string ruleName = "SampleDefaultAdminRule";
NullableResponse<BaseAdminRuleResource> response = await collection.GetIfExistsAsync(ruleName);
BaseAdminRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BaseAdminRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
