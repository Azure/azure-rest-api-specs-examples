using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: 2024-12-30/CheckNameAvailability.json
// this example is just showing the usage of "CheckNameAvailability_CheckMySqlFlexibleServerNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation locationName = new AzureLocation("SouthEastAsia");
MySqlFlexibleServerNameAvailabilityContent content = new MySqlFlexibleServerNameAvailabilityContent("name1")
{
    ResourceType = new ResourceType("Microsoft.DBforMySQL/flexibleServers"),
};
MySqlFlexibleServerNameAvailabilityResult result = await subscriptionResource.CheckMySqlFlexibleServerNameAvailabilityAsync(locationName, content);

Console.WriteLine($"Succeeded: {result}");
