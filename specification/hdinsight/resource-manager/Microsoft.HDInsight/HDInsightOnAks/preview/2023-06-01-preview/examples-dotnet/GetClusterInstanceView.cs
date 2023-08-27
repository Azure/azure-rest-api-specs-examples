using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HDInsight.Containers;
using Azure.ResourceManager.HDInsight.Containers.Models;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/GetClusterInstanceView.json
// this example is just showing the usage of "Clusters_GetInstanceView" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightClusterResource created on azure
// for more information of creating HDInsightClusterResource, please refer to the document of HDInsightClusterResource
string subscriptionId = "10e32bab-26da-4cc4-a441-52b318f824e6";
string resourceGroupName = "rg1";
string clusterPoolName = "clusterPool1";
string clusterName = "cluster1";
ResourceIdentifier hdInsightClusterResourceId = HDInsightClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterPoolName, clusterName);
HDInsightClusterResource hdInsightCluster = client.GetHDInsightClusterResource(hdInsightClusterResourceId);

// invoke the operation
ClusterInstanceViewResult result = await hdInsightCluster.GetInstanceViewAsync();

Console.WriteLine($"Succeeded: {result}");
