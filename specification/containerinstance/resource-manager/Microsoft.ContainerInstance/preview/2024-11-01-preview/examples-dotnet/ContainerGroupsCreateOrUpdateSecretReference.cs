using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerInstance.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerInstance;

// Generated from example definition: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-11-01-preview/examples/ContainerGroupsCreateOrUpdateSecretReference.json
// this example is just showing the usage of "ContainerGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ContainerGroupResource
ContainerGroupCollection collection = resourceGroupResource.GetContainerGroups();

// invoke the operation
string containerGroupName = "demo1";
ContainerGroupData data = new ContainerGroupData(new AzureLocation("west us"), new ContainerInstanceContainer[]
{
new ContainerInstanceContainer("demo1", "privateRegistryImage", new ContainerResourceRequirements(new ContainerResourceRequestsContent(1.5, 1)
{
Gpu = new ContainerGpuResourceInfo(1, ContainerGpuSku.K80),
}))
{
Command = {},
Ports = {new ContainerPort(80)},
EnvironmentVariables = {new ContainerEnvironmentVariable("envSecret")
{
SecureValueReference = "envSecretRef",
}},
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
}
}, ContainerInstanceOperatingSystemType.Linux)
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name")] = new UserAssignedIdentity()
        },
    },
    SecretReferences = { new ContainerGroupSecretReference("envSecretRef", new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name"), new Uri("https://keyvaultname.vault.azure.net/secrets/envSecret")), new ContainerGroupSecretReference("accountKeyRef", new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name"), new Uri("https://keyvaultname.vault.azure.net/secrets/accountKey")), new ContainerGroupSecretReference("volumeSecretRef", new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name"), new Uri("https://keyvaultname.vault.azure.net/secrets/volumeSecret")), new ContainerGroupSecretReference("privateRegistryKeyRef", new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name"), new Uri("https://keyvaultname.vault.azure.net/secrets/privateRegistryKey")) },
    ImageRegistryCredentials = {new ContainerGroupImageRegistryCredential("demoregistry.azurecr.io")
    {
    Username = "registryUserName",
    PasswordReference = "privateRegistryKeyRef",
    }},
    IPAddress = new ContainerGroupIPAddress(new ContainerGroupPort[]
{
new ContainerGroupPort(80)
{
Protocol = ContainerGroupNetworkProtocol.Tcp,
}
}, ContainerGroupIPAddressType.Public)
    {
        DnsNameLabel = "dnsnamelabel1",
        AutoGeneratedDomainNameLabelScope = DnsNameLabelReusePolicy.Unsecure,
    },
    Volumes = {new ContainerVolume("volume1")
    {
    AzureFile = new ContainerInstanceAzureFileVolume("shareName", "accountName")
    {
    StorageAccountKeyReference = "accountKeyRef",
    },
    }, new ContainerVolume("volume2")
    {
    EmptyDir = BinaryData.FromObjectAsJson(new object()),
    }, new ContainerVolume("volume3")
    {
    Secret =
    {
    ["secretKey1"] = "SecretValue1InBase64"
    },
    SecretReference =
    {
    ["secretKey2"] = "volumeSecretRef"
    },
    }},
    DiagnosticsLogAnalytics = new ContainerGroupLogAnalytics("workspaceid", "workspaceKey")
    {
        LogType = ContainerGroupLogAnalyticsLogType.ContainerInsights,
        Metadata =
        {
        ["test-key"] = "test-metadata-value"
        },
        WorkspaceResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/microsoft.operationalinsights/workspaces/workspace"),
    },
    SubnetIds = { new ContainerGroupSubnetId(new ResourceIdentifier("[resourceId('Microsoft.Network/virtualNetworks/subnets', parameters('vnetName'), parameters('subnetName'))]")) },
    DnsConfig = new ContainerGroupDnsConfiguration(new string[] { "1.1.1.1" })
    {
        SearchDomains = "cluster.local svc.cluster.local",
        Options = "ndots:2",
    },
};
ArmOperation<ContainerGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, containerGroupName, data);
ContainerGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
