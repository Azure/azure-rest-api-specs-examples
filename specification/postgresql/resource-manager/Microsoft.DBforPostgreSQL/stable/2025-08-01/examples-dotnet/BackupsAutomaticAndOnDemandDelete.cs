using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/BackupsAutomaticAndOnDemandDelete.json
// this example is just showing the usage of "BackupsAutomaticAndOnDemand_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerBackupResource created on azure
// for more information of creating PostgreSqlFlexibleServerBackupResource, please refer to the document of PostgreSqlFlexibleServerBackupResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "exampleresourcegroup";
string serverName = "exampleserver";
string backupName = "ondemandbackup-20250601T183022";
ResourceIdentifier postgreSqlFlexibleServerBackupResourceId = PostgreSqlFlexibleServerBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, backupName);
PostgreSqlFlexibleServerBackupResource postgreSqlFlexibleServerBackup = client.GetPostgreSqlFlexibleServerBackupResource(postgreSqlFlexibleServerBackupResourceId);

// invoke the operation
await postgreSqlFlexibleServerBackup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
