using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/PrivateEndpointConnectionsDelete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServersPrivateEndpointConnectionResource created on azure
// for more information of creating PostgreSqlFlexibleServersPrivateEndpointConnectionResource, please refer to the document of PostgreSqlFlexibleServersPrivateEndpointConnectionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "exampleresourcegroup";
string serverName = "exampleserver";
string privateEndpointConnectionName = "private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e";
ResourceIdentifier postgreSqlFlexibleServersPrivateEndpointConnectionResourceId = PostgreSqlFlexibleServersPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, privateEndpointConnectionName);
PostgreSqlFlexibleServersPrivateEndpointConnectionResource postgreSqlFlexibleServersPrivateEndpointConnection = client.GetPostgreSqlFlexibleServersPrivateEndpointConnectionResource(postgreSqlFlexibleServersPrivateEndpointConnectionResourceId);

// invoke the operation
await postgreSqlFlexibleServersPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
