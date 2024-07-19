using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SelfHelp.Models;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/DiscoverSolutionsAtSubscriptionScope.json
// this example is just showing the usage of "DiscoverySolutionNLP_DiscoverSolutionsBySubscription" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "0d0fcd2e-c4fd-4349-8497-200edb3923c6";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
DiscoveryNlpContent content = new DiscoveryNlpContent("how to retrieve certs from deleted keyvault.")
{
    ResourceId = "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read",
    ServiceId = "0d0fcd2e-c4fd-4349-8497-200edb39s3ca",
};
await foreach (SolutionNlpMetadata item in subscriptionResource.DiscoverSolutionsNlpAsync(content: content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
