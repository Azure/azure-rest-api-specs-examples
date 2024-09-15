using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolCreate_VirtualMachineConfiguration_ManagedOSDisk.json
// this example is just showing the usage of "Pool_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchAccountResource created on azure
// for more information of creating BatchAccountResource, please refer to the document of BatchAccountResource
string subscriptionId = "subid";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
ResourceIdentifier batchAccountResourceId = BatchAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
BatchAccountResource batchAccount = client.GetBatchAccountResource(batchAccountResourceId);

// get the collection of this BatchAccountPoolResource
BatchAccountPoolCollection collection = batchAccount.GetBatchAccountPools();

// invoke the operation
string poolName = "testpool";
BatchAccountPoolData data = new BatchAccountPoolData()
{
    VmSize = "Standard_d2s_v3",
    DeploymentVmConfiguration = new BatchVmConfiguration(new BatchImageReference()
    {
        Publisher = "microsoftwindowsserver",
        Offer = "windowsserver",
        Sku = "2022-datacenter-smalldisk",
    }, "batch.node.windows amd64")
    {
        OSDisk = new BatchOSDisk()
        {
            Caching = BatchDiskCachingType.ReadWrite,
            ManagedDisk = new ManagedDisk()
            {
                StorageAccountType = BatchStorageAccountType.StandardSsdLrs,
            },
            DiskSizeGB = 100,
            IsWriteAcceleratorEnabled = false,
        },
    },
    ScaleSettings = new BatchAccountPoolScaleSettings()
    {
        FixedScale = new BatchAccountFixedScaleSettings()
        {
            TargetDedicatedNodes = 1,
            TargetLowPriorityNodes = 0,
        },
    },
};
ArmOperation<BatchAccountPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, poolName, data);
BatchAccountPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BatchAccountPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
