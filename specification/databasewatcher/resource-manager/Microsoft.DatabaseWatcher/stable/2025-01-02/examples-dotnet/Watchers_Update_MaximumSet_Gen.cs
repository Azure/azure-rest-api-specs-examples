using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DatabaseWatcher.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.DatabaseWatcher;

// Generated from example definition: 2025-01-02/Watchers_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Watcher_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseWatcherResource created on azure
// for more information of creating DatabaseWatcherResource, please refer to the document of DatabaseWatcherResource
string subscriptionId = "A76F9850-996B-40B3-94D4-C98110A0EEC9";
string resourceGroupName = "rgWatcher";
string watcherName = "testWatcher";
ResourceIdentifier databaseWatcherResourceId = DatabaseWatcherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, watcherName);
DatabaseWatcherResource databaseWatcher = client.GetDatabaseWatcherResource(databaseWatcherResourceId);

// invoke the operation
DatabaseWatcherPatch patch = new DatabaseWatcherPatch
{
    Identity = (ManagedServiceIdentity)null,
    Tags = { },
    Properties = new DatabaseWatcherUpdateProperties
    {
        Datastore = new DatabaseWatcherDatastore(new Uri("https://kustouri-adx.eastus.kusto.windows.net"), new Uri("https://ingest-kustouri-adx.eastus.kusto.windows.net"), "kustoDatabaseName1", new Uri("https://portal.azure.com/"), KustoOfferingType.Adx)
        {
            AdxClusterResourceId = new ResourceIdentifier("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
            KustoClusterDisplayName = "kustoUri-adx",
        },
        DefaultAlertRuleIdentityResourceId = new ResourceIdentifier("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/newtest"),
    },
};
ArmOperation<DatabaseWatcherResource> lro = await databaseWatcher.UpdateAsync(WaitUntil.Completed, patch);
DatabaseWatcherResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatabaseWatcherData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
