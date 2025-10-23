using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QuerySubscriptionScope.json
// this example is just showing the usage of "ComponentPolicyStates_ListQueryResultsForSubscription" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "fff10b27-fff3-fff5-fff8-fffbe01e86a5";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
ComponentPolicyStatesResource componentPolicyStatesResource = ComponentPolicyStatesResource.Latest;
SubscriptionResourceGetQueryResultsForSubscriptionComponentPolicyStatesOptions options = new SubscriptionResourceGetQueryResultsForSubscriptionComponentPolicyStatesOptions(componentPolicyStatesResource);
await foreach (ComponentPolicyState item in subscriptionResource.GetQueryResultsForSubscriptionComponentPolicyStatesAsync(options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
