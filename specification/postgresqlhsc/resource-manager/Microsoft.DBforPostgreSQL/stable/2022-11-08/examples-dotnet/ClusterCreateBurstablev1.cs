using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDBForPostgreSql;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterCreateBurstablev1.json
// this example is just showing the usage of "Clusters_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CosmosDBForPostgreSqlClusterResource
CosmosDBForPostgreSqlClusterCollection collection = resourceGroupResource.GetCosmosDBForPostgreSqlClusters();

// invoke the operation
string clusterName = "testcluster-burstablev1";
CosmosDBForPostgreSqlClusterData data = new CosmosDBForPostgreSqlClusterData(new AzureLocation("westus"))
{
    AdministratorLoginPassword = "password",
    PostgresqlVersion = "15",
    CitusVersion = "11.3",
    PreferredPrimaryZone = "1",
    IsShardsOnCoordinatorEnabled = true,
    IsHAEnabled = false,
    CoordinatorServerEdition = "BurstableMemoryOptimized",
    CoordinatorStorageQuotaInMb = 131072,
    CoordinatorVCores = 1,
    IsCoordinatorPublicIPAccessEnabled = true,
    NodeCount = 0,
    Tags =
    {
    ["owner"] = "JohnDoe",
    },
};
ArmOperation<CosmosDBForPostgreSqlClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
CosmosDBForPostgreSqlClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBForPostgreSqlClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
