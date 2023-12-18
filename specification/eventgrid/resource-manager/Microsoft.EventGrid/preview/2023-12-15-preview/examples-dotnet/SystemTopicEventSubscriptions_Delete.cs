using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/SystemTopicEventSubscriptions_Delete.json
// this example is just showing the usage of "SystemTopicEventSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SystemTopicEventSubscriptionResource created on azure
// for more information of creating SystemTopicEventSubscriptionResource, please refer to the document of SystemTopicEventSubscriptionResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string systemTopicName = "exampleSystemTopic1";
string eventSubscriptionName = "examplesubscription1";
ResourceIdentifier systemTopicEventSubscriptionResourceId = SystemTopicEventSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, systemTopicName, eventSubscriptionName);
SystemTopicEventSubscriptionResource systemTopicEventSubscription = client.GetSystemTopicEventSubscriptionResource(systemTopicEventSubscriptionResourceId);

// invoke the operation
await systemTopicEventSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
