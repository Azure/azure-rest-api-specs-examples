using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NotificationHubs.Models;
using Azure.ResourceManager.NotificationHubs;

// Generated from example definition: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/Namespaces/PrivateEndpointConnectionUpdate.json
// this example is just showing the usage of "PrivateEndpointConnections_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NotificationHubNamespaceResource created on azure
// for more information of creating NotificationHubNamespaceResource, please refer to the document of NotificationHubNamespaceResource
string subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
string resourceGroupName = "5ktrial";
string namespaceName = "nh-sdk-ns";
ResourceIdentifier notificationHubNamespaceResourceId = NotificationHubNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
NotificationHubNamespaceResource notificationHubNamespace = client.GetNotificationHubNamespaceResource(notificationHubNamespaceResourceId);

// get the collection of this NotificationHubPrivateEndpointConnectionResource
NotificationHubPrivateEndpointConnectionCollection collection = notificationHubNamespace.GetNotificationHubPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "nh-sdk-ns.1fa229cd-bf3f-47f0-8c49-afb36723997e";
NotificationHubPrivateEndpointConnectionData data = new NotificationHubPrivateEndpointConnectionData
{
    Properties = new NotificationHubPrivateEndpointConnectionProperties
    {
        ConnectionState = new RemotePrivateLinkServiceConnectionState
        {
            Status = NotificationHubPrivateLinkConnectionStatus.Approved,
        },
    },
};
ArmOperation<NotificationHubPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
NotificationHubPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NotificationHubPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
