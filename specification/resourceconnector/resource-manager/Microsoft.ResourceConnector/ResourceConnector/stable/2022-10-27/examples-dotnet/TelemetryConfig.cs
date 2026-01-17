using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceConnector.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ResourceConnector;

// Generated from example definition: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/ResourceConnector/stable/2022-10-27/examples/TelemetryConfig.json
// this example is just showing the usage of "Appliances_GetTelemetryConfig" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
ApplianceTelemetryConfigResult result = await subscriptionResource.GetApplianceTelemetryConfigAsync();

Console.WriteLine($"Succeeded: {result}");
