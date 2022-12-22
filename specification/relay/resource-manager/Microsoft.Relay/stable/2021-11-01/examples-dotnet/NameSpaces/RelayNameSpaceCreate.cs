using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/NameSpaces/RelayNameSpaceCreate.json
// this example is just showing the usage of "Namespaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "resourcegroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this RelayNamespaceResource
RelayNamespaceCollection collection = resourceGroupResource.GetRelayNamespaces();

// invoke the operation
string namespaceName = "example-RelayNamespace-5849";
RelayNamespaceData data = new RelayNamespaceData(new AzureLocation("South Central US"))
{
    Sku = new RelaySku(RelaySkuName.Standard)
    {
        Tier = RelaySkuTier.Standard,
    },
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
};
ArmOperation<RelayNamespaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, namespaceName, data);
RelayNamespaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RelayNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
