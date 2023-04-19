using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventHubs;
using Azure.ResourceManager.EventHubs.Models;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/SchemaRegistry/SchemaRegistryGet.json
// this example is just showing the usage of "SchemaRegistry_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsNamespaceResource created on azure
// for more information of creating EventHubsNamespaceResource, please refer to the document of EventHubsNamespaceResource
string subscriptionId = "e8baea74-64ce-459b-bee3-5aa4c47b3ae3";
string resourceGroupName = "alitest";
string namespaceName = "ali-ua-test-eh-system-1";
ResourceIdentifier eventHubsNamespaceResourceId = EventHubsNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
EventHubsNamespaceResource eventHubsNamespace = client.GetEventHubsNamespaceResource(eventHubsNamespaceResourceId);

// get the collection of this EventHubsSchemaGroupResource
EventHubsSchemaGroupCollection collection = eventHubsNamespace.GetEventHubsSchemaGroups();

// invoke the operation
string schemaGroupName = "testSchemaGroup1";
bool result = await collection.ExistsAsync(schemaGroupName);

Console.WriteLine($"Succeeded: {result}");
