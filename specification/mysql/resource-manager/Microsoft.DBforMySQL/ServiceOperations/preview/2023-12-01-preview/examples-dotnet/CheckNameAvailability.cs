using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql.FlexibleServers;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2023-12-01-preview/examples/CheckNameAvailability.json
// this example is just showing the usage of "CheckNameAvailabilityWithoutLocation_Execute" operation, for the dependent resources, they will have to be created separately.

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
MySqlFlexibleServerNameAvailabilityContent content = new MySqlFlexibleServerNameAvailabilityContent("name1")
{
    ResourceType = new ResourceType("Microsoft.DBforMySQL/flexibleServers"),
};
MySqlFlexibleServerNameAvailabilityResult result = await subscriptionResource.CheckMySqlFlexibleServerNameAvailabilityWithoutLocationAsync(content);

Console.WriteLine($"Succeeded: {result}");
