using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/NameSpaces/EHNameSpaceCreate.json
// this example is just showing the usage of "Namespaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "SampleSubscription";
string resourceGroupName = "ResurceGroupSample";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this EventHubsNamespaceResource
EventHubsNamespaceCollection collection = resourceGroupResource.GetEventHubsNamespaces();

// invoke the operation
string namespaceName = "NamespaceSample";
EventHubsNamespaceData data = new EventHubsNamespaceData(new AzureLocation("East US"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned, UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1")] = new UserAssignedIdentity(),
        [new ResourceIdentifier("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2")] = new UserAssignedIdentity()
        },
    },
    ClusterArmId = new ResourceIdentifier("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.EventHub/clusters/enc-test"),
    Encryption = new EventHubsEncryption
    {
        KeyVaultProperties = {new EventHubsKeyVaultProperties
        {
        KeyName = "Samplekey",
        KeyVaultUri = new Uri("https://aprao-keyvault-user.vault-int.azure-int.net/"),
        UserAssignedIdentity = "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1",
        }},
        KeySource = EventHubsKeySource.MicrosoftKeyVault,
    },
    GeoDataReplication = new NamespaceGeoDataReplicationProperties
    {
        MaxReplicationLagDurationInSeconds = 300,
        Locations = {new EventHubsNamespaceReplicaLocation
        {
        LocationName = "eastus",
        RoleType = EventHubsNamespaceGeoDRRoleType.Primary,
        }, new EventHubsNamespaceReplicaLocation
        {
        LocationName = "southcentralus",
        RoleType = EventHubsNamespaceGeoDRRoleType.Secondary,
        }},
    },
};
ArmOperation<EventHubsNamespaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, namespaceName, data);
EventHubsNamespaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
