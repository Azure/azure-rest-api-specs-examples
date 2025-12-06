using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/NameSpaces/SBNameSpaceCreate.json
// this example is just showing the usage of "Namespaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ServiceBusNamespaceResource
ServiceBusNamespaceCollection collection = resourceGroupResource.GetServiceBusNamespaces();

// invoke the operation
string namespaceName = "sdk-Namespace2924";
ServiceBusNamespaceData data = new ServiceBusNamespaceData(new AzureLocation("South Central US"))
{
    Sku = new ServiceBusSku(ServiceBusSkuName.Premium)
    {
        Tier = ServiceBusSkuTier.Premium,
        Capacity = 4,
    },
    PremiumMessagingPartitions = 2,
    GeoDataReplication = new GeoDataReplicationProperties
    {
        MaxReplicationLagDurationInSeconds = 300,
        Locations = {new ServiceBusNamespaceReplicaLocation
        {
        LocationName = "eastus",
        RoleType = GeoDRRoleType.Primary,
        }, new ServiceBusNamespaceReplicaLocation
        {
        LocationName = "southcentralus",
        RoleType = GeoDRRoleType.Secondary,
        }},
    },
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ArmOperation<ServiceBusNamespaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, namespaceName, data);
ServiceBusNamespaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
