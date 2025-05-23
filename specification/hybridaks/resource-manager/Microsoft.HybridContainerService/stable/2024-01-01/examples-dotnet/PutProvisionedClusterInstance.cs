using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridContainerService.Models;
using Azure.ResourceManager.HybridContainerService;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/PutProvisionedClusterInstance.json
// this example is just showing the usage of "provisionedClusterInstances_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProvisionedClusterResource created on azure
// for more information of creating ProvisionedClusterResource, please refer to the document of ProvisionedClusterResource
string connectedClusterResourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster";
ResourceIdentifier provisionedClusterResourceId = ProvisionedClusterResource.CreateResourceIdentifier(connectedClusterResourceUri);
ProvisionedClusterResource provisionedCluster = client.GetProvisionedClusterResource(provisionedClusterResourceId);

// invoke the operation
ProvisionedClusterData data = new ProvisionedClusterData
{
    Properties = new ProvisionedClusterProperties
    {
        SshPublicKeys = {new LinuxSshPublicKey
        {
        KeyData = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY.......",
        }},
        ControlPlane = new ProvisionedClusterControlPlaneProfile
        {
            Count = 1,
            VmSize = "Standard_A4_v2",
        },
        KubernetesVersion = "v1.20.5",
        NetworkProfile = new ProvisionedClusterNetworkProfile
        {
            NetworkPolicy = ProvisionedClusterNetworkPolicy.Calico,
            PodCidr = "10.244.0.0/16",
        },
        ClusterVmAccessAuthorizedIPRanges = "127.0.0.1,127.0.0.2",
        AgentPoolProfiles = {new HybridContainerServiceNamedAgentPoolProfile
        {
        Count = 1,
        VmSize = "Standard_A4_v2",
        Name = "default-nodepool-1",
        OSType = HybridContainerServiceOSType.Linux,
        NodeLabels =
        {
        ["env"] = "dev",
        ["goal"] = "test"
        },
        NodeTaints = {"env=prod:NoSchedule", "sku=gpu:NoSchedule"},
        }},
        InfraNetworkVnetSubnetIds = { new ResourceIdentifier("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.AzureStackHCI/logicalNetworks/test-vnet-static") },
        LicenseAzureHybridBenefit = ProvisionedClusterAzureHybridBenefit.NotApplicable,
    },
    ExtendedLocation = new HybridContainerServiceExtendedLocation
    {
        ExtendedLocationType = HybridContainerServiceExtendedLocationType.CustomLocation,
        Name = "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation",
    },
};
ArmOperation<ProvisionedClusterResource> lro = await provisionedCluster.CreateOrUpdateAsync(WaitUntil.Completed, data);
ProvisionedClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProvisionedClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
