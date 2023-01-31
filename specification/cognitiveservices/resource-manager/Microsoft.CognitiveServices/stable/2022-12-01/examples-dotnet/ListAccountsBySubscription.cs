using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/ListAccountsBySubscription.json
// this example is just showing the usage of "Accounts_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
await foreach (CognitiveServicesAccountResource item in subscriptionResource.GetCognitiveServicesAccountsAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CognitiveServicesAccountData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
