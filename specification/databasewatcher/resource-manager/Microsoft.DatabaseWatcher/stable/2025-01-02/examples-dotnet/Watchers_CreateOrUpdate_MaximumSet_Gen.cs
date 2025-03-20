using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DatabaseWatcher.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DatabaseWatcher;

// Generated from example definition: 2025-01-02/Watchers_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "Watcher_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "A76F9850-996B-40B3-94D4-C98110A0EEC9";
string resourceGroupName = "rgWatcher";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DatabaseWatcherResource
DatabaseWatcherCollection collection = resourceGroupResource.GetDatabaseWatchers();

// invoke the operation
string watcherName = "testWatcher";
DatabaseWatcherData data = new DatabaseWatcherData(new AzureLocation("eastus2"))
{
    Properties = new DatabaseWatcherProperties
    {
        Datastore = new DatabaseWatcherDatastore(new Uri("https://kustouri-adx.eastus.kusto.windows.net"), new Uri("https://ingest-kustouri-adx.eastus.kusto.windows.net"), "kustoDatabaseName1", new Uri("https://portal.azure.com/"), KustoOfferingType.Adx)
        {
            AdxClusterResourceId = new ResourceIdentifier("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
            KustoClusterDisplayName = "kustoUri-adx",
        },
        DefaultAlertRuleIdentityResourceId = new ResourceIdentifier("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
    },
    Identity = (ManagedServiceIdentity)null,
    Tags = { },
};
ArmOperation<DatabaseWatcherResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, watcherName, data);
DatabaseWatcherResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatabaseWatcherData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
