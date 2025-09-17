using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerInstance.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerInstance;

// Generated from example definition: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-11-01-preview/examples/ContainerGroupProfilesCreateOrUpdate.json
// this example is just showing the usage of "CGProfile_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "demo";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerGroupProfileResource
ContainerGroupProfileCollection collection = resourceGroupResource.GetContainerGroupProfiles();

// invoke the operation
string containerGroupProfileName = "demo1";
ContainerGroupProfileData data = new ContainerGroupProfileData(new AzureLocation("west us"))
{
    Containers = {new ContainerInstanceContainer("demo1")
    {
    Image = "nginx",
    Command = {},
    Ports = {new ContainerPort(80)},
    EnvironmentVariables = {},
    Resources = new ContainerResourceRequirements(new ContainerResourceRequestsContent(1.5, 1)
    {
    Gpu = new ContainerGpuResourceInfo(1, ContainerGpuSku.K80),
    }),
    VolumeMounts = {new ContainerVolumeMount("volume1", "/mnt/volume1")
    {
    IsReadOnly = false,
    }, new ContainerVolumeMount("volume2", "/mnt/volume2")
    {
    IsReadOnly = false,
    }, new ContainerVolumeMount("volume3", "/mnt/volume3")
    {
    IsReadOnly = true,
    }},
    }},
    ImageRegistryCredentials = { },
    IPAddress = new ContainerGroupIPAddress(new ContainerGroupPort[]
{
new ContainerGroupPort(80)
{
Protocol = ContainerGroupNetworkProtocol.Tcp,
}
}, ContainerGroupIPAddressType.Public),
    OSType = ContainerInstanceOperatingSystemType.Linux,
    Volumes = {new ContainerVolume("volume1")
    {
    AzureFile = new ContainerInstanceAzureFileVolume("shareName", "accountName")
    {
    StorageAccountKey = "accountKey",
    },
    }, new ContainerVolume("volume2")
    {
    EmptyDir = BinaryData.FromObjectAsJson(new object()),
    }, new ContainerVolume("volume3")
    {
    Secret =
    {
    ["secretKey1"] = "SecretValue1InBase64",
    ["secretKey2"] = "SecretValue2InBase64"
    },
    }},
    DiagnosticsLogAnalytics = new ContainerGroupLogAnalytics("workspaceid", "workspaceKey")
    {
        LogType = ContainerGroupLogAnalyticsLogType.ContainerInsights,
        Metadata =
        {
        ["pod-uuid"] = "test-metadata-value"
        },
        WorkspaceResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/microsoft.operationalinsights/workspaces/workspace"),
    },
    Zones = { "1" },
};
ArmOperation<ContainerGroupProfileResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, containerGroupProfileName, data);
ContainerGroupProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerGroupProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
