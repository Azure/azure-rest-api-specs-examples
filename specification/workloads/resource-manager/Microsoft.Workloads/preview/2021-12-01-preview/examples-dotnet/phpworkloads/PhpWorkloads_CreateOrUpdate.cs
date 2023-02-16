using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/PhpWorkloads_CreateOrUpdate.json
// this example is just showing the usage of "PhpWorkloads_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
string resourceGroupName = "test-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PhpWorkloadResource
PhpWorkloadResourceCollection collection = resourceGroupResource.GetPhpWorkloadResources();

// invoke the operation
string phpWorkloadName = "wp39";
PhpWorkloadResourceData data = new PhpWorkloadResourceData(new AzureLocation("eastus2"), WorkloadKind.WordPress)
{
    Sku = new WorkloadsSku("Large"),
    AppLocation = new AzureLocation("eastus"),
    ManagedResourceGroupName = "php-mrg-wp39",
    AdminUserProfile = new UserProfile("wpadmin", "===SSH=PUBLIC=KEY==="),
    WebNodesProfile = new VmssNodesProfile("Standard_DS2_v2", new OSImageProfile()
    {
        Publisher = OSImagePublisher.Canonical,
        Offer = OSImageOffer.UbuntuServer,
        Sku = new OSImageSku("18.0-LTS"),
        Version = OSImageVersion.Latest,
    }, new DiskInfo(DiskStorageType.PremiumLrs))
    {
        AutoScaleMinCount = 1,
        AutoScaleMaxCount = 1,
        Name = "web-server",
    },
    ControllerProfile = new NodeProfile("Standard_DS2_v2", new OSImageProfile()
    {
        Publisher = OSImagePublisher.Canonical,
        Offer = OSImageOffer.UbuntuServer,
        Sku = new OSImageSku("18.0-LTS"),
        Version = OSImageVersion.Latest,
    }, new DiskInfo(DiskStorageType.PremiumLrs))
    {
        Name = "contoller-vm",
        DataDisks =
        {
        new DiskInfo(DiskStorageType.PremiumLrs)
        {
        SizeInGB = 100,
        }
        },
    },
    NetworkProfile = new NetworkProfile(LoadBalancerType.LoadBalancer)
    {
        LoadBalancerSku = "Standard",
        AzureFrontDoorEnabled = AzureFrontDoorEnabled.Enabled,
    },
    DatabaseProfile = new DatabaseProfile(DatabaseType.MySql, "Standard_D32s_v4", DatabaseTier.GeneralPurpose)
    {
        ServerName = "wp-db-server",
        Version = "5.7",
        HAEnabled = HAEnabled.Disabled,
        StorageSku = "Premium_LRS",
        StorageInGB = 128,
        StorageIops = 200,
        BackupRetentionDays = 7,
        SslEnforcementEnabled = EnableSslEnforcement.Enabled,
    },
    SiteDomainName = "www.example.com",
    FileshareProfile = new FileshareProfile(FileShareType.AzureFiles, FileShareStorageType.PremiumLrs)
    {
        ShareSizeInGB = 100,
    },
    PhpVersion = PhpVersion.V7_3,
    SearchProfile = new SearchProfile("Standard_DS2_v2", new OSImageProfile()
    {
        Publisher = OSImagePublisher.Canonical,
        Offer = OSImageOffer.UbuntuServer,
        Sku = new OSImageSku("18.0-LTS"),
        Version = OSImageVersion.Latest,
    }, new DiskInfo(DiskStorageType.PremiumLrs), SearchType.Elastic),
    CacheProfile = new CacheProfile("Basic", RedisCacheFamily.C, 0)
    {
        Name = "wp-cache",
    },
    BackupProfile = new BackupProfile(EnableBackup.Disabled),
    Tags =
    {
    },
};
ArmOperation<PhpWorkloadResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, phpWorkloadName, data);
PhpWorkloadResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PhpWorkloadResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
