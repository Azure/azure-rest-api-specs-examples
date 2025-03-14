using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HDInsight.Containers.Models;
using Azure.ResourceManager.HDInsight.Containers;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/UninstallExistingClusterLibraries.json
// this example is just showing the usage of "ClusterLibraries_ManageLibraries" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightClusterResource created on azure
// for more information of creating HDInsightClusterResource, please refer to the document of HDInsightClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "hiloResourceGroup";
string clusterPoolName = "clusterPool";
string clusterName = "cluster";
ResourceIdentifier hdInsightClusterResourceId = HDInsightClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterPoolName, clusterName);
HDInsightClusterResource hdInsightCluster = client.GetHDInsightClusterResource(hdInsightClusterResourceId);

// invoke the operation
ClusterLibraryManagementOperationContent content = new ClusterLibraryManagementOperationContent(new ClusterLibraryManagementOperationProperties(LibraryManagementAction.Uninstall, new ClusterLibrary[]
{
new ClusterLibrary(new ClusterPyPILibraryProperties("tensorflow")),
new ClusterLibrary(new ClusterMavenLibraryProperties("org.apache.flink", "flink-connector-hudi"))
}));
await hdInsightCluster.ManageLibrariesClusterLibraryAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
