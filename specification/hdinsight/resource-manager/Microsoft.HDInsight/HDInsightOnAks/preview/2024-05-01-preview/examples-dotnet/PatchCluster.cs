using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HDInsight.Containers.Models;
using Azure.ResourceManager.HDInsight.Containers;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/PatchCluster.json
// this example is just showing the usage of "Clusters_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightClusterResource created on azure
// for more information of creating HDInsightClusterResource, please refer to the document of HDInsightClusterResource
string subscriptionId = "10e32bab-26da-4cc4-a441-52b318f824e6";
string resourceGroupName = "hiloResourcegroup";
string clusterPoolName = "clusterpool1";
string clusterName = "cluster1";
ResourceIdentifier hdInsightClusterResourceId = HDInsightClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterPoolName, clusterName);
HDInsightClusterResource hdInsightCluster = client.GetHDInsightClusterResource(hdInsightClusterResourceId);

// invoke the operation
HDInsightClusterPatch patch = new HDInsightClusterPatch()
{
    ClusterProfile = new UpdatableClusterProfile()
    {
        ServiceConfigsProfiles =
        {
        new ClusterServiceConfigsProfile("TestService1",new ClusterServiceConfig[]
        {
        new ClusterServiceConfig("TestComp1",new ClusterConfigFile[]
        {
        new ClusterConfigFile("TestFile1")
        {
        Values =
        {
        ["Test.config.1"] = "1",
        ["Test.config.2"] = "2",
        },
        },new ClusterConfigFile("TestFile2")
        {
        Values =
        {
        ["Test.config.3"] = "3",
        ["Test.config.4"] = "4",
        },
        }
        }),new ClusterServiceConfig("TestComp2",new ClusterConfigFile[]
        {
        new ClusterConfigFile("TestFile3")
        {
        Content = "TestContent",
        Path = "TestPath",
        },new ClusterConfigFile("TestFile4")
        {
        Values =
        {
        ["Test.config.7"] = "7",
        ["Test.config.8"] = "8",
        },
        }
        })
        }),new ClusterServiceConfigsProfile("TestService2",new ClusterServiceConfig[]
        {
        new ClusterServiceConfig("TestComp3",new ClusterConfigFile[]
        {
        new ClusterConfigFile("TestFile5")
        {
        Values =
        {
        ["Test.config.9"] = "9",
        },
        }
        })
        })
        },
        SshProfile = new ClusterSshProfile(2),
        AutoscaleProfile = new ClusterAutoscaleProfile(true)
        {
            GracefulDecommissionTimeout = -1,
            AutoscaleType = ClusterAutoscaleType.ScheduleBased,
            ScheduleBasedConfig = new ScheduleBasedConfig("Cen. Australia Standard Time", 3, new AutoscaleSchedule[]
{
new AutoscaleSchedule("00:00","12:00",3,new AutoscaleScheduleDay[]
{
new AutoscaleScheduleDay("Monday, Tuesday, Wednesday")
}),new AutoscaleSchedule("00:00","12:00",3,new AutoscaleScheduleDay[]
{
AutoscaleScheduleDay.Sunday
})
}),
        },
        AuthorizationProfile = new AuthorizationProfile()
        {
            UserIds =
            {
            "Testuser1","Testuser2"
            },
        },
        LogAnalyticsProfile = new ClusterLogAnalyticsProfile(true)
        {
            ApplicationLogs = new ClusterLogAnalyticsApplicationLogs()
            {
                IsStdOutEnabled = true,
                IsStdErrorEnabled = true,
            },
            IsMetricsEnabled = true,
        },
    },
};
ArmOperation<HDInsightClusterResource> lro = await hdInsightCluster.UpdateAsync(WaitUntil.Completed, patch);
HDInsightClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HDInsightClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
