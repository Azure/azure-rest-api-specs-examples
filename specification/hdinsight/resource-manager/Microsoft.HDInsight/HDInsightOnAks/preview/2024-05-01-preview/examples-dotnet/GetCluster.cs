using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HDInsight.Containers.Models;
using Azure.ResourceManager.HDInsight.Containers;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/GetCluster.json
// this example is just showing the usage of "Clusters_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightClusterPoolResource created on azure
// for more information of creating HDInsightClusterPoolResource, please refer to the document of HDInsightClusterPoolResource
string subscriptionId = "10e32bab-26da-4cc4-a441-52b318f824e6";
string resourceGroupName = "hiloResourcegroup";
string clusterPoolName = "clusterpool1";
ResourceIdentifier hdInsightClusterPoolResourceId = HDInsightClusterPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterPoolName);
HDInsightClusterPoolResource hdInsightClusterPool = client.GetHDInsightClusterPoolResource(hdInsightClusterPoolResourceId);

// get the collection of this HDInsightClusterResource
HDInsightClusterCollection collection = hdInsightClusterPool.GetHDInsightClusters();

// invoke the operation
string clusterName = "cluster1";
NullableResponse<HDInsightClusterResource> response = await collection.GetIfExistsAsync(clusterName);
HDInsightClusterResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HDInsightClusterData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
