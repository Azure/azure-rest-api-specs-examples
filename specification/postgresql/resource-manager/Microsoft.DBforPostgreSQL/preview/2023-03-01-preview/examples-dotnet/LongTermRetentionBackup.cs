using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/LongTermRetentionBackup.json
// this example is just showing the usage of "FlexibleServer_StartLtrBackup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerResource created on azure
// for more information of creating PostgreSqlFlexibleServerResource, please refer to the document of PostgreSqlFlexibleServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "rgLongTermRetention";
string serverName = "pgsqlltrtestserver";
ResourceIdentifier postgreSqlFlexibleServerResourceId = PostgreSqlFlexibleServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
PostgreSqlFlexibleServerResource postgreSqlFlexibleServer = client.GetPostgreSqlFlexibleServerResource(postgreSqlFlexibleServerResourceId);

// invoke the operation
PostgreSqlFlexibleServerLtrBackupContent content = new PostgreSqlFlexibleServerLtrBackupContent(new PostgreSqlFlexibleServerBackupSettings("backup1"), new PostgreSqlFlexibleServerBackupStoreDetails(new string[]
{
"sasuri"
}));
ArmOperation<PostgreSqlFlexibleServerLtrBackupResult> lro = await postgreSqlFlexibleServer.StartLtrBackupFlexibleServerAsync(WaitUntil.Completed, content);
PostgreSqlFlexibleServerLtrBackupResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
