using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PostgreSql.FlexibleServers;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/Migrations_Delete.json
// this example is just showing the usage of "Migrations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlMigrationResource created on azure
// for more information of creating PostgreSqlMigrationResource, please refer to the document of PostgreSqlMigrationResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string targetDbServerName = "testtarget";
string migrationName = "testmigration";
ResourceIdentifier postgreSqlMigrationResourceId = PostgreSqlMigrationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, targetDbServerName, migrationName);
PostgreSqlMigrationResource postgreSqlMigration = client.GetPostgreSqlMigrationResource(postgreSqlMigrationResourceId);

// invoke the operation
await postgreSqlMigration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
