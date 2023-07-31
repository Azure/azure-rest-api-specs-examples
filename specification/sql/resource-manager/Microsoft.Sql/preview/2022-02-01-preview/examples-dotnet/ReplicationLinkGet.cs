using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ReplicationLinkGet.json
// this example is just showing the usage of "ReplicationLinks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "sourcesvr";
string databaseName = "gamma-db";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// get the collection of this SqlServerDatabaseReplicationLinkResource
SqlServerDatabaseReplicationLinkCollection collection = sqlDatabase.GetSqlServerDatabaseReplicationLinks();

// invoke the operation
string linkId = "4891ca10-ebd0-47d7-9182-c722651780fb";
bool result = await collection.ExistsAsync(linkId);

Console.WriteLine($"Succeeded: {result}");
