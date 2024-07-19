using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/Migrations_Create_With_Other_Users.json
// this example is just showing the usage of "Migrations_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerResource created on azure
// for more information of creating PostgreSqlFlexibleServerResource, please refer to the document of PostgreSqlFlexibleServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string targetDbServerName = "testtarget";
ResourceIdentifier postgreSqlFlexibleServerResourceId = PostgreSqlFlexibleServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, targetDbServerName);
PostgreSqlFlexibleServerResource postgreSqlFlexibleServer = client.GetPostgreSqlFlexibleServerResource(postgreSqlFlexibleServerResourceId);

// get the collection of this PostgreSqlMigrationResource
PostgreSqlMigrationCollection collection = postgreSqlFlexibleServer.GetPostgreSqlMigrations();

// invoke the operation
string migrationName = "testmigration";
PostgreSqlMigrationData data = new PostgreSqlMigrationData(new AzureLocation("westus"))
{
    MigrationMode = PostgreSqlMigrationMode.Offline,
    SourceDbServerResourceId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/servers/testsource"),
    SecretParameters = new PostgreSqlMigrationSecretParameters(new PostgreSqlMigrationAdminCredentials("xxxxxxxx", "xxxxxxxx"))
    {
        SourceServerUsername = "newadmin@testsource",
        TargetServerUsername = "targetadmin",
    },
    DbsToMigrate =
    {
    "db1","db2","db3","db4"
    },
};
ArmOperation<PostgreSqlMigrationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, migrationName, data);
PostgreSqlMigrationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlMigrationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
