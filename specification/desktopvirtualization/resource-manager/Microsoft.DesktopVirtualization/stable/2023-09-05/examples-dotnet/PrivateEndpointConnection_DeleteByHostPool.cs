using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/PrivateEndpointConnection_DeleteByHostPool.json
// this example is just showing the usage of "PrivateEndpointConnections_DeleteByHostPool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HostPoolPrivateEndpointConnectionResource created on azure
// for more information of creating HostPoolPrivateEndpointConnectionResource, please refer to the document of HostPoolPrivateEndpointConnectionResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string hostPoolName = "hostPool1";
string privateEndpointConnectionName = "hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b";
ResourceIdentifier hostPoolPrivateEndpointConnectionResourceId = HostPoolPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostPoolName, privateEndpointConnectionName);
HostPoolPrivateEndpointConnectionResource hostPoolPrivateEndpointConnection = client.GetHostPoolPrivateEndpointConnectionResource(hostPoolPrivateEndpointConnectionResourceId);

// invoke the operation
await hostPoolPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
