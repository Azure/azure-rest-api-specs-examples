using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDBForPostgreSql.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CosmosDBForPostgreSql;

// Generated from example definition: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/CheckNameAvailability.json
// this example is just showing the usage of "Clusters_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

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
CosmosDBForPostgreSqlClusterNameAvailabilityContent content = new CosmosDBForPostgreSqlClusterNameAvailabilityContent("name1");
CosmosDBForPostgreSqlClusterNameAvailabilityResult result = await subscriptionResource.CheckCosmosDBForPostgreSqlClusterNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
