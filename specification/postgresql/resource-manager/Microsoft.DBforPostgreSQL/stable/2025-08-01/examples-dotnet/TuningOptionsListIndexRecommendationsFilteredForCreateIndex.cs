using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/TuningOptionsListIndexRecommendationsFilteredForCreateIndex.json
// this example is just showing the usage of "TuningOptions_ListRecommendations" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerTuningOptionResource created on azure
// for more information of creating PostgreSqlFlexibleServerTuningOptionResource, please refer to the document of PostgreSqlFlexibleServerTuningOptionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "exampleresourcegroup";
string serverName = "exampleserver";
PostgreSqlFlexibleServerTuningOptionType tuningOption = PostgreSqlFlexibleServerTuningOptionType.Index;
ResourceIdentifier postgreSqlFlexibleServerTuningOptionResourceId = PostgreSqlFlexibleServerTuningOptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, tuningOption);
PostgreSqlFlexibleServerTuningOptionResource postgreSqlFlexibleServerTuningOption = client.GetPostgreSqlFlexibleServerTuningOptionResource(postgreSqlFlexibleServerTuningOptionResourceId);

// invoke the operation and iterate over the result
PostgreSqlFlexibleServerRecommendationFilterType? recommendationType = PostgreSqlFlexibleServerRecommendationFilterType.CreateIndex;
await foreach (ObjectRecommendation item in postgreSqlFlexibleServerTuningOption.GetRecommendationsAsync(recommendationType: recommendationType))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
