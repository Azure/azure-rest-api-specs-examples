using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StorageSync;
using Azure.ResourceManager.StorageSync.Models;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/StorageSyncServiceCheckNameAvailability_Available.json
// this example is just showing the usage of "StorageSyncServices_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "5c6bc8e1-1eaf-4192-94d8-58ce463ac86c";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
string locationName = "westus";
StorageSyncNameAvailabilityContent content = new StorageSyncNameAvailabilityContent("newstoragesyncservicename");
StorageSyncNameAvailabilityResult result = await subscriptionResource.CheckStorageSyncNameAvailabilityAsync(locationName, content);

Console.WriteLine($"Succeeded: {result}");
