using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceNotificationRecipientEmail.json
// this example is just showing the usage of "WorkspaceNotificationRecipientEmail_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceNotificationResource created on azure
// for more information of creating ServiceWorkspaceNotificationResource, please refer to the document of ServiceWorkspaceNotificationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
NotificationName notificationName = NotificationName.RequestPublisherNotificationMessage;
ResourceIdentifier serviceWorkspaceNotificationResourceId = ServiceWorkspaceNotificationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, notificationName);
ServiceWorkspaceNotificationResource serviceWorkspaceNotification = client.GetServiceWorkspaceNotificationResource(serviceWorkspaceNotificationResourceId);

// invoke the operation
string email = "foobar@live.com";
RecipientEmailContract result = await serviceWorkspaceNotification.CreateOrUpdateWorkspaceNotificationRecipientEmailAsync(email);

Console.WriteLine($"Succeeded: {result}");
