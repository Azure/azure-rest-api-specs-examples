using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/private-endpoint-connection-delete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesPrivateEndpointConnectionResource created on azure
// for more information of creating MediaServicesPrivateEndpointConnectionResource, please refer to the document of MediaServicesPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso";
string accountName = "contososports";
string name = "connectionName1";
ResourceIdentifier mediaServicesPrivateEndpointConnectionResourceId = MediaServicesPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, name);
MediaServicesPrivateEndpointConnectionResource mediaServicesPrivateEndpointConnection = client.GetMediaServicesPrivateEndpointConnectionResource(mediaServicesPrivateEndpointConnectionResourceId);

// invoke the operation
await mediaServicesPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
