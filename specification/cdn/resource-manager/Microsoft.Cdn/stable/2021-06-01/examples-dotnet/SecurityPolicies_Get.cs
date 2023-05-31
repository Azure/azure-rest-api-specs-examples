using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/SecurityPolicies_Get.json
// this example is just showing the usage of "FrontDoorSecurityPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProfileResource created on azure
// for more information of creating ProfileResource, please refer to the document of ProfileResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
ResourceIdentifier profileResourceId = ProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
ProfileResource profile = client.GetProfileResource(profileResourceId);

// get the collection of this FrontDoorSecurityPolicyResource
FrontDoorSecurityPolicyCollection collection = profile.GetFrontDoorSecurityPolicies();

// invoke the operation
string securityPolicyName = "securityPolicy1";
bool result = await collection.ExistsAsync(securityPolicyName);

Console.WriteLine($"Succeeded: {result}");
