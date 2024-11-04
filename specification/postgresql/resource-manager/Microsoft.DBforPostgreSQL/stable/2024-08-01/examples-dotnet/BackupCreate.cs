using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/BackupCreate.json
// this example is just showing the usage of "Backups_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerBackupResource created on azure
// for more information of creating PostgreSqlFlexibleServerBackupResource, please refer to the document of PostgreSqlFlexibleServerBackupResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "postgresqltestserver";
string backupName = "backup_20210615T160516";
ResourceIdentifier postgreSqlFlexibleServerBackupResourceId = PostgreSqlFlexibleServerBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, backupName);
PostgreSqlFlexibleServerBackupResource postgreSqlFlexibleServerBackup = client.GetPostgreSqlFlexibleServerBackupResource(postgreSqlFlexibleServerBackupResourceId);

// invoke the operation
ArmOperation<PostgreSqlFlexibleServerBackupResource> lro = await postgreSqlFlexibleServerBackup.UpdateAsync(WaitUntil.Completed);
PostgreSqlFlexibleServerBackupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlFlexibleServerBackupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
