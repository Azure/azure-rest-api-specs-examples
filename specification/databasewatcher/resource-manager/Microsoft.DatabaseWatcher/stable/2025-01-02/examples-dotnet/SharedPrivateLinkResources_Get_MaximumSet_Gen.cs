using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DatabaseWatcher.Models;
using Azure.ResourceManager.DatabaseWatcher;

// Generated from example definition: 2025-01-02/SharedPrivateLinkResources_Get_MaximumSet_Gen.json
// this example is just showing the usage of "SharedPrivateLinkResource_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseWatcherResource created on azure
// for more information of creating DatabaseWatcherResource, please refer to the document of DatabaseWatcherResource
string subscriptionId = "49e0fbd3-75e8-44e7-96fd-5b64d9ad818d";
string resourceGroupName = "apiTest-ddat4p";
string watcherName = "databasemo3ej9ih";
ResourceIdentifier databaseWatcherResourceId = DatabaseWatcherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, watcherName);
DatabaseWatcherResource databaseWatcher = client.GetDatabaseWatcherResource(databaseWatcherResourceId);

// get the collection of this DatabaseWatcherSharedPrivateLinkResource
DatabaseWatcherSharedPrivateLinkResourceCollection collection = databaseWatcher.GetDatabaseWatcherSharedPrivateLinkResources();

// invoke the operation
string sharedPrivateLinkResourceName = "monitoringh22eed";
NullableResponse<DatabaseWatcherSharedPrivateLinkResource> response = await collection.GetIfExistsAsync(sharedPrivateLinkResourceName);
DatabaseWatcherSharedPrivateLinkResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DatabaseWatcherSharedPrivateLinkResourceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
