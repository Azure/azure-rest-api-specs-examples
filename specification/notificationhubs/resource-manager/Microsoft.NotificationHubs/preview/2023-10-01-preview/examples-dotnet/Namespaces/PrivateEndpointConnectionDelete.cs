using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NotificationHubs.Models;
using Azure.ResourceManager.NotificationHubs;

// Generated from example definition: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/Namespaces/PrivateEndpointConnectionDelete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NotificationHubPrivateEndpointConnectionResource created on azure
// for more information of creating NotificationHubPrivateEndpointConnectionResource, please refer to the document of NotificationHubPrivateEndpointConnectionResource
string subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
string resourceGroupName = "5ktrial";
string namespaceName = "nh-sdk-ns";
string privateEndpointConnectionName = "nh-sdk-ns.1fa229cd-bf3f-47f0-8c49-afb36723997e";
ResourceIdentifier notificationHubPrivateEndpointConnectionResourceId = NotificationHubPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, privateEndpointConnectionName);
NotificationHubPrivateEndpointConnectionResource notificationHubPrivateEndpointConnection = client.GetNotificationHubPrivateEndpointConnectionResource(notificationHubPrivateEndpointConnectionResourceId);

// invoke the operation
await notificationHubPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
