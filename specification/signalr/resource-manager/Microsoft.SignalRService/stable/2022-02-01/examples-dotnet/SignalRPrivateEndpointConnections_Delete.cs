using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SignalR;
using Azure.ResourceManager.SignalR.Models;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRPrivateEndpointConnections_Delete.json
// this example is just showing the usage of "SignalRPrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SignalRPrivateEndpointConnectionResource created on azure
// for more information of creating SignalRPrivateEndpointConnectionResource, please refer to the document of SignalRPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
string privateEndpointConnectionName = "mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e";
ResourceIdentifier signalRPrivateEndpointConnectionResourceId = SignalRPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, privateEndpointConnectionName);
SignalRPrivateEndpointConnectionResource signalRPrivateEndpointConnection = client.GetSignalRPrivateEndpointConnectionResource(signalRPrivateEndpointConnectionResourceId);

// invoke the operation
await signalRPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
