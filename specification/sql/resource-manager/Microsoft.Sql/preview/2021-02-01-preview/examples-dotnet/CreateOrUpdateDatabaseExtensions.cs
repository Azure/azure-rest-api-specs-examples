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

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/CreateOrUpdateDatabaseExtensions.json
// this example is just showing the usage of "DatabaseExtensions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "a1c0814d-3c18-4e1e-a247-c128c12b1677";
string resourceGroupName = "rg_20cbe0f0-c2d9-4522-9177-5469aad53029";
string serverName = "srv_1ffd1cf8-9951-47fb-807d-a9c384763849";
string databaseName = "878e303f-1ea0-4f17-aa3d-a22ac5e9da08";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// invoke the operation
string extensionName = "polybaseimport";
SqlDatabaseExtension sqlDatabaseExtension = new SqlDatabaseExtension()
{
    OperationMode = DatabaseExtensionOperationMode.PolybaseImport,
    StorageKeyType = StorageKeyType.StorageAccessKey,
    StorageKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    StorageUri = new Uri("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml"),
};
ArmOperation<ImportExportExtensionsOperationResult> lro = await sqlDatabase.CreateOrUpdateDatabaseExtensionAsync(WaitUntil.Completed, extensionName, sqlDatabaseExtension);
ImportExportExtensionsOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
