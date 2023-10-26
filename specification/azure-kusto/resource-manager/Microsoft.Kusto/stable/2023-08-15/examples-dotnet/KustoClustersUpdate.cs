using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClustersUpdate.json
// this example is just showing the usage of "Clusters_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoClusterResource created on azure
// for more information of creating KustoClusterResource, please refer to the document of KustoClusterResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster2";
ResourceIdentifier kustoClusterResourceId = KustoClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
KustoClusterResource kustoCluster = client.GetKustoClusterResource(kustoClusterResourceId);

// invoke the operation
KustoClusterPatch patch = new KustoClusterPatch(new AzureLocation("westus"));
string ifMatch = "*";
ArmOperation<KustoClusterResource> lro = await kustoCluster.UpdateAsync(WaitUntil.Completed, patch, ifMatch: ifMatch);
KustoClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KustoClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
