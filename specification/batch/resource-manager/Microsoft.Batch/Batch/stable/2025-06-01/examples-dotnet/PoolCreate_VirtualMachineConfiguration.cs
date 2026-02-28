using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/PoolCreate_VirtualMachineConfiguration.json
// this example is just showing the usage of "Pool_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchAccountResource created on azure
// for more information of creating BatchAccountResource, please refer to the document of BatchAccountResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
ResourceIdentifier batchAccountResourceId = BatchAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
BatchAccountResource batchAccount = client.GetBatchAccountResource(batchAccountResourceId);

// get the collection of this BatchAccountPoolResource
BatchAccountPoolCollection collection = batchAccount.GetBatchAccountPools();

// invoke the operation
string poolName = "testpool";
BatchAccountPoolData data = new BatchAccountPoolData
{
    VmSize = "STANDARD_D4",
    DeploymentVmConfiguration = new BatchVmConfiguration(new BatchImageReference
    {
        Publisher = "MicrosoftWindowsServer",
        Offer = "WindowsServer",
        Sku = "2016-Datacenter-SmallDisk",
        Version = "latest",
    }, "batch.node.windows amd64")
    {
        IsAutomaticUpdateEnabled = false,
        DataDisks = {new BatchVmDataDisk(0, 30)
        {
        Caching = BatchDiskCachingType.ReadWrite,
        ManagedDisk = new ManagedDisk
        {
        StorageAccountType = BatchStorageAccountType.StandardSsdLrs,
        },
        }, new BatchVmDataDisk(1, 200)
        {
        Caching = BatchDiskCachingType.None,
        ManagedDisk = new ManagedDisk
        {
        StorageAccountType = BatchStorageAccountType.PremiumLrs,
        },
        }},
        LicenseType = "Windows_Server",
        DiskEncryptionConfiguration = new BatchDiskEncryptionConfiguration
        {
            Targets = { BatchDiskEncryptionTarget.OSDisk, BatchDiskEncryptionTarget.TemporaryDisk },
        },
        NodePlacementPolicy = BatchNodePlacementPolicyType.Zonal,
        OSDisk = new BatchOSDisk
        {
            EphemeralOSDiskPlacement = BatchDiffDiskPlacement.CacheDisk,
        },
    },
    ScaleSettings = new BatchAccountPoolScaleSettings
    {
        AutoScale = new BatchAccountAutoScaleSettings("$TargetDedicatedNodes=1")
        {
            EvaluationInterval = XmlConvert.ToTimeSpan("PT5M"),
        },
    },
    NetworkConfiguration = new BatchNetworkConfiguration
    {
        EndpointInboundNatPools = {new BatchInboundNatPool("testnat", BatchInboundEndpointProtocol.Tcp, 12001, 15000, 15100)
        {
        NetworkSecurityGroupRules = {new BatchNetworkSecurityGroupRule(150, BatchNetworkSecurityGroupRuleAccess.Allow, "192.100.12.45")
        {
        SourcePortRanges = {"1", "2"},
        }, new BatchNetworkSecurityGroupRule(3500, BatchNetworkSecurityGroupRuleAccess.Deny, "*")
        {
        SourcePortRanges = {"*"},
        }},
        }},
    },
};
ArmOperation<BatchAccountPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, poolName, data);
BatchAccountPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BatchAccountPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
