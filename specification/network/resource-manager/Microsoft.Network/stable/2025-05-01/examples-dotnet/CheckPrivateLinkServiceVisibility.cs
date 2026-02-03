using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/CheckPrivateLinkServiceVisibility.json
// this example is just showing the usage of "PrivateLinkServices_CheckPrivateLinkServiceVisibility" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subid";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("westus");
CheckPrivateLinkServiceVisibilityRequest checkPrivateLinkServiceVisibilityRequest = new CheckPrivateLinkServiceVisibilityRequest
{
    PrivateLinkServiceAlias = "mypls.00000000-0000-0000-0000-000000000000.azure.privatelinkservice",
};
ArmOperation<PrivateLinkServiceVisibility> lro = await subscriptionResource.CheckPrivateLinkServiceVisibilityPrivateLinkServiceAsync(WaitUntil.Completed, location, checkPrivateLinkServiceVisibilityRequest);
PrivateLinkServiceVisibility result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
