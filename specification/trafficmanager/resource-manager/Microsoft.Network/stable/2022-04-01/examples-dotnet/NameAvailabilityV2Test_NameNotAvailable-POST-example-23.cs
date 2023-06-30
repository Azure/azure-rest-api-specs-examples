using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.TrafficManager;
using Azure.ResourceManager.TrafficManager.Models;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/NameAvailabilityV2Test_NameNotAvailable-POST-example-23.json
// this example is just showing the usage of "Profiles_checkTrafficManagerNameAvailabilityV2" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
TrafficManagerRelativeDnsNameAvailabilityContent content = new TrafficManagerRelativeDnsNameAvailabilityContent()
{
    Name = "azsmnet4696",
    ResourceType = new ResourceType("microsoft.network/trafficmanagerprofiles"),
};
TrafficManagerNameAvailabilityResult result = await subscriptionResource.CheckTrafficManagerNameAvailabilityV2Async(content);

Console.WriteLine($"Succeeded: {result}");
