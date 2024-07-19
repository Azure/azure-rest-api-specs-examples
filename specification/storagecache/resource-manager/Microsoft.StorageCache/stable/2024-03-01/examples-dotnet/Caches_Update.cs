using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StorageCache.Models;
using Azure.ResourceManager.StorageCache;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/Caches_Update.json
// this example is just showing the usage of "Caches_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageCacheResource created on azure
// for more information of creating StorageCacheResource, please refer to the document of StorageCacheResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string cacheName = "sc1";
ResourceIdentifier storageCacheResourceId = StorageCacheResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName);
StorageCacheResource storageCache = client.GetStorageCacheResource(storageCacheResourceId);

// invoke the operation
StorageCacheData data = new StorageCacheData(new AzureLocation("westus"))
{
    SkuName = "Standard_2G",
    CacheSizeGB = 3072,
    Subnet = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/sub1"),
    UpgradeSettings = new StorageCacheUpgradeSettings()
    {
        EnableUpgradeSchedule = true,
        ScheduledOn = DateTimeOffset.Parse("2022-04-26T18:25:43.511Z"),
    },
    NetworkSettings = new StorageCacheNetworkSettings()
    {
        Mtu = 1500,
        DnsServers =
        {
        IPAddress.Parse("10.1.22.33"),IPAddress.Parse("10.1.12.33")
        },
        DnsSearchDomain = "contoso.com",
        NtpServer = "time.contoso.com",
    },
    SecurityAccessPolicies =
    {
    new NfsAccessPolicy("default",new NfsAccessRule[]
    {
    new NfsAccessRule(NfsAccessRuleScope.Default,NfsAccessRuleAccess.ReadWrite)
    {
    AllowSuid = false,
    AllowSubmountAccess = true,
    EnableRootSquash = false,
    }
    }),new NfsAccessPolicy("restrictive",new NfsAccessRule[]
    {
    new NfsAccessRule(NfsAccessRuleScope.Host,NfsAccessRuleAccess.ReadWrite)
    {
    Filter = "10.99.3.145",
    AllowSuid = true,
    AllowSubmountAccess = true,
    EnableRootSquash = false,
    },new NfsAccessRule(NfsAccessRuleScope.Network,NfsAccessRuleAccess.ReadWrite)
    {
    Filter = "10.99.1.0/24",
    AllowSuid = true,
    AllowSubmountAccess = true,
    EnableRootSquash = false,
    },new NfsAccessRule(NfsAccessRuleScope.Default,NfsAccessRuleAccess.No)
    {
    AllowSuid = false,
    AllowSubmountAccess = true,
    EnableRootSquash = true,
    AnonymousUID = "65534",
    AnonymousGID = "65534",
    }
    })
    },
    DirectoryServicesSettings = new StorageCacheDirectorySettings()
    {
        ActiveDirectory = new StorageCacheActiveDirectorySettings(IPAddress.Parse("192.0.2.10"), "contosoAd.contoso.local", "contosoAd", "contosoSmb")
        {
            SecondaryDnsIPAddress = IPAddress.Parse("192.0.2.11"),
        },
        UsernameDownload = new StorageCacheUsernameDownloadSettings()
        {
            EnableExtendedGroups = true,
            UsernameSource = StorageCacheUsernameSourceType.AD,
        },
    },
    Tags =
    {
    ["Dept"] = "Contoso",
    },
};
ArmOperation<StorageCacheResource> lro = await storageCache.UpdateAsync(WaitUntil.Completed, data);
StorageCacheResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageCacheData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
