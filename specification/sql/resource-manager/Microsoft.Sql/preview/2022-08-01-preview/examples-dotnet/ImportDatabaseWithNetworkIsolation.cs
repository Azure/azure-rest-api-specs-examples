using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ImportDatabaseWithNetworkIsolation.json
// this example is just showing the usage of "Databases_Import" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string serverName = "testsvr";
string databaseName = "testdb";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// invoke the operation
ImportExistingDatabaseDefinition importExistingDatabaseDefinition = new ImportExistingDatabaseDefinition(StorageKeyType.StorageAccessKey, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx==", new Uri("https://test.blob.core.windows.net/test.bacpac"), "login", "password")
{
    AuthenticationType = "Sql",
    NetworkIsolation = new NetworkIsolationSettings()
    {
        StorageAccountResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Storage/storageAccounts/test-privatelink"),
        SqlServerResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr"),
    },
};
ArmOperation<ImportExportOperationResult> lro = await sqlDatabase.ImportAsync(WaitUntil.Completed, importExistingDatabaseDefinition);
ImportExportOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
