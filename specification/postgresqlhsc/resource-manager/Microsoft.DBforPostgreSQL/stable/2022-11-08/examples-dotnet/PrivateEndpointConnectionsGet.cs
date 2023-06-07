using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDBForPostgreSql;
using Azure.ResourceManager.CosmosDBForPostgreSql.Models;

// Generated from example definition: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/PrivateEndpointConnectionsGet.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBForPostgreSqlPrivateEndpointConnectionResource created on azure
// for more information of creating CosmosDBForPostgreSqlPrivateEndpointConnectionResource, please refer to the document of CosmosDBForPostgreSqlPrivateEndpointConnectionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string clusterName = "testcluster";
string privateEndpointConnectionName = "private-endpoint-connection-name";
ResourceIdentifier cosmosDBForPostgreSqlPrivateEndpointConnectionResourceId = CosmosDBForPostgreSqlPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, privateEndpointConnectionName);
CosmosDBForPostgreSqlPrivateEndpointConnectionResource cosmosDBForPostgreSqlPrivateEndpointConnection = client.GetCosmosDBForPostgreSqlPrivateEndpointConnectionResource(cosmosDBForPostgreSqlPrivateEndpointConnectionResourceId);

// invoke the operation
CosmosDBForPostgreSqlPrivateEndpointConnectionResource result = await cosmosDBForPostgreSqlPrivateEndpointConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBForPostgreSqlPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
