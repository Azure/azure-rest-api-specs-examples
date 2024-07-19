using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/DeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

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
