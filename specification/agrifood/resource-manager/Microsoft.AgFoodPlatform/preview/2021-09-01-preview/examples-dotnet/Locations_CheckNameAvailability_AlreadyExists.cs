using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AgFoodPlatform.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AgFoodPlatform;

// Generated from example definition: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Locations_CheckNameAvailability_AlreadyExists.json
// this example is just showing the usage of "Locations_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
CheckNameAvailabilityContent content = new CheckNameAvailabilityContent
{
    Name = "existingaccountname",
    ResourceType = "Microsoft.AgFoodPlatform/farmBeats",
};
CheckNameAvailabilityResponse result = await subscriptionResource.CheckNameAvailabilityLocationAsync(content);

Console.WriteLine($"Succeeded: {result}");
