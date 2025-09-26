using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.DataMigration;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/SqlDbGetDatabaseMigration.json
// this example is just showing the usage of "DatabaseMigrationsSqlDb_Get" operation, for the dependent resources, they will have to be created separately.

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
DatabaseMigrationSqlDBResource result = await databaseMigrationSqlDB.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatabaseMigrationSqlDBData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
