using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HDInsight.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.HDInsight;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/EnableOrUpdateAutoScaleWithLoadBasedConfiguration.json
// this example is just showing the usage of "Clusters_UpdateAutoScaleConfiguration" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightClusterResource created on azure
// for more information of creating HDInsightClusterResource, please refer to the document of HDInsightClusterResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string clusterName = "cluster1";
ResourceIdentifier hdInsightClusterResourceId = HDInsightClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
HDInsightClusterResource hdInsightCluster = client.GetHDInsightClusterResource(hdInsightClusterResourceId);

// invoke the operation
HDInsightRoleName roleName = HDInsightRoleName.Workernode;
HDInsightAutoScaleConfigurationUpdateContent content = new HDInsightAutoScaleConfigurationUpdateContent
{
    AutoScale = new HDInsightAutoScaleConfiguration
    {
        Capacity = new HDInsightAutoScaleCapacity
        {
            MinInstanceCount = 3,
            MaxInstanceCount = 5,
        },
    },
};
await hdInsightCluster.UpdateAutoScaleConfigurationAsync(WaitUntil.Completed, roleName, content);

Console.WriteLine("Succeeded");
