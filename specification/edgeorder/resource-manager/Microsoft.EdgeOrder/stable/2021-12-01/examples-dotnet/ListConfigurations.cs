using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EdgeOrder.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.EdgeOrder;

// Generated from example definition: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListConfigurations.json
// this example is just showing the usage of "ListConfigurations" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
ConfigurationsContent content = new ConfigurationsContent(new ConfigurationFilters[]
{
new ConfigurationFilters(new HierarchyInformation
{
ProductFamilyName = "AzureStackEdge",
ProductLineName = "AzureStackEdge",
ProductName = "AzureStackEdgeGPU",
})
{
FilterableProperty = {new FilterableProperty(SupportedFilterType.ShipToCountries, new string[]{"US"})},
}
});
await foreach (ProductConfiguration item in subscriptionResource.GetConfigurationsAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
