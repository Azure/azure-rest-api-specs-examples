using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/NotificationRegistrations_CreateOrUpdate.json
// this example is just showing the usage of "NotificationRegistrations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProviderRegistrationResource created on azure
// for more information of creating ProviderRegistrationResource, please refer to the document of ProviderRegistrationResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
ResourceIdentifier providerRegistrationResourceId = ProviderRegistrationResource.CreateResourceIdentifier(subscriptionId, providerNamespace);
ProviderRegistrationResource providerRegistration = client.GetProviderRegistrationResource(providerRegistrationResourceId);

// get the collection of this NotificationRegistrationResource
NotificationRegistrationCollection collection = providerRegistration.GetNotificationRegistrations();

// invoke the operation
string notificationRegistrationName = "fooNotificationRegistration";
NotificationRegistrationData data = new NotificationRegistrationData
{
    Properties = new NotificationRegistrationProperties
    {
        NotificationMode = NotificationMode.EventHub,
        MessageScope = MessageScope.RegisteredSubscriptions,
        IncludedEvents = { "*/write", "Microsoft.Contoso/employees/delete" },
        NotificationEndpoints = {new NotificationEndpoint
        {
        NotificationDestination = new ResourceIdentifier("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-eastus/providers/Microsoft.EventHub/namespaces/unitedstates-mgmtexpint/eventhubs/armlinkednotifications"),
        Locations = {new AzureLocation(""), new AzureLocation("East US")},
        }, new NotificationEndpoint
        {
        NotificationDestination = new ResourceIdentifier("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-northeurope/providers/Microsoft.EventHub/namespaces/europe-mgmtexpint/eventhubs/armlinkednotifications"),
        Locations = {new AzureLocation("North Europe")},
        }},
    },
};
ArmOperation<NotificationRegistrationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, notificationRegistrationName, data);
NotificationRegistrationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NotificationRegistrationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
