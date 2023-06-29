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

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateHDInsightClusterWithAvailabilityZones.json
// this example is just showing the usage of "Clusters_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HDInsightClusterResource
HDInsightClusterCollection collection = resourceGroupResource.GetHDInsightClusters();

// invoke the operation
string clusterName = "cluster1";
HDInsightClusterCreateOrUpdateContent content = new HDInsightClusterCreateOrUpdateContent()
{
    Zones =
    {
    "1"
    },
    Properties = new HDInsightClusterCreateOrUpdateProperties()
    {
        ClusterVersion = "3.6",
        OSType = HDInsightOSType.Linux,
        ClusterDefinition = new HDInsightClusterDefinition()
        {
            Kind = "hadoop",
            Configurations = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
            {
                ["ambari-conf"] = new Dictionary<string, object>()
                {
                    ["database-name"] = "{ambari database name}",
                    ["database-server"] = "{sql server name}.database.windows.net",
                    ["database-user-name"] = "**********",
                    ["database-user-password"] = "**********"
                },
                ["gateway"] = new Dictionary<string, object>()
                {
                    ["restAuthCredential.isEnabled"] = "true",
                    ["restAuthCredential.password"] = "**********",
                    ["restAuthCredential.username"] = "admin"
                },
                ["hive-env"] = new Dictionary<string, object>()
                {
                    ["hive_database"] = "Existing MSSQL Server database with SQL authentication",
                    ["hive_database_name"] = "{hive metastore name}",
                    ["hive_database_type"] = "mssql",
                    ["hive_existing_mssql_server_database"] = "{hive metastore name}",
                    ["hive_existing_mssql_server_host"] = "{sql server name}.database.windows.net",
                    ["hive_hostname"] = "{sql server name}.database.windows.net"
                },
                ["hive-site"] = new Dictionary<string, object>()
                {
                    ["javax.jdo.option.ConnectionDriverName"] = "com.microsoft.sqlserver.jdbc.SQLServerDriver",
                    ["javax.jdo.option.ConnectionPassword"] = "**********!",
                    ["javax.jdo.option.ConnectionURL"] = "jdbc:sqlserver://{sql server name}.database.windows.net;database={hive metastore name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0",
                    ["javax.jdo.option.ConnectionUserName"] = "**********"
                },
                ["oozie-env"] = new Dictionary<string, object>()
                {
                    ["oozie_database"] = "Existing MSSQL Server database with SQL authentication",
                    ["oozie_database_name"] = "{oozie metastore name}",
                    ["oozie_database_type"] = "mssql",
                    ["oozie_existing_mssql_server_database"] = "{oozie metastore name}",
                    ["oozie_existing_mssql_server_host"] = "{sql server name}.database.windows.net",
                    ["oozie_hostname"] = "{sql server name}.database.windows.net"
                },
                ["oozie-site"] = new Dictionary<string, object>()
                {
                    ["oozie.db.schema.name"] = "oozie",
                    ["oozie.service.JPAService.jdbc.driver"] = "com.microsoft.sqlserver.jdbc.SQLServerDriver",
                    ["oozie.service.JPAService.jdbc.password"] = "**********",
                    ["oozie.service.JPAService.jdbc.url"] = "jdbc:sqlserver://{sql server name}.database.windows.net;database={oozie metastore name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0",
                    ["oozie.service.JPAService.jdbc.username"] = "**********"
                }
            }),
        },
        ComputeRoles =
        {
        new HDInsightClusterRole()
        {
        Name = "headnode",
        TargetInstanceCount = 2,
        HardwareVmSize = "standard_d3",
        OSLinuxProfile = new HDInsightLinuxOSProfile()
        {
        Username = "sshuser",
        Password = "**********",
        SshPublicKeys =
        {
        new HDInsightSshPublicKey()
        {
        CertificateData = "**********",
        }
        },
        },
        VirtualNetworkProfile = new HDInsightVirtualNetworkProfile()
        {
        Id = new ResourceIdentifier("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
        Subnet = "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
        },
        },new HDInsightClusterRole()
        {
        Name = "workernode",
        TargetInstanceCount = 2,
        HardwareVmSize = "standard_d3",
        OSLinuxProfile = new HDInsightLinuxOSProfile()
        {
        Username = "sshuser",
        Password = "**********",
        SshPublicKeys =
        {
        new HDInsightSshPublicKey()
        {
        CertificateData = "**********",
        }
        },
        },
        VirtualNetworkProfile = new HDInsightVirtualNetworkProfile()
        {
        Id = new ResourceIdentifier("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
        Subnet = "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
        },
        }
        },
        StorageAccounts =
        {
        new HDInsightStorageAccountInfo()
        {
        Name = "mystorage",
        IsDefault = true,
        Container = "containername",
        Key = "storage account key",
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
