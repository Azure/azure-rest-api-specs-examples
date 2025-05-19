using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/ExportDatabaseUsingDatabaseExtensions.json
// this example is just showing the usage of "DatabaseExtensions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "0ca8cd24-0b47-4ad5-bc7e-d70e35c44adf";
string resourceGroupName = "rg_d1ef9eae-044d-4710-ba59-b82e84ad3157";
string serverName = "srv_9243d320-ac4e-4f97-8e06-b1167dae5f4c";
string databaseName = "db_7fe424c8-23cf-4ac3-bdc3-e21f424bdb68";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// invoke the operation
string extensionName = "Export";
SqlDatabaseExtension sqlDatabaseExtension = new SqlDatabaseExtension
{
    OperationMode = DatabaseExtensionOperationMode.Export,
    StorageKeyType = StorageKeyType.StorageAccessKey,
    StorageKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    StorageUri = new Uri("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml"),
    AdministratorLogin = "login",
    AdministratorLoginPassword = "password",
    AuthenticationType = "Sql",
};
ArmOperation<ImportExportExtensionsOperationResult> lro = await sqlDatabase.CreateOrUpdateDatabaseExtensionAsync(WaitUntil.Completed, extensionName, sqlDatabaseExtension);
ImportExportExtensionsOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
