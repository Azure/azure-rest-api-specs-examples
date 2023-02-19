using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopicEventSubscriptions_Get.json
// this example is just showing the usage of "SystemTopicEventSubscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SystemTopicResource created on azure
// for more information of creating SystemTopicResource, please refer to the document of SystemTopicResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string systemTopicName = "exampleSystemTopic1";
ResourceIdentifier systemTopicResourceId = SystemTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, systemTopicName);
SystemTopicResource systemTopic = client.GetSystemTopicResource(systemTopicResourceId);

// get the collection of this SystemTopicEventSubscriptionResource
SystemTopicEventSubscriptionCollection collection = systemTopic.GetSystemTopicEventSubscriptions();

// invoke the operation
string eventSubscriptionName = "examplesubscription1";
bool result = await collection.ExistsAsync(eventSubscriptionName);

Console.WriteLine($"Succeeded: {result}");
