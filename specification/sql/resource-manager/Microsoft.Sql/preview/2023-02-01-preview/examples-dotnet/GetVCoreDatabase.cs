using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/GetVCoreDatabase.json
// this example is just showing the usage of "Databases_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerResource created on azure
// for more information of creating SqlServerResource, please refer to the document of SqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string serverName = "testsvr";
ResourceIdentifier sqlServerResourceId = SqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
SqlServerResource sqlServer = client.GetSqlServerResource(sqlServerResourceId);

// get the collection of this SqlDatabaseResource
SqlDatabaseCollection collection = sqlServer.GetSqlDatabases();

// invoke the operation
string databaseName = "testdb";
NullableResponse<SqlDatabaseResource> response = await collection.GetIfExistsAsync(databaseName);
SqlDatabaseResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SqlDatabaseData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
