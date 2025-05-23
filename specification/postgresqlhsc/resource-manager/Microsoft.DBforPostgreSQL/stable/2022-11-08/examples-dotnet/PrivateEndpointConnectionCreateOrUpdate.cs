using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDBForPostgreSql.Models;
using Azure.ResourceManager.CosmosDBForPostgreSql;

// Generated from example definition: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/PrivateEndpointConnectionCreateOrUpdate.json
// this example is just showing the usage of "PrivateEndpointConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBForPostgreSqlClusterResource created on azure
// for more information of creating CosmosDBForPostgreSqlClusterResource, please refer to the document of CosmosDBForPostgreSqlClusterResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string clusterName = "testcluster";
ResourceIdentifier cosmosDBForPostgreSqlClusterResourceId = CosmosDBForPostgreSqlClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
CosmosDBForPostgreSqlClusterResource cosmosDBForPostgreSqlCluster = client.GetCosmosDBForPostgreSqlClusterResource(cosmosDBForPostgreSqlClusterResourceId);

// get the collection of this CosmosDBForPostgreSqlPrivateEndpointConnectionResource
CosmosDBForPostgreSqlPrivateEndpointConnectionCollection collection = cosmosDBForPostgreSqlCluster.GetCosmosDBForPostgreSqlPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "private-endpoint-connection-name";
CosmosDBForPostgreSqlPrivateEndpointConnectionData data = new CosmosDBForPostgreSqlPrivateEndpointConnectionData
{
    ConnectionState = new CosmosDBForPostgreSqlPrivateLinkServiceConnectionState
    {
        Status = CosmosDBForPostgreSqlPrivateEndpointServiceConnectionStatus.Approved,
        Description = "Approved by johndoe@contoso.com",
    },
};
ArmOperation<CosmosDBForPostgreSqlPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
CosmosDBForPostgreSqlPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBForPostgreSqlPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
