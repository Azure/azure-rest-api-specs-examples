using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/ReplicationLinkDelete.json
// this example is just showing the usage of "ReplicationLinks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerDatabaseReplicationLinkResource created on azure
// for more information of creating SqlServerDatabaseReplicationLinkResource, please refer to the document of SqlServerDatabaseReplicationLinkResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "sourcesvr";
string databaseName = "gamma-db";
string linkId = "4891ca10-ebd0-47d7-9182-c722651780fb";
ResourceIdentifier sqlServerDatabaseReplicationLinkResourceId = SqlServerDatabaseReplicationLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, linkId);
SqlServerDatabaseReplicationLinkResource sqlServerDatabaseReplicationLink = client.GetSqlServerDatabaseReplicationLinkResource(sqlServerDatabaseReplicationLinkResourceId);

// invoke the operation
await sqlServerDatabaseReplicationLink.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
