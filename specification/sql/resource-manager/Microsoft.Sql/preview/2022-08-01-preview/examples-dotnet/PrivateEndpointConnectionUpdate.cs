using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/PrivateEndpointConnectionUpdate.json
// this example is just showing the usage of "PrivateEndpointConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerResource created on azure
// for more information of creating SqlServerResource, please refer to the document of SqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "test-svr";
ResourceIdentifier sqlServerResourceId = SqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
SqlServerResource sqlServer = client.GetSqlServerResource(sqlServerResourceId);

// get the collection of this SqlPrivateEndpointConnectionResource
SqlPrivateEndpointConnectionCollection collection = sqlServer.GetSqlPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "private-endpoint-connection-name";
SqlPrivateEndpointConnectionData data = new SqlPrivateEndpointConnectionData()
{
    ConnectionState = new SqlPrivateLinkServiceConnectionStateProperty(SqlPrivateLinkServiceConnectionStatus.Approved, "Approved by johndoe@contoso.com"),
};
ArmOperation<SqlPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
SqlPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
