using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NotificationHubs;
using Azure.ResourceManager.NotificationHubs.Models;

// Generated from example definition: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceAuthorizationRuleGet.json
// this example is just showing the usage of "Namespaces_GetAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this NotificationHubNamespaceAuthorizationRuleResource
NotificationHubNamespaceAuthorizationRuleCollection collection = notificationHubNamespace.GetNotificationHubNamespaceAuthorizationRules();

// invoke the operation
string authorizationRuleName = "RootManageSharedAccessKey";
bool result = await collection.ExistsAsync(authorizationRuleName);

Console.WriteLine($"Succeeded: {result}");
