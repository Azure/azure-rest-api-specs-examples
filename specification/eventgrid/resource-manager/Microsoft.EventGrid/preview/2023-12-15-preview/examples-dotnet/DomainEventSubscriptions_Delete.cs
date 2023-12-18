using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/DomainEventSubscriptions_Delete.json
// this example is just showing the usage of "DomainEventSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DomainEventSubscriptionResource created on azure
// for more information of creating DomainEventSubscriptionResource, please refer to the document of DomainEventSubscriptionResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string domainName = "exampleDomain1";
string eventSubscriptionName = "examplesubscription1";
ResourceIdentifier domainEventSubscriptionResourceId = DomainEventSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, domainName, eventSubscriptionName);
DomainEventSubscriptionResource domainEventSubscription = client.GetDomainEventSubscriptionResource(domainEventSubscriptionResourceId);

// invoke the operation
await domainEventSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
