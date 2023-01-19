using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NotificationHubs;
using Azure.ResourceManager.NotificationHubs.Models;

// Generated from example definition: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleDelete.json
// this example is just showing the usage of "NotificationHubs_DeleteAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NotificationHubAuthorizationRuleResource created on azure
// for more information of creating NotificationHubAuthorizationRuleResource, please refer to the document of NotificationHubAuthorizationRuleResource
string subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
string resourceGroupName = "5ktrial";
string namespaceName = "nh-sdk-ns";
string notificationHubName = "nh-sdk-hub";
string authorizationRuleName = "DefaultListenSharedAccessSignature";
ResourceIdentifier notificationHubAuthorizationRuleResourceId = NotificationHubAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, notificationHubName, authorizationRuleName);
NotificationHubAuthorizationRuleResource notificationHubAuthorizationRule = client.GetNotificationHubAuthorizationRuleResource(notificationHubAuthorizationRuleResourceId);

// invoke the operation
await notificationHubAuthorizationRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
