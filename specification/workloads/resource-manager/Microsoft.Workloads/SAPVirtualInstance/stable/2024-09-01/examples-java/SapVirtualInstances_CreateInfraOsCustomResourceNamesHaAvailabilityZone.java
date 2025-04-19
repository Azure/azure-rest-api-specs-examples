
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ApplicationServerConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ApplicationServerFullResourceNames;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.CentralServerConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.CentralServerFullResourceNames;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DatabaseConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DatabaseServerFullResourceNames;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DeploymentWithOSConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.HighAvailabilityConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ImageReference;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.LinuxConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.LoadBalancerResourceNames;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.NetworkInterfaceResourceNames;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.OSProfile;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.OsSapConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapHighAvailabilityType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapVirtualInstanceProperties;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SharedStorageResourceNames;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SshKeyPair;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ThreeTierConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ThreeTierFullResourceNames;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.VirtualMachineConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.VirtualMachineResourceNames;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapVirtualInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateInfraOsCustomResourceNamesHaAvailabilityZone.json
     */
    /**
     * Sample code: Create Infrastructure (with OS configuration) with custom resource names for HA system with
     * Availability Zone.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void createInfrastructureWithOSConfigurationWithCustomResourceNamesForHASystemWithAvailabilityZone(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().define("X00").withRegion("westcentralus").withExistingResourceGroup("test-rg")
            .withTags(mapOf())
            .withProperties(new SapVirtualInstanceProperties().withEnvironment(SapEnvironmentType.PROD)
                .withSapProduct(SapProductType.S4HANA)
                .withConfiguration(new DeploymentWithOSConfiguration().withAppLocation("eastus")
                    .withInfrastructureConfiguration(new ThreeTierConfiguration().withAppResourceGroup("X00-RG")
                        .withCentralServer(new CentralServerConfiguration().withSubnetId(
                            "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                            .withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                                .withVmSize("Standard_E16ds_v4")
                                .withImageReference(new ImageReference().withPublisher("RedHat").withOffer("RHEL-SAP")
                                    .withSku("84sapha-gen2").withVersion("latest"))
                                .withOsProfile(new OSProfile().withAdminUsername("{your-username}").withOsConfiguration(
                                    new LinuxConfiguration().withDisablePasswordAuthentication(true)
                                        .withSshKeyPair(new SshKeyPair().withPublicKey("fakeTokenPlaceholder")
                                            .withPrivateKey("fakeTokenPlaceholder")))))
                            .withInstanceCount(2L))
                        .withApplicationServer(new ApplicationServerConfiguration().withSubnetId(
                            "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                            .withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                                .withVmSize("Standard_E32ds_v4")
                                .withImageReference(new ImageReference().withPublisher("RedHat").withOffer("RHEL-SAP")
                                    .withSku("84sapha-gen2").withVersion("latest"))
                                .withOsProfile(new OSProfile().withAdminUsername("{your-username}")
                                    .withOsConfiguration(new LinuxConfiguration()
                                        .withDisablePasswordAuthentication(true)
                                        .withSshKeyPair(new SshKeyPair().withPublicKey("fakeTokenPlaceholder")
                                            .withPrivateKey("fakeTokenPlaceholder")))))
                            .withInstanceCount(6L))
                        .withDatabaseServer(new DatabaseConfiguration().withDatabaseType(SapDatabaseType.HANA)
                            .withSubnetId(
                                "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet")
                            .withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                                .withVmSize("Standard_M32ts")
                                .withImageReference(new ImageReference().withPublisher("RedHat").withOffer("RHEL-SAP")
                                    .withSku("84sapha-gen2").withVersion("latest"))
                                .withOsProfile(new OSProfile().withAdminUsername("{your-username}").withOsConfiguration(
                                    new LinuxConfiguration().withDisablePasswordAuthentication(true)
                                        .withSshKeyPair(new SshKeyPair().withPublicKey("fakeTokenPlaceholder")
                                            .withPrivateKey("fakeTokenPlaceholder")))))
                            .withInstanceCount(2L))
                        .withHighAvailabilityConfig(new HighAvailabilityConfiguration()
                            .withHighAvailabilityType(SapHighAvailabilityType.AVAILABILITY_ZONE))
                        .withCustomResourceNames(
                            new ThreeTierFullResourceNames()
                                .withCentralServer(new CentralServerFullResourceNames()
                                    .withVirtualMachines(Arrays.asList(
                                        new VirtualMachineResourceNames()
                                            .withVmName("ascsvm").withHostName("ascshostName")
                                            .withNetworkInterfaces(Arrays
                                                .asList(new NetworkInterfaceResourceNames()
                                                    .withNetworkInterfaceName("ascsnic")))
                                            .withOsDiskName("ascsosdisk"),
                                        new VirtualMachineResourceNames().withVmName("ersvm")
                                            .withHostName("ershostName")
                                            .withNetworkInterfaces(Arrays.asList(
                                                new NetworkInterfaceResourceNames().withNetworkInterfaceName("ersnic")))
                                            .withOsDiskName("ersosdisk")))
                                    .withLoadBalancer(new LoadBalancerResourceNames().withLoadBalancerName("ascslb")
                                        .withFrontendIpConfigurationNames(Arrays.asList("ascsip0", "ersip0"))
                                        .withBackendPoolNames(Arrays.asList("ascsBackendPool"))
                                        .withHealthProbeNames(Arrays.asList("ascsHealthProbe", "ersHealthProbe"))))
                                .withApplicationServer(
                                    new ApplicationServerFullResourceNames().withVirtualMachines(Arrays.asList(
                                        new VirtualMachineResourceNames().withVmName("appvm0")
                                            .withHostName("apphostName0")
                                            .withNetworkInterfaces(Arrays.asList(new NetworkInterfaceResourceNames()
                                                .withNetworkInterfaceName("appnic0")))
                                            .withOsDiskName("app0osdisk")
                                            .withDataDiskNames(mapOf("default", Arrays.asList("app0disk0"))),
                                        new VirtualMachineResourceNames()
                                            .withVmName("appvm1").withHostName("apphostName1")
                                            .withNetworkInterfaces(Arrays.asList(new NetworkInterfaceResourceNames()
                                                .withNetworkInterfaceName("appnic1")))
                                            .withOsDiskName("app1osdisk")
                                            .withDataDiskNames(mapOf("default", Arrays.asList("app1disk0"))))))
                                .withDatabaseServer(new DatabaseServerFullResourceNames()
                                    .withVirtualMachines(Arrays.asList(
                                        new VirtualMachineResourceNames()
                                            .withVmName("dbvmpr").withHostName("dbprhostName")
                                            .withNetworkInterfaces(Arrays.asList(new NetworkInterfaceResourceNames()
                                                .withNetworkInterfaceName("dbprnic")))
                                            .withOsDiskName("dbprosdisk").withDataDiskNames(
                                                mapOf("hanaData", Arrays.asList("hanadatapr0", "hanadatapr1"),
                                                    "hanaLog", Arrays.asList("hanalogpr0", "hanalogpr1", "hanalogpr2"),
                                                    "hanaShared", Arrays.asList("hanasharedpr0", "hanasharedpr1"),
                                                    "usrSap", Arrays.asList("usrsappr0"))),
                                        new VirtualMachineResourceNames().withVmName("dbvmsr")
                                            .withHostName("dbsrhostName")
                                            .withNetworkInterfaces(Arrays.asList(new NetworkInterfaceResourceNames()
                                                .withNetworkInterfaceName("dbsrnic")))
                                            .withOsDiskName("dbsrosdisk").withDataDiskNames(
                                                mapOf("hanaData", Arrays.asList("hanadatasr0", "hanadatasr1"),
                                                    "hanaLog", Arrays.asList("hanalogsr0", "hanalogsr1", "hanalogsr2"),
                                                    "hanaShared", Arrays.asList("hanasharedsr0", "hanasharedsr1"),
                                                    "usrSap", Arrays.asList("usrsapsr0")))))
                                    .withLoadBalancer(new LoadBalancerResourceNames().withLoadBalancerName("dblb")
                                        .withFrontendIpConfigurationNames(Arrays.asList("dbip"))
                                        .withBackendPoolNames(Arrays.asList("dbBackendPool"))
                                        .withHealthProbeNames(Arrays.asList("dbHealthProbe"))))
                                .withSharedStorage(
                                    new SharedStorageResourceNames().withSharedStorageAccountName("storageacc")
                                        .withSharedStorageAccountPrivateEndPointName("peForxNFS"))))
                    .withOsSapConfiguration(new OsSapConfiguration().withSapFqdn("xyz.test.com"))))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
