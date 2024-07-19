using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupRefreshHubSchema.json
// this example is just showing the usage of "SyncGroups_RefreshHubSchema" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
await syncGroup.RefreshHubSchemaAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
