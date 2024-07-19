using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NotificationHubs.Models;
using Azure.ResourceManager.NotificationHubs;

// Generated from example definition: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/NotificationHubs/PnsCredentialsGet.json
// this example is just showing the usage of "NotificationHubs_GetPnsCredentials" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NotificationHubResource created on azure
// for more information of creating NotificationHubResource, please refer to the document of NotificationHubResource
string subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
string resourceGroupName = "5ktrial";
string namespaceName = "nh-sdk-ns";
string notificationHubName = "nh-sdk-hub";
ResourceIdentifier notificationHubResourceId = NotificationHubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, notificationHubName);
NotificationHubResource notificationHub = client.GetNotificationHubResource(notificationHubResourceId);

// invoke the operation
NotificationHubPnsCredentials result = await notificationHub.GetPnsCredentialsAsync();

Console.WriteLine($"Succeeded: {result}");
