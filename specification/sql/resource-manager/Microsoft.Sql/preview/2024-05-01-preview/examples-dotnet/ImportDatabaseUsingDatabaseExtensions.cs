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

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/ImportDatabaseUsingDatabaseExtensions.json
// this example is just showing the usage of "DatabaseExtensions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "17ca4d13-bf7d-4c33-a60e-b87a2820a325";
string resourceGroupName = "rg_062866bf-c4f4-41f9-abf0-b59132ca7924";
string serverName = "srv_2d6be2d2-26c8-4930-8fb6-82a5e95e0e82";
string databaseName = "db_2a47e946-e414-4c00-94ac-ed886bb78302";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// invoke the operation
string extensionName = "Import";
SqlDatabaseExtension sqlDatabaseExtension = new SqlDatabaseExtension
{
    OperationMode = DatabaseExtensionOperationMode.Import,
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
