using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/PrivateEndpointConnectionUpdate.json
// this example is just showing the usage of "PrivateEndpointConnection_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServersPrivateEndpointConnectionResource created on azure
// for more information of creating PostgreSqlFlexibleServersPrivateEndpointConnectionResource, please refer to the document of PostgreSqlFlexibleServersPrivateEndpointConnectionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "Default";
string serverName = "test-svr";
string privateEndpointConnectionName = "private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e";
ResourceIdentifier postgreSqlFlexibleServersPrivateEndpointConnectionResourceId = PostgreSqlFlexibleServersPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, privateEndpointConnectionName);
PostgreSqlFlexibleServersPrivateEndpointConnectionResource postgreSqlFlexibleServersPrivateEndpointConnection = client.GetPostgreSqlFlexibleServersPrivateEndpointConnectionResource(postgreSqlFlexibleServersPrivateEndpointConnectionResourceId);

// invoke the operation
PostgreSqlFlexibleServersPrivateEndpointConnectionData data = new PostgreSqlFlexibleServersPrivateEndpointConnectionData
{
    ConnectionState = new PostgreSqlFlexibleServersPrivateLinkServiceConnectionState
    {
        Status = PostgreSqlFlexibleServersPrivateEndpointServiceConnectionStatus.Approved,
        Description = "Approved by johndoe@contoso.com",
    },
};
ArmOperation<PostgreSqlFlexibleServersPrivateEndpointConnectionResource> lro = await postgreSqlFlexibleServersPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, data);
PostgreSqlFlexibleServersPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlFlexibleServersPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
