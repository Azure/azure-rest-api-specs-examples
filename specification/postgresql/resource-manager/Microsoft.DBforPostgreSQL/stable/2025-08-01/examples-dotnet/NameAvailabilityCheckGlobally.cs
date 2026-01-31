using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/NameAvailabilityCheckGlobally.json
// this example is just showing the usage of "NameAvailability_CheckGlobally" operation, for the dependent resources, they will have to be created separately.

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
PostgreSqlFlexibleServerNameAvailabilityContent content = new PostgreSqlFlexibleServerNameAvailabilityContent("exampleserver")
{
    ResourceType = new ResourceType("Microsoft.DBforPostgreSQL/flexibleServers"),
};
PostgreSqlFlexibleServerNameAvailabilityResult result = await subscriptionResource.CheckPostgreSqlFlexibleServerNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
