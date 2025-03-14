using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CustomerInsights.Models;
using Azure.ResourceManager.CustomerInsights;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesRegeneratePrimaryKey.json
// this example is just showing the usage of "AuthorizationPolicies_RegeneratePrimaryKey" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AuthorizationPolicyResourceFormatResource created on azure
// for more information of creating AuthorizationPolicyResourceFormatResource, please refer to the document of AuthorizationPolicyResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "azSdkTestHub";
string authorizationPolicyName = "testPolicy4222";
ResourceIdentifier authorizationPolicyResourceFormatResourceId = AuthorizationPolicyResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, authorizationPolicyName);
AuthorizationPolicyResourceFormatResource authorizationPolicyResourceFormat = client.GetAuthorizationPolicyResourceFormatResource(authorizationPolicyResourceFormatResourceId);

// invoke the operation
AuthorizationPolicy result = await authorizationPolicyResourceFormat.RegeneratePrimaryKeyAsync();

Console.WriteLine($"Succeeded: {result}");
