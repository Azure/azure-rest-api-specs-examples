using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Kusto;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoManagedPrivateEndpointsCreateOrUpdate.json
// this example is just showing the usage of "ManagedPrivateEndpoints_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoClusterResource created on azure
// for more information of creating KustoClusterResource, please refer to the document of KustoClusterResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
ResourceIdentifier kustoClusterResourceId = KustoClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
KustoClusterResource kustoCluster = client.GetKustoClusterResource(kustoClusterResourceId);

// get the collection of this KustoManagedPrivateEndpointResource
KustoManagedPrivateEndpointCollection collection = kustoCluster.GetKustoManagedPrivateEndpoints();

// invoke the operation
string managedPrivateEndpointName = "managedPrivateEndpointTest";
KustoManagedPrivateEndpointData data = new KustoManagedPrivateEndpointData
{
    PrivateLinkResourceId = new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/storageAccountTest"),
    GroupId = "blob",
    RequestMessage = "Please Approve.",
};
ArmOperation<KustoManagedPrivateEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, managedPrivateEndpointName, data);
KustoManagedPrivateEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KustoManagedPrivateEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
