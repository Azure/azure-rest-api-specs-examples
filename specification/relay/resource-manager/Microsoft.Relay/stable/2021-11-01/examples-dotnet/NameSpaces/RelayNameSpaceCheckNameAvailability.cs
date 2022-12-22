using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/NameSpaces/RelayNameSpaceCheckNameAvailability.json
// this example is just showing the usage of "Namespaces_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
RelayNameAvailabilityContent content = new RelayNameAvailabilityContent("example-RelayNamespace1321");
RelayNameAvailabilityResult result = await subscriptionResource.CheckRelayNamespaceNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
