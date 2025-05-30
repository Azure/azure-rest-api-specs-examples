using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DataMigration;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/SqlDbCreateOrUpdateDatabaseMigrationMIN.json
// this example is just showing the usage of "DatabaseMigrationsSqlDb_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DatabaseMigrationSqlDBResource
DatabaseMigrationSqlDBCollection collection = resourceGroupResource.GetDatabaseMigrationSqlDBs();

// invoke the operation
string sqlDBInstanceName = "sqldbinstance";
string targetDBName = "db1";
DatabaseMigrationSqlDBData data = new DatabaseMigrationSqlDBData
{
    Properties = new DatabaseMigrationSqlDBProperties
    {
        TargetSqlConnection = new SqlConnectionInformation
        {
            DataSource = "sqldbinstance",
            Authentication = "SqlAuthentication",
            UserName = "bbb",
            Password = "placeholder",
            EncryptConnection = true,
            TrustServerCertificate = true,
        },
        Scope = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/servers/sqldbinstance",
        SourceSqlConnection = new SqlConnectionInformation
        {
            DataSource = "aaa",
            Authentication = "WindowsAuthentication",
            UserName = "bbb",
            Password = "placeholder",
            EncryptConnection = true,
            TrustServerCertificate = true,
        },
        SourceDatabaseName = "aaa",
        MigrationService = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent",
    },
};
ArmOperation<DatabaseMigrationSqlDBResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sqlDBInstanceName, targetDBName, data);
DatabaseMigrationSqlDBResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatabaseMigrationSqlDBData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
