using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/LocationCheckNameAvailability_Available.json
// this example is just showing the usage of "Location_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation locationName = new AzureLocation("japaneast");
BatchNameAvailabilityContent content = new BatchNameAvailabilityContent("newaccountname");
BatchNameAvailabilityResult result = await subscriptionResource.CheckBatchNameAvailabilityAsync(locationName, content);

Console.WriteLine($"Succeeded: {result}");
