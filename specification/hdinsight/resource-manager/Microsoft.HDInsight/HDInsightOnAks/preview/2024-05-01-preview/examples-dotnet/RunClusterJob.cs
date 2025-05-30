using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HDInsight.Containers.Models;
using Azure.ResourceManager.HDInsight.Containers;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/RunClusterJob.json
// this example is just showing the usage of "ClusterJobs_RunJob" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightClusterResource created on azure
// for more information of creating HDInsightClusterResource, please refer to the document of HDInsightClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "hiloResourcegroup";
string clusterPoolName = "clusterpool1";
string clusterName = "cluster1";
ResourceIdentifier hdInsightClusterResourceId = HDInsightClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterPoolName, clusterName);
HDInsightClusterResource hdInsightCluster = client.GetHDInsightClusterResource(hdInsightClusterResourceId);

// invoke the operation
ClusterJob clusterJob = new ClusterJob(new FlinkJobProperties
{
    JobName = "flink-job-name",
    JobJarDirectory = "abfs://flinkjob@hilosa.dfs.core.windows.net/jars",
    JarName = "flink-sleep-job-0.0.1-SNAPSHOT.jar",
    EntryClass = "com.microsoft.hilo.flink.job.streaming.SleepJob",
    Action = FlinkJobAction.Start,
    FlinkConfiguration =
    {
    ["parallelism"] = "1",
    ["savepoint.directory"] = "abfs://flinkjob@hilosa.dfs.core.windows.net/savepoint"
    },
});
ArmOperation<ClusterJob> lro = await hdInsightCluster.RunJobClusterJobAsync(WaitUntil.Completed, clusterJob);
ClusterJob result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
