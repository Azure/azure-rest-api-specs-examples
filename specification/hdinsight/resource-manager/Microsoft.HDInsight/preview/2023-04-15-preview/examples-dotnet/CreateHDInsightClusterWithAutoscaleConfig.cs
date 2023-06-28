using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HDInsight;
using Azure.ResourceManager.HDInsight.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateHDInsightClusterWithAutoscaleConfig.json
// this example is just showing the usage of "Clusters_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HDInsightClusterResource
HDInsightClusterCollection collection = resourceGroupResource.GetHDInsightClusters();

// invoke the operation
string clusterName = "cluster1";
HDInsightClusterCreateOrUpdateContent content = new HDInsightClusterCreateOrUpdateContent()
{
    Properties = new HDInsightClusterCreateOrUpdateProperties()
    {
        ClusterVersion = "3.6",
        OSType = HDInsightOSType.Linux,
        Tier = HDInsightTier.Standard,
        ClusterDefinition = new HDInsightClusterDefinition()
        {
            Kind = "hadoop",
            ComponentVersion =
            {
            ["Hadoop"] = "2.7",
            },
            Configurations = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
            {
                ["gateway"] = new Dictionary<string, object>()
                {
                    ["restAuthCredential.isEnabled"] = "true",
                    ["restAuthCredential.password"] = "**********",
                    ["restAuthCredential.username"] = "admin"
                }
            }),
        },
        ComputeRoles =
        {
        new HDInsightClusterRole()
        {
        Name = "workernode",
        TargetInstanceCount = 4,
        AutoScaleConfiguration = new HDInsightAutoScaleConfiguration()
        {
        Capacity = null,
        Recurrence = new HDInsightAutoScaleRecurrence()
        {
        TimeZone = "China Standard Time",
        Schedule =
        {
        new HDInsightAutoScaleSchedule()
        {
        Days =
        {
        HDInsightDayOfWeek.Monday,HDInsightDayOfWeek.Tuesday,HDInsightDayOfWeek.Wednesday,HDInsightDayOfWeek.Thursday,HDInsightDayOfWeek.Friday
        },
        TimeAndCapacity = new HDInsightAutoScaleTimeAndCapacity()
        {
        Time = "09:00",
        MinInstanceCount = 3,
        MaxInstanceCount = 3,
        },
        },new HDInsightAutoScaleSchedule()
        {
        Days =
        {
        HDInsightDayOfWeek.Monday,HDInsightDayOfWeek.Tuesday,HDInsightDayOfWeek.Wednesday,HDInsightDayOfWeek.Thursday,HDInsightDayOfWeek.Friday
        },
        TimeAndCapacity = new HDInsightAutoScaleTimeAndCapacity()
        {
        Time = "18:00",
        MinInstanceCount = 6,
        MaxInstanceCount = 6,
        },
        },new HDInsightAutoScaleSchedule()
        {
        Days =
        {
        HDInsightDayOfWeek.Saturday,HDInsightDayOfWeek.Sunday
        },
        TimeAndCapacity = new HDInsightAutoScaleTimeAndCapacity()
        {
        Time = "09:00",
        MinInstanceCount = 2,
        MaxInstanceCount = 2,
        },
        },new HDInsightAutoScaleSchedule()
        {
        Days =
        {
        HDInsightDayOfWeek.Saturday,HDInsightDayOfWeek.Sunday
        },
        TimeAndCapacity = new HDInsightAutoScaleTimeAndCapacity()
        {
        Time = "18:00",
        MinInstanceCount = 4,
        MaxInstanceCount = 4,
        },
        }
        },
        },
        },
        HardwareVmSize = "Standard_D4_V2",
        OSLinuxProfile = new HDInsightLinuxOSProfile()
        {
        Username = "sshuser",
        Password = "**********",
        },
        VirtualNetworkProfile = null,
        DataDisksGroups =
        {
        },
        ScriptActions =
        {
        },
        }
        },
        StorageAccounts =
        {
        new HDInsightStorageAccountInfo()
        {
        Name = "mystorage.blob.core.windows.net",
        IsDefault = true,
        Container = "hdinsight-autoscale-tes-2019-06-18t05-49-16-591z",
        Key = "storagekey",
        EnableSecureChannel = true,
        }
        },
    },
};
ArmOperation<HDInsightClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, content);
HDInsightClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HDInsightClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
