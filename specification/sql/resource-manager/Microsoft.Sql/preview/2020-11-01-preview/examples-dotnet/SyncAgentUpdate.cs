using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncAgentUpdate.json
// this example is just showing the usage of "SyncAgents_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SyncAgentResource created on azure
// for more information of creating SyncAgentResource, please refer to the document of SyncAgentResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "syncagentcrud-65440";
string serverName = "syncagentcrud-8475";
string syncAgentName = "syncagentcrud-3187";
ResourceIdentifier syncAgentResourceId = SyncAgentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, syncAgentName);
SyncAgentResource syncAgent = client.GetSyncAgentResource(syncAgentResourceId);

// invoke the operation
SyncAgentData data = new SyncAgentData()
{
    SyncDatabaseId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/syncagentcrud-8475/databases/sync"),
};
ArmOperation<SyncAgentResource> lro = await syncAgent.UpdateAsync(WaitUntil.Completed, data);
SyncAgentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SyncAgentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
