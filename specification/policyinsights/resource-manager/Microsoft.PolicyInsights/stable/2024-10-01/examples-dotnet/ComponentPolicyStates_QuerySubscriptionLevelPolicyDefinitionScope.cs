using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QuerySubscriptionLevelPolicyDefinitionScope.json
// this example is just showing the usage of "ComponentPolicyStates_ListQueryResultsForPolicyDefinition" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "fffedd8f-ffff-fffd-fffd-fffed2f84852";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
string policyDefinitionName = "24813039-7534-408a-9842-eb99f45721b1";
ComponentPolicyStatesResource componentPolicyStatesResource = ComponentPolicyStatesResource.Latest;
SubscriptionResourceGetQueryResultsForPolicyDefinitionComponentPolicyStatesOptions options = new SubscriptionResourceGetQueryResultsForPolicyDefinitionComponentPolicyStatesOptions(policyDefinitionName, componentPolicyStatesResource);
await foreach (ComponentPolicyState item in subscriptionResource.GetQueryResultsForPolicyDefinitionComponentPolicyStatesAsync(options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
