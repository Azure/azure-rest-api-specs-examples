using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/LocationListVirtualMachineSkus.json
// this example is just showing the usage of "Location_ListSupportedVirtualMachineSkus" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subid";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
AzureLocation locationName = new AzureLocation("japaneast");
await foreach (BatchSupportedSku item in subscriptionResource.GetBatchSupportedVirtualMachineSkusAsync(locationName))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
