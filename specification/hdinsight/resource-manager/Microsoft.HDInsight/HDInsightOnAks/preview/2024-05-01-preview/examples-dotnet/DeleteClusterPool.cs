using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HDInsight.Containers.Models;
using Azure.ResourceManager.HDInsight.Containers;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/DeleteClusterPool.json
// this example is just showing the usage of "ClusterPools_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightClusterPoolResource created on azure
// for more information of creating HDInsightClusterPoolResource, please refer to the document of HDInsightClusterPoolResource
string subscriptionId = "10e32bab-26da-4cc4-a441-52b318f824e6";
string resourceGroupName = "rg1";
string clusterPoolName = "clusterpool1";
ResourceIdentifier hdInsightClusterPoolResourceId = HDInsightClusterPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterPoolName);
HDInsightClusterPoolResource hdInsightClusterPool = client.GetHDInsightClusterPoolResource(hdInsightClusterPoolResourceId);

// invoke the operation
await hdInsightClusterPool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
