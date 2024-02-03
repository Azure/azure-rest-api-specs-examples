using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Purview;
using Azure.ResourceManager.Purview.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/purview/resource-manager/Microsoft.Purview/preview/2023-05-01-preview/examples/Accounts_CheckNameAvailability.json
// this example is just showing the usage of "Accounts_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
PurviewAccountNameAvailabilityContent content = new PurviewAccountNameAvailabilityContent()
{
    Name = "account1",
    ResourceType = "Microsoft.Purview/accounts",
};
PurviewAccountNameAvailabilityResult result = await subscriptionResource.CheckPurviewAccountNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
