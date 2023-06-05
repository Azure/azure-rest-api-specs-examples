using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventHubs;
using Azure.ResourceManager.EventHubs.Models;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/EventHubs/EHEventHubAuthorizationRuleCreate.json
// this example is just showing the usage of "EventHubs_CreateOrUpdateAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubResource created on azure
// for more information of creating EventHubResource, please refer to the document of EventHubResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-960";
string eventHubName = "sdk-EventHub-532";
ResourceIdentifier eventHubResourceId = EventHubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, eventHubName);
EventHubResource eventHub = client.GetEventHubResource(eventHubResourceId);

// get the collection of this EventHubAuthorizationRuleResource
EventHubAuthorizationRuleCollection collection = eventHub.GetEventHubAuthorizationRules();

// invoke the operation
string authorizationRuleName = "sdk-Authrules-2513";
EventHubsAuthorizationRuleData data = new EventHubsAuthorizationRuleData()
{
    Rights =
    {
    EventHubsAccessRight.Listen,EventHubsAccessRight.Send
    },
};
ArmOperation<EventHubAuthorizationRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, authorizationRuleName, data);
EventHubAuthorizationRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsAuthorizationRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
