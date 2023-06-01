using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/PrivateEndpointConnectionDelete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlPrivateEndpointConnectionResource created on azure
// for more information of creating SqlPrivateEndpointConnectionResource, please refer to the document of SqlPrivateEndpointConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "test-svr";
string privateEndpointConnectionName = "private-endpoint-connection-name";
ResourceIdentifier sqlPrivateEndpointConnectionResourceId = SqlPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, privateEndpointConnectionName);
SqlPrivateEndpointConnectionResource sqlPrivateEndpointConnection = client.GetSqlPrivateEndpointConnectionResource(sqlPrivateEndpointConnectionResourceId);

// invoke the operation
await sqlPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
