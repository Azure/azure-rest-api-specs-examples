using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.DataMigration;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/SqlMiCancelDatabaseMigration.json
// this example is just showing the usage of "DatabaseMigrationsSqlMi_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseMigrationSqlMIResource created on azure
// for more information of creating DatabaseMigrationSqlMIResource, please refer to the document of DatabaseMigrationSqlMIResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string managedInstanceName = "managedInstance1";
string targetDBName = "db1";
ResourceIdentifier databaseMigrationSqlMIResourceId = DatabaseMigrationSqlMIResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, targetDBName);
DatabaseMigrationSqlMIResource databaseMigrationSqlMI = client.GetDatabaseMigrationSqlMIResource(databaseMigrationSqlMIResourceId);

// invoke the operation
MigrationOperationInput input = new MigrationOperationInput
{
    MigrationOperationId = Guid.Parse("4124fe90-d1b6-4b50-b4d9-46d02381f59a"),
};
await databaseMigrationSqlMI.CancelAsync(WaitUntil.Completed, input);

Console.WriteLine("Succeeded");
