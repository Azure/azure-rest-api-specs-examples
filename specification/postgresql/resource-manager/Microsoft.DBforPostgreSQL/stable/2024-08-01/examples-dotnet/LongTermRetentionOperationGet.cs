using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/LongTermRetentionOperationGet.json
// this example is just showing the usage of "ltrBackupOperations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlLtrServerBackupOperationResource created on azure
// for more information of creating PostgreSqlLtrServerBackupOperationResource, please refer to the document of PostgreSqlLtrServerBackupOperationResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "rgLongTermRetention";
string serverName = "pgsqlltrtestserver";
string backupName = "backup1";
ResourceIdentifier postgreSqlLtrServerBackupOperationResourceId = PostgreSqlLtrServerBackupOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, backupName);
PostgreSqlLtrServerBackupOperationResource postgreSqlLtrServerBackupOperation = client.GetPostgreSqlLtrServerBackupOperationResource(postgreSqlLtrServerBackupOperationResourceId);

// invoke the operation
PostgreSqlLtrServerBackupOperationResource result = await postgreSqlLtrServerBackupOperation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlLtrServerBackupOperationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
