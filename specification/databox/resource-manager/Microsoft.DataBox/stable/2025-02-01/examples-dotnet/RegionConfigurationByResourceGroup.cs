using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBox.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DataBox;

// Generated from example definition: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/RegionConfigurationByResourceGroup.json
// this example is just showing the usage of "Service_RegionConfigurationByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "YourSubscriptionId";
string resourceGroupName = "YourResourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("westus");
RegionConfigurationContent content = new RegionConfigurationContent
{
    ScheduleAvailabilityRequest = new DataBoxScheduleAvailabilityContent(new AzureLocation("westus"))
    {
        Model = DeviceModelName.DataBox,
    },
    DeviceCapabilityRequest = new DeviceCapabilityContent
    {
        SkuName = DataBoxSkuName.DataBoxDisk,
        Model = DeviceModelName.DataBoxDisk,
    },
};
RegionConfigurationResult result = await resourceGroupResource.GetRegionConfigurationAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
