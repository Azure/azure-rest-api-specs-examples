using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDBForPostgreSql;

// Generated from example definition: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ConfigurationListByServer.json
// this example is just showing the usage of "Configurations_ListByServer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBForPostgreSqlClusterServerResource created on azure
// for more information of creating CosmosDBForPostgreSqlClusterServerResource, please refer to the document of CosmosDBForPostgreSqlClusterServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestResourceGroup";
string clusterName = "testcluster";
string serverName = "testserver";
ResourceIdentifier cosmosDBForPostgreSqlClusterServerResourceId = CosmosDBForPostgreSqlClusterServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, serverName);
CosmosDBForPostgreSqlClusterServerResource cosmosDBForPostgreSqlClusterServer = client.GetCosmosDBForPostgreSqlClusterServerResource(cosmosDBForPostgreSqlClusterServerResourceId);

// invoke the operation and iterate over the result
await foreach (CosmosDBForPostgreSqlServerConfigurationData item in cosmosDBForPostgreSqlClusterServer.GetConfigurationsAsync())
{
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {item.Id}");
}

Console.WriteLine("Succeeded");
