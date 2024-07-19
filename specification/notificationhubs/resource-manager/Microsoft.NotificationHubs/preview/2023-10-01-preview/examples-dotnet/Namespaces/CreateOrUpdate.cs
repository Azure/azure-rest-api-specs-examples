using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NotificationHubs.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NotificationHubs;

// Generated from example definition: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/Namespaces/CreateOrUpdate.json
// this example is just showing the usage of "Namespaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
string resourceGroupName = "5ktrial";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NotificationHubNamespaceResource
NotificationHubNamespaceCollection collection = resourceGroupResource.GetNotificationHubNamespaces();

// invoke the operation
string namespaceName = "nh-sdk-ns";
NotificationHubNamespaceData data = new NotificationHubNamespaceData(new AzureLocation("South Central US"), new NotificationHubSku(NotificationHubSkuName.Standard)
{
    Tier = "Standard",
})
{
    ZoneRedundancy = ZoneRedundancyPreference.Enabled,
    NetworkAcls = new NotificationHubNetworkAcls()
    {
        IPRules =
        {
        new NotificationHubIPRule("185.48.100.00/24",new AuthorizationRuleAccessRightExt[]
        {
        AuthorizationRuleAccessRightExt.Manage,AuthorizationRuleAccessRightExt.Send,AuthorizationRuleAccessRightExt.Listen
        })
        },
        PublicNetworkRuleAccessRights =
        {
        AuthorizationRuleAccessRightExt.Listen
        },
    },
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
};
ArmOperation<NotificationHubNamespaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, namespaceName, data);
NotificationHubNamespaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NotificationHubNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
