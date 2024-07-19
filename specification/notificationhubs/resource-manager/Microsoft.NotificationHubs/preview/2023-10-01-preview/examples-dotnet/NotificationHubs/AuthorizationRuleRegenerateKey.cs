using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NotificationHubs.Models;
using Azure.ResourceManager.NotificationHubs;

// Generated from example definition: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/NotificationHubs/AuthorizationRuleRegenerateKey.json
// this example is just showing the usage of "NotificationHubs_RegenerateKeys" operation, for the dependent resources, they will have to be created separately.

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
NotificationHubPolicyKey notificationHubPolicyKey = new NotificationHubPolicyKey("PrimaryKey");
NotificationHubResourceKeys result = await notificationHubAuthorizationRule.RegenerateKeysAsync(notificationHubPolicyKey);

Console.WriteLine($"Succeeded: {result}");
