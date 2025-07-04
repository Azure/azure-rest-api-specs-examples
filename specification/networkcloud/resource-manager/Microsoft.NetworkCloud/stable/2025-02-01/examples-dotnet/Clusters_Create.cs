using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/Clusters_Create.json
// this example is just showing the usage of "Clusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkCloudClusterResource
NetworkCloudClusterCollection collection = resourceGroupResource.GetNetworkCloudClusters();

// invoke the operation
string clusterName = "clusterName";
NetworkCloudClusterData data = new NetworkCloudClusterData(
    new AzureLocation("location"),
    new ExtendedLocation("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterManagerExtendedLocationName", "CustomLocation"),
    new NetworkCloudRackDefinition(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"), "AA1234", new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"))
    {
        BareMetalMachineConfigurationData = {new BareMetalMachineConfiguration(new AdministrativeCredentials("username")
        {
        Password = "https://keyvaultname.vault.azure.net/secrets/secretName",
        }, "AA:BB:CC:DD:EE:FF", "00:BB:CC:DD:EE:FF", 1L, "BM1219XXX")
        {
        MachineDetails = "extraDetails",
        MachineName = "bmmName1",
        }, new BareMetalMachineConfiguration(new AdministrativeCredentials("username")
        {
        Password = "https://keyvaultname.vault.azure.net/secrets/secretName",
        }, "AA:BB:CC:DD:EE:00", "00:BB:CC:DD:EE:00", 2L, "BM1219YYY")
        {
        MachineDetails = "extraDetails",
        MachineName = "bmmName2",
        }},
        RackLocation = "Foo Datacenter, Floor 3, Aisle 9, Rack 2",
        StorageApplianceConfigurationData = {new StorageApplianceConfiguration(new AdministrativeCredentials("username")
        {
        Password = "https://keyvaultname.vault.azure.net/secrets/secretName",
        }, 1L, "BM1219XXX")
        {
        StorageApplianceName = "vmName",
        }},
    },
    ClusterType.SingleRack,
    "1.0.0",
    new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabrics/fabricName"))
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1")] = new UserAssignedIdentity()
        },
    },
    AnalyticsOutputSettings = new AnalyticsOutputSettings
    {
        AnalyticsWorkspaceId = new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName"),
        AssociatedIdentity = new ManagedServiceIdentitySelector
        {
            IdentityType = ManagedServiceIdentitySelectorType.UserAssignedIdentity,
            UserAssignedIdentityResourceId = new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1"),
        },
    },
    ClusterLocation = "Foo Street, 3rd Floor, row 9",
    ClusterServicePrincipal = new ServicePrincipalInformation("12345678-1234-1234-1234-123456789012", "00000008-0004-0004-0004-000000000012", "80000000-4000-4000-4000-120000000000")
    {
        Password = "{password}",
    },
    CommandOutputSettings = new CommandOutputSettings
    {
        AssociatedIdentity = new ManagedServiceIdentitySelector
        {
            IdentityType = ManagedServiceIdentitySelectorType.UserAssignedIdentity,
            UserAssignedIdentityResourceId = new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1"),
        },
        ContainerUri = new Uri("https://myaccount.blob.core.windows.net/mycontainer?restype=container"),
    },
    ComputeDeploymentThreshold = new ValidationThreshold(ValidationThresholdGrouping.PerCluster, ValidationThresholdType.PercentSuccess, 90L),
    ComputeRackDefinitions = {new NetworkCloudRackDefinition(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"), "AA1234", new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"))
    {
    BareMetalMachineConfigurationData = {new BareMetalMachineConfiguration(new AdministrativeCredentials("username")
    {
    Password = "https://keyvaultname.vault.azure.net/secrets/secretName",
    }, "AA:BB:CC:DD:EE:FF", "00:BB:CC:DD:EE:FF", 1L, "BM1219XXX")
    {
    MachineDetails = "extraDetails",
    MachineName = "bmmName1",
    }, new BareMetalMachineConfiguration(new AdministrativeCredentials("username")
    {
    Password = "https://keyvaultname.vault.azure.net/secrets/secretName",
    }, "AA:BB:CC:DD:EE:00", "00:BB:CC:DD:EE:00", 2L, "BM1219YYY")
    {
    MachineDetails = "extraDetails",
    MachineName = "bmmName2",
    }},
    RackLocation = "Foo Datacenter, Floor 3, Aisle 9, Rack 2",
    StorageApplianceConfigurationData = {new StorageApplianceConfiguration(new AdministrativeCredentials("username")
    {
    Password = "https://keyvaultname.vault.azure.net/secrets/secretName",
    }, 1L, "BM1219XXX")
    {
    StorageApplianceName = "vmName",
    }},
    }},
    ManagedResourceGroupConfiguration = new ManagedResourceGroupConfiguration
    {
        Location = new AzureLocation("East US"),
        Name = "my-managed-rg",
    },
    RuntimeProtectionEnforcementLevel = RuntimeProtectionEnforcementLevel.OnDemand,
    SecretArchiveSettings = new SecretArchiveSettings
    {
        AssociatedIdentity = new ManagedServiceIdentitySelector
        {
            IdentityType = ManagedServiceIdentitySelectorType.UserAssignedIdentity,
            UserAssignedIdentityResourceId = new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1"),
        },
        VaultUri = new Uri("https://keyvaultname.vault.azure.net/"),
    },
    UpdateStrategy = new ClusterUpdateStrategy(ClusterUpdateStrategyType.Rack, ValidationThresholdType.CountSuccess, 4L)
    {
        MaxUnavailable = 4L,
        WaitTimeMinutes = 10L,
    },
    VulnerabilityScanningContainerScan = VulnerabilityScanningSettingsContainerScan.Enabled,
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
};
ArmOperation<NetworkCloudClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
NetworkCloudClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
