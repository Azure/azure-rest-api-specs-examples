using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoClustersCreateOrUpdate.json
// this example is just showing the usage of "Clusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this KustoClusterResource
KustoClusterCollection collection = resourceGroupResource.GetKustoClusters();

// invoke the operation
string clusterName = "kustoCluster";
KustoClusterData data = new KustoClusterData(new AzureLocation("westus"), new KustoSku(KustoSkuName.StandardL16asV3, KustoSkuTier.Standard)
{
    Capacity = 2,
})
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    IsStreamingIngestEnabled = true,
    IsPurgeEnabled = true,
    IsDoubleEncryptionEnabled = false,
    PublicNetworkAccess = KustoClusterPublicNetworkAccess.Enabled,
    AllowedIPRangeList =
    {
    "0.0.0.0/0"
    },
    IsAutoStopEnabled = true,
    PublicIPType = KustoClusterPublicIPType.DualStack,
};
ArmOperation<KustoClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
KustoClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KustoClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
