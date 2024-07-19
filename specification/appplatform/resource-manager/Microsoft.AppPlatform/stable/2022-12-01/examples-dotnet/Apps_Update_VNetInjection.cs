using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Apps_Update_VNetInjection.json
// this example is just showing the usage of "Apps_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformAppResource created on azure
// for more information of creating AppPlatformAppResource, please refer to the document of AppPlatformAppResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string appName = "myapp";
ResourceIdentifier appPlatformAppResourceId = AppPlatformAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, appName);
AppPlatformAppResource appPlatformApp = client.GetAppPlatformAppResource(appPlatformAppResourceId);

// invoke the operation
AppPlatformAppData data = new AppPlatformAppData()
{
    Properties = new AppPlatformAppProperties()
    {
        IsPublic = true,
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
        },
        },
        }
        },
        IsEndToEndTlsEnabled = false,
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
ArmOperation<AppPlatformAppResource> lro = await appPlatformApp.UpdateAsync(WaitUntil.Completed, data);
AppPlatformAppResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformAppData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
