using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/CheckDomainAvailability.json
// this example is just showing the usage of "CheckDomainAvailability" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
CognitiveServicesDomainAvailabilityContent content = new CognitiveServicesDomainAvailabilityContent("contosodemoapp1", new ResourceType("Microsoft.CognitiveServices/accounts"));
CognitiveServicesDomainAvailabilityList result = await subscriptionResource.CheckDomainAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
