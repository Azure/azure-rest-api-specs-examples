using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/DatabaseDelete.json
// this example is just showing the usage of "Databases_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerDatabaseResource created on azure
// for more information of creating PostgreSqlFlexibleServerDatabaseResource, please refer to the document of PostgreSqlFlexibleServerDatabaseResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "testserver";
string databaseName = "db1";
ResourceIdentifier postgreSqlFlexibleServerDatabaseResourceId = PostgreSqlFlexibleServerDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
PostgreSqlFlexibleServerDatabaseResource postgreSqlFlexibleServerDatabase = client.GetPostgreSqlFlexibleServerDatabaseResource(postgreSqlFlexibleServerDatabaseResourceId);

// invoke the operation
await postgreSqlFlexibleServerDatabase.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
