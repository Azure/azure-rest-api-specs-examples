using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/vmClusters_create.json
// this example is just showing the usage of "CloudVmClusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg000";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CloudVmClusterResource
CloudVmClusterCollection collection = resourceGroupResource.GetCloudVmClusters();

// invoke the operation
string cloudvmclustername = "cluster1";
CloudVmClusterData data = new CloudVmClusterData(new AzureLocation("eastus"))
{
    Properties = new CloudVmClusterProperties(
    "hostname1",
    2,
    new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
    new string[] { "ssh-key 1" },
    new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
    "19.0.0.0",
    new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
    "cluster 1")
    {
        DataStorageSizeInTbs = 1000,
        DBNodeStorageSizeInGbs = 1000,
        MemorySizeInGbs = 1000,
        TimeZone = "UTC",
        Domain = "domain1",
        OcpuCount = 3,
        ClusterName = "cluster1",
        DataStoragePercentage = 100,
        IsLocalBackupEnabled = false,
        IsSparseDiskgroupEnabled = false,
        LicenseModel = OracleLicenseModel.LicenseIncluded,
        ScanListenerPortTcp = 1050,
        ScanListenerPortTcpSsl = 1025,
        BackupSubnetCidr = "172.17.5.0/24",
        NsgCidrs = {new CloudVmClusterNsgCidr("10.0.0.0/16")
        {
        DestinationPortRange = new CloudVmClusterPortRange(1520, 1522),
        }, new CloudVmClusterNsgCidr("10.10.0.0/24")},
        DataCollectionOptions = new DiagnosticCollectionConfig
        {
            IsDiagnosticsEventsEnabled = false,
            IsHealthMonitoringEnabled = false,
            IsIncidentLogsEnabled = false,
        },
        DBServers = { new ResourceIdentifier("ocid1..aaaa") },
    },
    Tags =
    {
    ["tagK1"] = "tagV1"
    },
};
ArmOperation<CloudVmClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, cloudvmclustername, data);
CloudVmClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudVmClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
