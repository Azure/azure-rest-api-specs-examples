using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ImportNewDatabaseWithNetworkIsolation.json
// this example is just showing the usage of "Servers_ImportDatabase" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerResource created on azure
// for more information of creating SqlServerResource, please refer to the document of SqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string serverName = "testsvr";
ResourceIdentifier sqlServerResourceId = SqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
SqlServerResource sqlServer = client.GetSqlServerResource(sqlServerResourceId);

// invoke the operation
DatabaseImportDefinition databaseImportDefinition = new DatabaseImportDefinition(StorageKeyType.StorageAccessKey, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx==", new Uri("https://test.blob.core.windows.net/test.bacpac"), "login", "password")
{
    DatabaseName = "testdb",
    AuthenticationType = "Sql",
    NetworkIsolation = new NetworkIsolationSettings
    {
        StorageAccountResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Storage/storageAccounts/test-privatelink"),
        SqlServerResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr"),
    },
};
ArmOperation<ImportExportOperationResult> lro = await sqlServer.ImportDatabaseAsync(WaitUntil.Completed, databaseImportDefinition);
ImportExportOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
