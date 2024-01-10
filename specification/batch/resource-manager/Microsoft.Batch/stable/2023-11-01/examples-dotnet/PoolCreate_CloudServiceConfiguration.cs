using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolCreate_CloudServiceConfiguration.json
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
    DisplayName = "my-pool-name",
    VmSize = "STANDARD_D4",
    DeploymentConfiguration = new BatchDeploymentConfiguration()
    {
        CloudServiceConfiguration = new BatchCloudServiceConfiguration("4")
        {
            OSVersion = "WA-GUEST-OS-4.45_201708-01",
        },
    },
    ScaleSettings = new BatchAccountPoolScaleSettings()
    {
        FixedScale = new BatchAccountFixedScaleSettings()
        {
            ResizeTimeout = XmlConvert.ToTimeSpan("PT8M"),
            TargetDedicatedNodes = 6,
            TargetLowPriorityNodes = 28,
            NodeDeallocationOption = BatchNodeDeallocationOption.TaskCompletion,
        },
    },
    InterNodeCommunication = InterNodeCommunicationState.Enabled,
    NetworkConfiguration = new BatchNetworkConfiguration()
    {
        SubnetId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
        PublicIPAddressConfiguration = new BatchPublicIPAddressConfiguration()
        {
            Provision = BatchIPAddressProvisioningType.UserManaged,
            IPAddressIds =
            {
            new ResourceIdentifier("/subscriptions/subid1/resourceGroups/rg13/providers/Microsoft.Network/publicIPAddresses/ip135"),new ResourceIdentifier("/subscriptions/subid2/resourceGroups/rg24/providers/Microsoft.Network/publicIPAddresses/ip268")
            },
        },
    },
    TaskSlotsPerNode = 13,
    TaskSchedulingNodeFillType = BatchNodeFillType.Pack,
    UserAccounts =
    {
    new BatchUserAccount("username1","<ExamplePassword>")
    {
    ElevationLevel = BatchUserAccountElevationLevel.Admin,
    LinuxUserConfiguration = new BatchLinuxUserConfiguration()
    {
    Uid = 1234,
    Gid = 4567,
    SshPrivateKey = "sshprivatekeyvalue",
    },
    }
    },
    Metadata =
    {
    new BatchAccountPoolMetadataItem("metadata-1","value-1"),new BatchAccountPoolMetadataItem("metadata-2","value-2")
    },
    StartTask = new BatchAccountPoolStartTask()
    {
        CommandLine = "cmd /c SET",
        ResourceFiles =
        {
        new BatchResourceFile()
        {
        HttpUri = new Uri("https://testaccount.blob.core.windows.net/example-blob-file"),
        FilePath = "c:\\temp\\gohere",
        FileMode = "777",
        }
        },
        EnvironmentSettings =
        {
        new BatchEnvironmentSetting("MYSET")
        {
        Value = "1234",
        }
        },
        UserIdentity = new BatchUserIdentity()
        {
            AutoUser = new BatchAutoUserSpecification()
            {
                Scope = BatchAutoUserScope.Pool,
                ElevationLevel = BatchUserAccountElevationLevel.Admin,
            },
        },
        MaxTaskRetryCount = 6,
        WaitForSuccess = true,
    },
    Certificates =
    {
    new BatchCertificateReference(new ResourceIdentifier("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"))
    {
    StoreLocation = BatchCertificateStoreLocation.LocalMachine,
    StoreName = "MY",
    Visibility =
    {
    BatchCertificateVisibility.RemoteUser
    },
    }
    },
    ApplicationPackages =
    {
    new BatchApplicationPackageReference(new ResourceIdentifier("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"))
    {
    Version = "asdf",
    }
    },
    ApplicationLicenses =
    {
    "app-license0","app-license1"
    },
};
ArmOperation<BatchAccountPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, poolName, data);
BatchAccountPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BatchAccountPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
