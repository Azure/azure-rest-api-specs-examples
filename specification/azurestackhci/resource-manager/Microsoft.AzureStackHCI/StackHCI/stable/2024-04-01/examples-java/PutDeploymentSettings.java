
import com.azure.resourcemanager.azurestackhci.models.DeploymentCluster;
import com.azure.resourcemanager.azurestackhci.models.DeploymentConfiguration;
import com.azure.resourcemanager.azurestackhci.models.DeploymentData;
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;
import com.azure.resourcemanager.azurestackhci.models.DeploymentSecuritySettings;
import com.azure.resourcemanager.azurestackhci.models.DeploymentSettingAdapterPropertyOverrides;
import com.azure.resourcemanager.azurestackhci.models.DeploymentSettingHostNetwork;
import com.azure.resourcemanager.azurestackhci.models.DeploymentSettingIntents;
import com.azure.resourcemanager.azurestackhci.models.DeploymentSettingStorageAdapterIpInfo;
import com.azure.resourcemanager.azurestackhci.models.DeploymentSettingStorageNetworks;
import com.azure.resourcemanager.azurestackhci.models.DeploymentSettingVirtualSwitchConfigurationOverrides;
import com.azure.resourcemanager.azurestackhci.models.EceDeploymentSecrets;
import com.azure.resourcemanager.azurestackhci.models.EceSecrets;
import com.azure.resourcemanager.azurestackhci.models.InfrastructureNetwork;
import com.azure.resourcemanager.azurestackhci.models.IpPools;
import com.azure.resourcemanager.azurestackhci.models.NetworkController;
import com.azure.resourcemanager.azurestackhci.models.Observability;
import com.azure.resourcemanager.azurestackhci.models.OperationType;
import com.azure.resourcemanager.azurestackhci.models.OptionalServices;
import com.azure.resourcemanager.azurestackhci.models.PhysicalNodes;
import com.azure.resourcemanager.azurestackhci.models.QosPolicyOverrides;
import com.azure.resourcemanager.azurestackhci.models.SbeCredentials;
import com.azure.resourcemanager.azurestackhci.models.SbeDeploymentInfo;
import com.azure.resourcemanager.azurestackhci.models.SbePartnerInfo;
import com.azure.resourcemanager.azurestackhci.models.SbePartnerProperties;
import com.azure.resourcemanager.azurestackhci.models.ScaleUnits;
import com.azure.resourcemanager.azurestackhci.models.SdnIntegration;
import com.azure.resourcemanager.azurestackhci.models.Storage;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for DeploymentSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * PutDeploymentSettings.json
     */
    /**
     * Sample code: Create Deployment Settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createDeploymentSettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.deploymentSettings().define("default").withExistingCluster("test-rg", "myCluster")
            .withArcNodeResourceIds(Arrays.asList(
                "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1",
                "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2"))
            .withDeploymentMode(DeploymentMode.DEPLOY).withOperationType(OperationType.CLUSTER_PROVISIONING)
            .withDeploymentConfiguration(new DeploymentConfiguration().withVersion("string")
                .withScaleUnits(Arrays.asList(new ScaleUnits()
                    .withDeploymentData(new DeploymentData()
                        .withSecuritySettings(new DeploymentSecuritySettings()
                            .withHvciProtection(true).withDrtmProtection(true).withDriftControlEnforced(true)
                            .withCredentialGuardEnforced(false).withSmbSigningEnforced(true)
                            .withSmbClusterEncryption(false).withSideChannelMitigationEnforced(true)
                            .withBitlockerBootVolume(true).withBitlockerDataVolumes(true).withWdacEnforced(true))
                        .withObservability(new Observability().withStreamingDataClient(true).withEuLocation(false)
                            .withEpisodicDataUpload(true))
                        .withCluster(new DeploymentCluster()
                            .withName("testHCICluster").withWitnessType("Cloud").withWitnessPath("Cloud")
                            .withCloudAccountName("myasestoragacct").withAzureServiceEndpoint("core.windows.net"))
                        .withStorage(new Storage().withConfigurationMode("Express")).withNamingPrefix("ms169")
                        .withDomainFqdn("ASZ1PLab8.nttest.microsoft.com")
                        .withInfrastructureNetwork(Arrays.asList(
                            new InfrastructureNetwork().withSubnetMask("255.255.248.0").withGateway("255.255.248.0")
                                .withIpPools(Arrays.asList(
                                    new IpPools().withStartingAddress("10.57.48.60").withEndingAddress("10.57.48.66")))
                                .withDnsServers(Arrays.asList("10.57.50.90"))))
                        .withPhysicalNodes(
                            Arrays.asList(new PhysicalNodes().withName("ms169host").withIpv4Address("10.57.51.224"),
                                new PhysicalNodes().withName("ms154host").withIpv4Address("10.57.53.236")))
                        .withHostNetwork(new DeploymentSettingHostNetwork()
                            .withIntents(Arrays.asList(new DeploymentSettingIntents().withName("Compute_Management")
                                .withTrafficType(Arrays.asList("Compute", "Management"))
                                .withAdapter(Arrays.asList("Port2")).withOverrideVirtualSwitchConfiguration(false)
                                .withVirtualSwitchConfigurationOverrides(
                                    new DeploymentSettingVirtualSwitchConfigurationOverrides()
                                        .withEnableIov("True").withLoadBalancingAlgorithm("HyperVPort"))
                                .withOverrideQosPolicy(false)
                                .withQosPolicyOverrides(new QosPolicyOverrides().withPriorityValue8021ActionCluster("7")
                                    .withPriorityValue8021ActionSmb("3").withBandwidthPercentageSmb("50"))
                                .withOverrideAdapterProperty(false).withAdapterPropertyOverrides(
                                    new DeploymentSettingAdapterPropertyOverrides().withJumboPacket("1514")
                                        .withNetworkDirect("Enabled").withNetworkDirectTechnology("iWARP"))))
                            .withStorageNetworks(
                                Arrays.asList(new DeploymentSettingStorageNetworks().withName("Storage1Network")
                                    .withNetworkAdapterName("Port3").withVlanId("5")
                                    .withStorageAdapterIpInfo(Arrays
                                        .asList(new DeploymentSettingStorageAdapterIpInfo().withPhysicalNode("string")
                                            .withIpv4Address("10.57.48.60").withSubnetMask("255.255.248.0")))))
                            .withStorageConnectivitySwitchless(true).withEnableStorageAutoIp(false))
                        .withSdnIntegration(new SdnIntegration()
                            .withNetworkController(new NetworkController().withMacAddressPoolStart("00-0D-3A-1B-C7-21")
                                .withMacAddressPoolStop("00-0D-3A-1B-C7-29").withNetworkVirtualizationEnabled(true)))
                        .withAdouPath("OU=ms169,DC=ASZ1PLab8,DC=nttest,DC=microsoft,DC=com")
                        .withSecretsLocation("fakeTokenPlaceholder")
                        .withSecrets(Arrays.asList(
                            new EceDeploymentSecrets().withSecretName("fakeTokenPlaceholder")
                                .withEceSecretName(EceSecrets.fromString("BMCAdminUserCred"))
                                .withSecretLocation("fakeTokenPlaceholder"),
                            new EceDeploymentSecrets().withSecretName("fakeTokenPlaceholder")
                                .withEceSecretName(EceSecrets.AZURE_STACK_LCMUSER_CREDENTIAL)
                                .withSecretLocation("fakeTokenPlaceholder")))
                        .withOptionalServices(new OptionalServices().withCustomLocation("customLocationName")))
                    .withSbePartnerInfo(new SbePartnerInfo()
                        .withSbeDeploymentInfo(new SbeDeploymentInfo().withVersion("4.0.2309.13").withFamily("Gen5")
                            .withPublisher("Contoso").withSbeManifestSource("default")
                            .withSbeManifestCreationDate(OffsetDateTime.parse("2023-07-25T02:40:33Z")))
                        .withPartnerProperties(
                            Arrays.asList(new SbePartnerProperties().withName("EnableBMCIpV6").withValue("false"),
                                new SbePartnerProperties().withName("PhoneHomePort").withValue("1653"),
                                new SbePartnerProperties().withName("BMCSecurityState").withValue("HighSecurity")))
                        .withCredentialList(Arrays.asList(new SbeCredentials().withSecretName("fakeTokenPlaceholder")
                            .withEceSecretName("fakeTokenPlaceholder").withSecretLocation("fakeTokenPlaceholder")))))))
            .create();
    }
}
