using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NotificationHubs.Models;
using Azure.ResourceManager.NotificationHubs;

// Generated from example definition: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/NotificationHubs/AuthorizationRuleCreateOrUpdate.json
// this example is just showing the usage of "NotificationHubs_CreateOrUpdateAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

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
string authorizationRuleName = "MyManageSharedAccessKey";
ResourceIdentifier notificationHubAuthorizationRuleResourceId = NotificationHubAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, notificationHubName, authorizationRuleName);
NotificationHubAuthorizationRuleResource notificationHubAuthorizationRule = client.GetNotificationHubAuthorizationRuleResource(notificationHubAuthorizationRuleResourceId);

// invoke the operation
NotificationHubAuthorizationRuleData data = new NotificationHubAuthorizationRuleData(default)
{
    AccessRights = { AuthorizationRuleAccessRightExt.Listen, AuthorizationRuleAccessRightExt.Send },
};
ArmOperation<NotificationHubAuthorizationRuleResource> lro = await notificationHubAuthorizationRule.UpdateAsync(WaitUntil.Completed, data);
NotificationHubAuthorizationRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NotificationHubAuthorizationRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
