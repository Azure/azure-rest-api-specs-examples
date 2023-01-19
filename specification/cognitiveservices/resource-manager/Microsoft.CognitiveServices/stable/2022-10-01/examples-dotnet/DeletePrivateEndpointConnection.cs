using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/DeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this CognitiveServicesPrivateEndpointConnectionResource created on azure
// for more information of creating CognitiveServicesPrivateEndpointConnectionResource, please refer to the document of CognitiveServicesPrivateEndpointConnectionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string accountName = "sto2527";
string privateEndpointConnectionName = "{privateEndpointConnectionName}";
ResourceIdentifier cognitiveServicesPrivateEndpointConnectionResourceId = CognitiveServicesPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, privateEndpointConnectionName);
CognitiveServicesPrivateEndpointConnectionResource cognitiveServicesPrivateEndpointConnection = client.GetCognitiveServicesPrivateEndpointConnectionResource(cognitiveServicesPrivateEndpointConnectionResourceId);

// invoke the operation
await cognitiveServicesPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
