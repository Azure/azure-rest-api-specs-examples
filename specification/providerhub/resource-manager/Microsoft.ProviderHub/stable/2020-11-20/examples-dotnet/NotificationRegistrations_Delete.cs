using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/NotificationRegistrations_Delete.json
// this example is just showing the usage of "NotificationRegistrations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NotificationRegistrationResource created on azure
// for more information of creating NotificationRegistrationResource, please refer to the document of NotificationRegistrationResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
string notificationRegistrationName = "fooNotificationRegistration";
ResourceIdentifier notificationRegistrationResourceId = NotificationRegistrationResource.CreateResourceIdentifier(subscriptionId, providerNamespace, notificationRegistrationName);
NotificationRegistrationResource notificationRegistration = client.GetNotificationRegistrationResource(notificationRegistrationResourceId);

// invoke the operation
await notificationRegistration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
