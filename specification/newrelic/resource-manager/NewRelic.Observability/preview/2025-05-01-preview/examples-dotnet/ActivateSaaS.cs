using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NewRelicObservability.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NewRelicObservability;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/ActivateSaaS.json
// this example is just showing the usage of "SaaS_ActivateResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
ActivateSaaSParameterContent content = new ActivateSaaSParameterContent("00000000-0000-0000-0000-000005430000", "publisherId");
NewRelicObservabilitySaaSResourceDetailsResult result = await subscriptionResource.ActivateResourceSaaSAsync(content);

Console.WriteLine($"Succeeded: {result}");
