using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.Models;
using Azure.ResourceManager.PostgreSql;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateEndpointConnectionUpdateTags.json
// this example is just showing the usage of "PrivateEndpointConnections_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlPrivateEndpointConnectionResource created on azure
// for more information of creating PostgreSqlPrivateEndpointConnectionResource, please refer to the document of PostgreSqlPrivateEndpointConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "test-svr";
string privateEndpointConnectionName = "private-endpoint-connection-name";
ResourceIdentifier postgreSqlPrivateEndpointConnectionResourceId = PostgreSqlPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, privateEndpointConnectionName);
PostgreSqlPrivateEndpointConnectionResource postgreSqlPrivateEndpointConnection = client.GetPostgreSqlPrivateEndpointConnectionResource(postgreSqlPrivateEndpointConnectionResourceId);

// invoke the operation
PostgreSqlPrivateEndpointConnectionPatch patch = new PostgreSqlPrivateEndpointConnectionPatch
{
    Tags =
    {
    ["key1"] = "val1",
    ["key2"] = "val2",
    ["key3"] = "val3"
    },
};
ArmOperation<PostgreSqlPrivateEndpointConnectionResource> lro = await postgreSqlPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, patch);
PostgreSqlPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
