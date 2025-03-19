using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DatabaseWatcher.Models;
using Azure.ResourceManager.DatabaseWatcher;

// Generated from example definition: 2025-01-02/SharedPrivateLinkResources_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "SharedPrivateLinkResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseWatcherSharedPrivateLinkResource created on azure
// for more information of creating DatabaseWatcherSharedPrivateLinkResource, please refer to the document of DatabaseWatcherSharedPrivateLinkResource
string subscriptionId = "49e0fbd3-75e8-44e7-96fd-5b64d9ad818d";
string resourceGroupName = "apiTest-ddat4p";
string watcherName = "databasemo3ej9ih";
string sharedPrivateLinkResourceName = "monitoringh22eed";
ResourceIdentifier databaseWatcherSharedPrivateLinkResourceId = DatabaseWatcherSharedPrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, watcherName, sharedPrivateLinkResourceName);
DatabaseWatcherSharedPrivateLinkResource databaseWatcherSharedPrivateLinkResource = client.GetDatabaseWatcherSharedPrivateLinkResource(databaseWatcherSharedPrivateLinkResourceId);

// invoke the operation
await databaseWatcherSharedPrivateLinkResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
