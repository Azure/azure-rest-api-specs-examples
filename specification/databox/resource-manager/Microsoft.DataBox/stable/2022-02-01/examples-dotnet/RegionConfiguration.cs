using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataBox;
using Azure.ResourceManager.DataBox.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/RegionConfiguration.json
// this example is just showing the usage of "Service_RegionConfiguration" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("westus");
RegionConfigurationContent content = new RegionConfigurationContent()
{
    ScheduleAvailabilityRequest = new DataBoxScheduleAvailabilityContent(new AzureLocation("westus")),
};
RegionConfigurationResult result = await subscriptionResource.GetRegionConfigurationAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
