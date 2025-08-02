using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/ColumnsListByDatabaseMax.json
// this example is just showing the usage of "DatabaseColumns_ListByDatabase" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string serverName = "serverName";
string databaseName = "myDatabase";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// invoke the operation and iterate over the result
IEnumerable<string> schema = new string[] { "dbo" };
IEnumerable<string> table = new string[] { "customer", "address" };
IEnumerable<string> column = new string[] { "username" };
IEnumerable<string> orderBy = new string[] { "schema asc", "table", "column desc" };
await foreach (SqlDatabaseColumnResource item in sqlDatabase.GetDatabaseColumnsAsync(schema: schema, table: table, column: column, orderBy: orderBy))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DatabaseColumnData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
