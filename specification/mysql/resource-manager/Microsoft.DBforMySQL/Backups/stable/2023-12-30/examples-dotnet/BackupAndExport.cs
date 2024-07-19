using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/stable/2023-12-30/examples/BackupAndExport.json
// this example is just showing the usage of "BackupAndExport_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerResource created on azure
// for more information of creating MySqlFlexibleServerResource, please refer to the document of MySqlFlexibleServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "mysqltestserver";
ResourceIdentifier mySqlFlexibleServerResourceId = MySqlFlexibleServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlFlexibleServerResource mySqlFlexibleServer = client.GetMySqlFlexibleServerResource(mySqlFlexibleServerResourceId);

// invoke the operation
MySqlFlexibleServerBackupAndExportContent content = new MySqlFlexibleServerBackupAndExportContent(new MySqlFlexibleServerBackupSettings("customer-backup-name"), new MySqlFlexibleServerFullBackupStoreDetails(new string[]
{
"sasuri1","sasuri2"
}));
ArmOperation<MySqlFlexibleServerBackupAndExportResult> lro = await mySqlFlexibleServer.CreateBackupAndExportAsync(WaitUntil.Completed, content);
MySqlFlexibleServerBackupAndExportResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
