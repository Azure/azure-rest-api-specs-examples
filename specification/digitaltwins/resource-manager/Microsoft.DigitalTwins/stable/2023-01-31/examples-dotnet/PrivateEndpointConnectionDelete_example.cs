using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DigitalTwins;
using Azure.ResourceManager.DigitalTwins.Models;

// Generated from example definition: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/PrivateEndpointConnectionDelete_example.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DigitalTwinsPrivateEndpointConnectionResource created on azure
// for more information of creating DigitalTwinsPrivateEndpointConnectionResource, please refer to the document of DigitalTwinsPrivateEndpointConnectionResource
string subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
string resourceGroupName = "resRg";
string resourceName = "myDigitalTwinsService";
string privateEndpointConnectionName = "myPrivateConnection";
ResourceIdentifier digitalTwinsPrivateEndpointConnectionResourceId = DigitalTwinsPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, privateEndpointConnectionName);
DigitalTwinsPrivateEndpointConnectionResource digitalTwinsPrivateEndpointConnection = client.GetDigitalTwinsPrivateEndpointConnectionResource(digitalTwinsPrivateEndpointConnectionResourceId);

// invoke the operation
await digitalTwinsPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
