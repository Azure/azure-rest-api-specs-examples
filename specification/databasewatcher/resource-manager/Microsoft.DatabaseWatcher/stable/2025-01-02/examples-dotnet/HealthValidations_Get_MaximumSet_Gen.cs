using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DatabaseWatcher;

// Generated from example definition: 2025-01-02/HealthValidations_Get_MaximumSet_Gen.json
// this example is just showing the usage of "HealthValidation_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this DatabaseWatcherHealthValidationResource
DatabaseWatcherHealthValidationCollection collection = databaseWatcher.GetDatabaseWatcherHealthValidations();

// invoke the operation
string healthValidationName = "testHealthValidation";
NullableResponse<DatabaseWatcherHealthValidationResource> response = await collection.GetIfExistsAsync(healthValidationName);
DatabaseWatcherHealthValidationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DatabaseWatcherHealthValidationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
