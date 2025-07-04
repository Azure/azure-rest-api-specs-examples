using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ApplicationGroup/ApplicationGroupCreate.json
// this example is just showing the usage of "ApplicationGroup_CreateOrUpdateApplicationGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsApplicationGroupResource created on azure
// for more information of creating EventHubsApplicationGroupResource, please refer to the document of EventHubsApplicationGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosotest";
string namespaceName = "contoso-ua-test-eh-system-1";
string applicationGroupName = "appGroup1";
ResourceIdentifier eventHubsApplicationGroupResourceId = EventHubsApplicationGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, applicationGroupName);
EventHubsApplicationGroupResource eventHubsApplicationGroup = client.GetEventHubsApplicationGroupResource(eventHubsApplicationGroupResourceId);

// invoke the operation
EventHubsApplicationGroupData data = new EventHubsApplicationGroupData
{
    IsEnabled = true,
    ClientAppGroupIdentifier = "SASKeyName=KeyName",
    Policies = { new EventHubsThrottlingPolicy("ThrottlingPolicy1", 7912L, EventHubsMetricId.IncomingMessages), new EventHubsThrottlingPolicy("ThrottlingPolicy2", 3951729L, EventHubsMetricId.IncomingBytes), new EventHubsThrottlingPolicy("ThrottlingPolicy3", 245175L, EventHubsMetricId.OutgoingBytes) },
};
ArmOperation<EventHubsApplicationGroupResource> lro = await eventHubsApplicationGroup.UpdateAsync(WaitUntil.Completed, data);
EventHubsApplicationGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsApplicationGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
