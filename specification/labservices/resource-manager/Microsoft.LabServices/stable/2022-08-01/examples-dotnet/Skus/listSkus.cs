using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.LabServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.LabServices;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Skus/listSkus.json
// this example is just showing the usage of "Skus_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
await foreach (AvailableLabServicesSku item in subscriptionResource.GetSkusAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
