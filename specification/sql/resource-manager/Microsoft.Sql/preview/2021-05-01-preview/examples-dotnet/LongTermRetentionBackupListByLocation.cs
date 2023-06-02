using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupListByLocation.json
// this example is just showing the usage of "LongTermRetentionBackups_ListByLocation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
AzureLocation locationName = new AzureLocation("japaneast");
await foreach (LongTermRetentionBackupData item in subscriptionResource.GetLongTermRetentionBackupsWithLocationAsync(locationName))
{
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {item.Id}");
}

Console.WriteLine($"Succeeded");
