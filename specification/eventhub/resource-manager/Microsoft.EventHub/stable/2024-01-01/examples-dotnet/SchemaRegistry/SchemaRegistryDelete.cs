using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/SchemaRegistry/SchemaRegistryDelete.json
// this example is just showing the usage of "SchemaRegistry_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsSchemaGroupResource created on azure
// for more information of creating EventHubsSchemaGroupResource, please refer to the document of EventHubsSchemaGroupResource
string subscriptionId = "e8baea74-64ce-459b-bee3-5aa4c47b3ae3";
string resourceGroupName = "alitest";
string namespaceName = "ali-ua-test-eh-system-1";
string schemaGroupName = "testSchemaGroup1";
ResourceIdentifier eventHubsSchemaGroupResourceId = EventHubsSchemaGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, schemaGroupName);
EventHubsSchemaGroupResource eventHubsSchemaGroup = client.GetEventHubsSchemaGroupResource(eventHubsSchemaGroupResourceId);

// invoke the operation
await eventHubsSchemaGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
