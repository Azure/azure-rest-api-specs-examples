using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NspAccessRuleGet.json
// this example is just showing the usage of "NetworkSecurityPerimeterAccessRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterProfileResource created on azure
// for more information of creating NetworkSecurityPerimeterProfileResource, please refer to the document of NetworkSecurityPerimeterProfileResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string networkSecurityPerimeterName = "nsp1";
string profileName = "profile1";
ResourceIdentifier networkSecurityPerimeterProfileResourceId = NetworkSecurityPerimeterProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityPerimeterName, profileName);
NetworkSecurityPerimeterProfileResource networkSecurityPerimeterProfile = client.GetNetworkSecurityPerimeterProfileResource(networkSecurityPerimeterProfileResourceId);

// get the collection of this NetworkSecurityPerimeterAccessRuleResource
NetworkSecurityPerimeterAccessRuleCollection collection = networkSecurityPerimeterProfile.GetNetworkSecurityPerimeterAccessRules();

// invoke the operation
string accessRuleName = "accessRule1";
NullableResponse<NetworkSecurityPerimeterAccessRuleResource> response = await collection.GetIfExistsAsync(accessRuleName);
NetworkSecurityPerimeterAccessRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkSecurityPerimeterAccessRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
