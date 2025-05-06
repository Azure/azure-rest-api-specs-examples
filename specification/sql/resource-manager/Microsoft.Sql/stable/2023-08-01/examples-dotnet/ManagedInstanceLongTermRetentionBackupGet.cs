using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/ManagedInstanceLongTermRetentionBackupGet.json
// this example is just showing the usage of "LongTermRetentionManagedInstanceBackups_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this SubscriptionLongTermRetentionManagedInstanceBackupResource
AzureLocation locationName = new AzureLocation("japaneast");
string managedInstanceName = "testInstance";
string databaseName = "testDatabase";
SubscriptionLongTermRetentionManagedInstanceBackupCollection collection = subscriptionResource.GetSubscriptionLongTermRetentionManagedInstanceBackups(locationName, managedInstanceName, databaseName);

// invoke the operation
string backupName = "55555555-6666-7777-8888-999999999999;131637960820000000;Archive";
NullableResponse<SubscriptionLongTermRetentionManagedInstanceBackupResource> response = await collection.GetIfExistsAsync(backupName);
SubscriptionLongTermRetentionManagedInstanceBackupResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ManagedInstanceLongTermRetentionBackupData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
