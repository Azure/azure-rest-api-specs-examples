using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/CheckNameAvailability.json
// this example is just showing the usage of "CheckNameAvailability_Execute" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
DevCenterNameAvailabilityContent content = new DevCenterNameAvailabilityContent
{
    Name = "name1",
    ResourceType = new ResourceType("Microsoft.DevCenter/devcenters"),
};
DevCenterNameAvailabilityResult result = await subscriptionResource.CheckDevCenterNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
