using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/SyncMemberUpdate.json
// this example is just showing the usage of "SyncMembers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SyncGroupResource created on azure
// for more information of creating SyncGroupResource, please refer to the document of SyncGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "syncgroupcrud-65440";
string serverName = "syncgroupcrud-8475";
string databaseName = "syncgroupcrud-4328";
string syncGroupName = "syncgroupcrud-3187";
ResourceIdentifier syncGroupResourceId = SyncGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, syncGroupName);
SyncGroupResource syncGroup = client.GetSyncGroupResource(syncGroupResourceId);

// get the collection of this SyncMemberResource
SyncMemberCollection collection = syncGroup.GetSyncMembers();

// invoke the operation
string syncMemberName = "syncmembercrud-4879";
SyncMemberData data = new SyncMemberData
{
    DatabaseType = SyncMemberDbType.AzureSqlDatabase,
    SyncMemberAzureDatabaseResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
    UsePrivateLinkConnection = true,
    ServerName = "syncgroupcrud-3379.database.windows.net",
    DatabaseName = "syncgroupcrud-7421",
    UserName = "myUser",
    SyncDirection = SyncDirection.Bidirectional,
};
ArmOperation<SyncMemberResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, syncMemberName, data);
SyncMemberResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SyncMemberData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
