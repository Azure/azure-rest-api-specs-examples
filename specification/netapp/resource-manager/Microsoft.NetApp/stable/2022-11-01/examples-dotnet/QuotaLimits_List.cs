using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/QuotaLimits_List.json
// this example is just showing the usage of "NetAppResourceQuotaLimits_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
AzureLocation location = new AzureLocation("eastus");
await foreach (NetAppSubscriptionQuotaItem item in subscriptionResource.GetNetAppQuotaLimitsAsync(location))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
