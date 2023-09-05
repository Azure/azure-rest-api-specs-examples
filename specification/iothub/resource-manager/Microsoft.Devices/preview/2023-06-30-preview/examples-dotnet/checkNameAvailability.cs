using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotHub;
using Azure.ResourceManager.IotHub.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/checkNameAvailability.json
// this example is just showing the usage of "IotHubResource_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
IotHubNameAvailabilityContent content = new IotHubNameAvailabilityContent("test-request");
IotHubNameAvailabilityResponse result = await subscriptionResource.CheckIotHubNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
