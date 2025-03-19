using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DatabaseWatcher.Models;
using Azure.ResourceManager.DatabaseWatcher;

// Generated from example definition: 2025-01-02/Targets_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "Target_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseWatcherTargetResource created on azure
// for more information of creating DatabaseWatcherTargetResource, please refer to the document of DatabaseWatcherTargetResource
string subscriptionId = "49e0fbd3-75e8-44e7-96fd-5b64d9ad818d";
string resourceGroupName = "apiTest-ddat4p";
string watcherName = "databasemo3ej9ih";
string targetName = "monitoringh22eed";
ResourceIdentifier databaseWatcherTargetResourceId = DatabaseWatcherTargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, watcherName, targetName);
DatabaseWatcherTargetResource databaseWatcherTarget = client.GetDatabaseWatcherTargetResource(databaseWatcherTargetResourceId);

// invoke the operation
DatabaseWatcherTargetData data = new DatabaseWatcherTargetData
{
    Properties = new SqlDBSingleDatabaseTargetProperties(TargetAuthenticationType.Aad, "sqlServero1ihe2", new ResourceIdentifier("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.Sql/servers/m1/databases/m2")),
};
ArmOperation<DatabaseWatcherTargetResource> lro = await databaseWatcherTarget.UpdateAsync(WaitUntil.Completed, data);
DatabaseWatcherTargetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatabaseWatcherTargetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
