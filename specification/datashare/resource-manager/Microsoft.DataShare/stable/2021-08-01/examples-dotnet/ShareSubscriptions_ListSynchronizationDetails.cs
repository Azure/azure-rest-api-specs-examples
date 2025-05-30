using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataShare.Models;
using Azure.ResourceManager.DataShare;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/ShareSubscriptions_ListSynchronizationDetails.json
// this example is just showing the usage of "ShareSubscriptions_ListSynchronizationDetails" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ShareSubscriptionResource created on azure
// for more information of creating ShareSubscriptionResource, please refer to the document of ShareSubscriptionResource
string subscriptionId = "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareSubscriptionName = "ShareSub1";
ResourceIdentifier shareSubscriptionResourceId = ShareSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareSubscriptionName);
ShareSubscriptionResource shareSubscription = client.GetShareSubscriptionResource(shareSubscriptionResourceId);

// invoke the operation and iterate over the result
ShareSubscriptionSynchronization shareSubscriptionSynchronization = new ShareSubscriptionSynchronization(Guid.Parse("7d0536a6-3fa5-43de-b152-3d07c4f6b2bb"));
await foreach (SynchronizationDetails item in shareSubscription.GetSynchronizationDetailsAsync(shareSubscriptionSynchronization))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
