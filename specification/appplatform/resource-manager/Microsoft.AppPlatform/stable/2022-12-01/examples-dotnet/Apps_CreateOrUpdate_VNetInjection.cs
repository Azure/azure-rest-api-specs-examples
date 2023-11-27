using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppPlatform;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Apps_CreateOrUpdate_VNetInjection.json
// this example is just showing the usage of "Apps_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformServiceResource created on azure
// for more information of creating AppPlatformServiceResource, please refer to the document of AppPlatformServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
ResourceIdentifier appPlatformServiceResourceId = AppPlatformServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
AppPlatformServiceResource appPlatformService = client.GetAppPlatformServiceResource(appPlatformServiceResourceId);

// get the collection of this AppPlatformAppResource
AppPlatformAppCollection collection = appPlatformService.GetAppPlatformApps();

// invoke the operation
string appName = "myapp";
AppPlatformAppData data = new AppPlatformAppData()
{
    Properties = new AppPlatformAppProperties()
    {
        IsPublic = true,
        AddonConfigs =
        {
        ["ApplicationConfigurationService"] = new Dictionary<string, BinaryData>()
        {
        ["resourceId"] = BinaryData.FromString("\"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/configurationServices/myacs\""),
        },
        ["ServiceRegistry"] = new Dictionary<string, BinaryData>()
        {
        ["resourceId"] = BinaryData.FromString("\"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/serviceRegistries/myServiceRegistry\""),
        },
        },
        IsHttpsOnly = false,
        TemporaryDisk = new AppTemporaryDisk()
        {
            SizeInGB = 2,
            MountPath = "/mytemporarydisk",
        },
        PersistentDisk = new AppPersistentDisk()
        {
            SizeInGB = 2,
            MountPath = "/mypersistentdisk",
        },
        CustomPersistentDisks =
        {
        new AppCustomPersistentDisk("myASCStorageID")
        {
        CustomPersistentDiskProperties = new AppPlatformAzureFileVolume("/mypath1/mypath2","myFileShare")
        {
        MountOptions =
        {
        "uid=0","gid=0","dir_mode=0777","file_mode=0777"
        },
        },
        }
        },
        IsEndToEndTlsEnabled = false,
        LoadedCertificates =
        {
        new AppLoadedCertificate(new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert1"))
        {
        LoadTrustStore = false,
        },new AppLoadedCertificate(new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert2"))
        {
        LoadTrustStore = true,
        }
        },
        VnetAddons = new AppVnetAddons()
        {
            IsPublicEndpoint = true,
        },
    },
    Identity = new ManagedServiceIdentity("SystemAssigned,UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/samplegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity(),
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/samplegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2")] = new UserAssignedIdentity(),
        },
    },
    Location = new AzureLocation("eastus"),
};
ArmOperation<AppPlatformAppResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, appName, data);
AppPlatformAppResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformAppData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
