using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CustomerInsights.Models;
using Azure.ResourceManager.CustomerInsights;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesGet.json
// this example is just showing the usage of "AuthorizationPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubResource created on azure
// for more information of creating HubResource, please refer to the document of HubResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "azSdkTestHub";
ResourceIdentifier hubResourceId = HubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName);
HubResource hub = client.GetHubResource(hubResourceId);

// get the collection of this AuthorizationPolicyResourceFormatResource
AuthorizationPolicyResourceFormatCollection collection = hub.GetAuthorizationPolicyResourceFormats();

// invoke the operation
string authorizationPolicyName = "testPolicy4222";
NullableResponse<AuthorizationPolicyResourceFormatResource> response = await collection.GetIfExistsAsync(authorizationPolicyName);
AuthorizationPolicyResourceFormatResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AuthorizationPolicyResourceFormatData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
