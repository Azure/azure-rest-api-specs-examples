using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppConfiguration.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppConfiguration;

// Generated from example definition: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-05-01/examples/CheckNameAvailable.json
// this example is just showing the usage of "CheckAppConfigurationNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AppConfigurationNameAvailabilityContent content = new AppConfigurationNameAvailabilityContent("contoso", AppConfigurationResourceType.MicrosoftAppConfigurationConfigurationStores);
AppConfigurationNameAvailabilityResult result = await subscriptionResource.CheckAppConfigurationNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
