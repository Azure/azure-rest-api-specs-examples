using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-11-11/examples/KustoPrivateLinkResourcesGet.json
// this example is just showing the usage of "PrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this KustoPrivateLinkResource
KustoPrivateLinkResourceCollection collection = kustoCluster.GetKustoPrivateLinkResources();

// invoke the operation
string privateLinkResourceName = "cluster";
bool result = await collection.ExistsAsync(privateLinkResourceName);

Console.WriteLine($"Succeeded: {result}");
