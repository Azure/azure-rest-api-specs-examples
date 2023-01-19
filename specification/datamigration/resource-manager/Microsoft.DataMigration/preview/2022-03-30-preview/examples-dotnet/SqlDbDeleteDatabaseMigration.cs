using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataMigration;
using Azure.ResourceManager.DataMigration.Models;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/SqlDbDeleteDatabaseMigration.json
// this example is just showing the usage of "DatabaseMigrationsSqlDb_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseMigrationSqlDBResource created on azure
// for more information of creating DatabaseMigrationSqlDBResource, please refer to the document of DatabaseMigrationSqlDBResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlDBInstanceName = "sqldbinstance";
string targetDBName = "db1";
ResourceIdentifier databaseMigrationSqlDBResourceId = DatabaseMigrationSqlDBResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlDBInstanceName, targetDBName);
DatabaseMigrationSqlDBResource databaseMigrationSqlDB = client.GetDatabaseMigrationSqlDBResource(databaseMigrationSqlDBResourceId);

// invoke the operation
await databaseMigrationSqlDB.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
