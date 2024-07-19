using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRestorePointsDelete.json
// this example is just showing the usage of "RestorePoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerDatabaseRestorePointResource created on azure
// for more information of creating SqlServerDatabaseRestorePointResource, please refer to the document of SqlServerDatabaseRestorePointResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string serverName = "testserver";
string databaseName = "testDatabase";
string restorePointName = "131546477590000000";
ResourceIdentifier sqlServerDatabaseRestorePointResourceId = SqlServerDatabaseRestorePointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, restorePointName);
SqlServerDatabaseRestorePointResource sqlServerDatabaseRestorePoint = client.GetSqlServerDatabaseRestorePointResource(sqlServerDatabaseRestorePointResourceId);

// invoke the operation
await sqlServerDatabaseRestorePoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
