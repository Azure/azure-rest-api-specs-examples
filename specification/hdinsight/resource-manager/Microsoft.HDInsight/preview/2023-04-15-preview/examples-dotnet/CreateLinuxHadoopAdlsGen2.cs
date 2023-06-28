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

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateLinuxHadoopAdlsGen2.json
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
    Tags =
    {
    ["key1"] = "val1",
    },
    Properties = new HDInsightClusterCreateOrUpdateProperties()
    {
        ClusterVersion = "3.6",
        OSType = HDInsightOSType.Linux,
        Tier = HDInsightTier.Standard,
        ClusterDefinition = new HDInsightClusterDefinition()
        {
            Kind = "Hadoop",
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
        Name = "headnode",
        MinInstanceCount = 1,
        TargetInstanceCount = 2,
        HardwareVmSize = "Standard_D3_V2",
        OSLinuxProfile = new HDInsightLinuxOSProfile()
        {
        Username = "sshuser",
        Password = "**********",
        },
        },new HDInsightClusterRole()
        {
        Name = "workernode",
        MinInstanceCount = 1,
        TargetInstanceCount = 4,
        HardwareVmSize = "Standard_D3_V2",
        OSLinuxProfile = new HDInsightLinuxOSProfile()
        {
        Username = "sshuser",
        Password = "**********",
        },
        },new HDInsightClusterRole()
        {
        Name = "zookeepernode",
        MinInstanceCount = 1,
        TargetInstanceCount = 3,
        HardwareVmSize = "Small",
        OSLinuxProfile = new HDInsightLinuxOSProfile()
        {
        Username = "sshuser",
        Password = "**********",
        },
        }
        },
        StorageAccounts =
        {
        new HDInsightStorageAccountInfo()
        {
        Name = "mystorage.dfs.core.windows.net",
        IsDefault = true,
        FileSystem = "default",
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
