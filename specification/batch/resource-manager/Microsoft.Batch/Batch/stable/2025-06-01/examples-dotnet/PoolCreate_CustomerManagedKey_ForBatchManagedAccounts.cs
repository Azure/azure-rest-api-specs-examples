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

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/PoolCreate_CustomerManagedKey_ForBatchManagedAccounts.json
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
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity()
        },
    },
    VmSize = "Standard_D4ds_v5",
    DeploymentVmConfiguration = new BatchVmConfiguration(new BatchImageReference
    {
        Publisher = "MicrosoftWindowsServer",
        Offer = "WindowsServer",
        Sku = "2022-Datacenter",
        Version = "latest",
    }, "batch.node.windows amd64")
    {
        DiskEncryptionConfiguration = new BatchDiskEncryptionConfiguration
        {
            Targets = { BatchDiskEncryptionTarget.OSDisk },
            CustomerManagedKey = new BatchDiskCustomerManagedKey
            {
                KeyUri = new Uri("http://sample.vault.azure.net//keys/cmk/bb60031a6d4545d3a60d3f94588538c9"),
                IdentityReferenceResourceId = new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"),
            },
        },
    },
    ScaleSettings = new BatchAccountPoolScaleSettings
    {
        FixedScale = new BatchAccountFixedScaleSettings
        {
            ResizeTimeout = XmlConvert.ToTimeSpan("PT15M"),
            TargetDedicatedNodes = 1,
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
