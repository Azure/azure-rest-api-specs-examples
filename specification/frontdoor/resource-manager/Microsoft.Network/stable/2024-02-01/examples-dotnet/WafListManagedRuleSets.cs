using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2024-02-01/examples/WafListManagedRuleSets.json
// this example is just showing the usage of "ManagedRuleSets_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subid";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
await foreach (ManagedRuleSetDefinition item in subscriptionResource.GetManagedRuleSetsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
