using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NewRelicObservability.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NewRelicObservability;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-03-01/examples/Accounts_List_MaximumSet_Gen.json
// this example is just showing the usage of "Accounts_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
string userEmail = "ruxvg@xqkmdhrnoo.hlmbpm";
AzureLocation location = new AzureLocation("egh");
await foreach (NewRelicAccountResourceData item in subscriptionResource.GetNewRelicAccountsAsync(userEmail, location))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
