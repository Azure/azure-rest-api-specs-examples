using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Reservations.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Reservations;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetCatalog.json
// this example is just showing the usage of "GetCatalog" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "23bc208b-083f-4901-ae85-4f98c0c3b4b6";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
SubscriptionResourceGetCatalogOptions options = new SubscriptionResourceGetCatalogOptions { ReservedResourceType = "VirtualMachines", Location = new AzureLocation("eastus") };
await foreach (ReservationCatalog item in subscriptionResource.GetCatalogAsync(options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
