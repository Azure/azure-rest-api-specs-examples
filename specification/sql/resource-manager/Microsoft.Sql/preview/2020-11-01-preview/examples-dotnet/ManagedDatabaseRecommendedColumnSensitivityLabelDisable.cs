using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseRecommendedColumnSensitivityLabelDisable.json
// this example is just showing the usage of "ManagedDatabaseSensitivityLabels_DisableRecommendation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseColumnResource created on azure
// for more information of creating ManagedDatabaseColumnResource, please refer to the document of ManagedDatabaseColumnResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string managedInstanceName = "myManagedInstanceName";
string databaseName = "myDatabase";
string schemaName = "dbo";
string tableName = "myTable";
string columnName = "myColumn";
ResourceIdentifier managedDatabaseColumnResourceId = ManagedDatabaseColumnResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName);
ManagedDatabaseColumnResource managedDatabaseColumn = client.GetManagedDatabaseColumnResource(managedDatabaseColumnResourceId);

// invoke the operation
await managedDatabaseColumn.DisableRecommendationManagedDatabaseSensitivityLabelAsync();

Console.WriteLine($"Succeeded");
