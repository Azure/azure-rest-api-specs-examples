using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/SyncMemberGetSchema.json
// this example is just showing the usage of "SyncMembers_ListMemberSchemas" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SyncMemberResource created on azure
// for more information of creating SyncMemberResource, please refer to the document of SyncMemberResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "syncgroupcrud-65440";
string serverName = "syncgroupcrud-8475";
string databaseName = "syncgroupcrud-4328";
string syncGroupName = "syncgroupcrud-3187";
string syncMemberName = "syncgroupcrud-4879";
ResourceIdentifier syncMemberResourceId = SyncMemberResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName);
SyncMemberResource syncMember = client.GetSyncMemberResource(syncMemberResourceId);

// invoke the operation and iterate over the result
await foreach (SyncFullSchemaProperties item in syncMember.GetMemberSchemasAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
