import com.azure.resourcemanager.workloads.models.ApplicationServerConfiguration;
import com.azure.resourcemanager.workloads.models.ApplicationServerFullResourceNames;
import com.azure.resourcemanager.workloads.models.CentralServerConfiguration;
import com.azure.resourcemanager.workloads.models.CentralServerFullResourceNames;
import com.azure.resourcemanager.workloads.models.DatabaseConfiguration;
import com.azure.resourcemanager.workloads.models.DatabaseServerFullResourceNames;
import com.azure.resourcemanager.workloads.models.DeploymentWithOSConfiguration;
import com.azure.resourcemanager.workloads.models.HighAvailabilityConfiguration;
import com.azure.resourcemanager.workloads.models.ImageReference;
import com.azure.resourcemanager.workloads.models.LinuxConfiguration;
import com.azure.resourcemanager.workloads.models.LoadBalancerResourceNames;
import com.azure.resourcemanager.workloads.models.NetworkInterfaceResourceNames;
import com.azure.resourcemanager.workloads.models.OSProfile;
import com.azure.resourcemanager.workloads.models.OsSapConfiguration;
import com.azure.resourcemanager.workloads.models.SapDatabaseType;
import com.azure.resourcemanager.workloads.models.SapEnvironmentType;
import com.azure.resourcemanager.workloads.models.SapHighAvailabilityType;
import com.azure.resourcemanager.workloads.models.SapProductType;
import com.azure.resourcemanager.workloads.models.SharedStorageResourceNames;
import com.azure.resourcemanager.workloads.models.SshKeyPair;
import com.azure.resourcemanager.workloads.models.ThreeTierConfiguration;
import com.azure.resourcemanager.workloads.models.ThreeTierFullResourceNames;
import com.azure.resourcemanager.workloads.models.VirtualMachineConfiguration;
import com.azure.resourcemanager.workloads.models.VirtualMachineResourceNames;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for SapVirtualInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Create_CustomFullResourceNames_HA_AvZone.json
     */
    /**
     * Sample code: Create Infrastructure (with OS configuration) with custom resource names for HA system with
     * Availability Zone.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createInfrastructureWithOSConfigurationWithCustomResourceNamesForHASystemWithAvailabilityZone(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapVirtualInstances()
            .define("X00")
            .withRegion("westcentralus")
            .withExistingResourceGroup("test-rg")
            .withEnvironment(SapEnvironmentType.PROD)
            .withSapProduct(SapProductType.S4HANA)
            .withConfiguration(
                new DeploymentWithOSConfiguration()
                    .withAppLocation("eastus")
                    .withInfrastructureConfiguration(
                        new ThreeTierConfiguration()
                            .withAppResourceGroup("X00-RG")
                            .withCentralServer(
                                new CentralServerConfiguration()
                                    .withSubnetId(
                                        "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                                    .withVirtualMachineConfiguration(
                                        new VirtualMachineConfiguration()
                                            .withVmSize("Standard_E16ds_v4")
                                            .withImageReference(
                                                new ImageReference()
                                                    .withPublisher("RedHat")
                                                    .withOffer("RHEL-SAP")
                                                    .withSku("84sapha-gen2")
                                                    .withVersion("latest"))
                                            .withOsProfile(
                                                new OSProfile()
                                                    .withAdminUsername("{your-username}")
                                                    .withOsConfiguration(
                                                        new LinuxConfiguration()
                                                            .withDisablePasswordAuthentication(true)
                                                            .withSshKeyPair(
                                                                new SshKeyPair()
                                                                    .withPublicKey("fakeTokenPlaceholder")
                                                                    .withPrivateKey("fakeTokenPlaceholder")))))
                                    .withInstanceCount(2L))
                            .withApplicationServer(
                                new ApplicationServerConfiguration()
                                    .withSubnetId(
                                        "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                                    .withVirtualMachineConfiguration(
                                        new VirtualMachineConfiguration()
                                            .withVmSize("Standard_E32ds_v4")
                                            .withImageReference(
                                                new ImageReference()
                                                    .withPublisher("RedHat")
                                                    .withOffer("RHEL-SAP")
                                                    .withSku("84sapha-gen2")
                                                    .withVersion("latest"))
                                            .withOsProfile(
                                                new OSProfile()
                                                    .withAdminUsername("{your-username}")
                                                    .withOsConfiguration(
                                                        new LinuxConfiguration()
                                                            .withDisablePasswordAuthentication(true)
                                                            .withSshKeyPair(
                                                                new SshKeyPair()
                                                                    .withPublicKey("fakeTokenPlaceholder")
                                                                    .withPrivateKey("fakeTokenPlaceholder")))))
                                    .withInstanceCount(6L))
                            .withDatabaseServer(
                                new DatabaseConfiguration()
                                    .withDatabaseType(SapDatabaseType.HANA)
                                    .withSubnetId(
                                        "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet")
                                    .withVirtualMachineConfiguration(
                                        new VirtualMachineConfiguration()
                                            .withVmSize("Standard_M32ts")
                                            .withImageReference(
                                                new ImageReference()
                                                    .withPublisher("RedHat")
                                                    .withOffer("RHEL-SAP")
                                                    .withSku("84sapha-gen2")
                                                    .withVersion("latest"))
                                            .withOsProfile(
                                                new OSProfile()
                                                    .withAdminUsername("{your-username}")
                                                    .withOsConfiguration(
                                                        new LinuxConfiguration()
                                                            .withDisablePasswordAuthentication(true)
                                                            .withSshKeyPair(
                                                                new SshKeyPair()
                                                                    .withPublicKey("fakeTokenPlaceholder")
                                                                    .withPrivateKey("fakeTokenPlaceholder")))))
                                    .withInstanceCount(2L))
                            .withHighAvailabilityConfig(
                                new HighAvailabilityConfiguration()
                                    .withHighAvailabilityType(SapHighAvailabilityType.AVAILABILITY_ZONE))
                            .withCustomResourceNames(
                                new ThreeTierFullResourceNames()
                                    .withCentralServer(
                                        new CentralServerFullResourceNames()
                                            .withVirtualMachines(
                                                Arrays
                                                    .asList(
                                                        new VirtualMachineResourceNames()
                                                            .withVmName("ascsvm")
                                                            .withHostname("ascshostName")
                                                            .withNetworkInterfaces(
                                                                Arrays
                                                                    .asList(
                                                                        new NetworkInterfaceResourceNames()
                                                                            .withNetworkInterfaceName("ascsnic")))
                                                            .withOsDiskName("ascsosdisk"),
                                                        new VirtualMachineResourceNames()
                                                            .withVmName("ersvm")
                                                            .withHostname("ershostName")
                                                            .withNetworkInterfaces(
                                                                Arrays
                                                                    .asList(
                                                                        new NetworkInterfaceResourceNames()
                                                                            .withNetworkInterfaceName("ersnic")))
                                                            .withOsDiskName("ersosdisk")))
                                            .withLoadBalancer(
                                                new LoadBalancerResourceNames()
                                                    .withLoadBalancerName("ascslb")
                                                    .withFrontendIpConfigurationNames(
                                                        Arrays.asList("ascsip0", "ersip0"))
                                                    .withBackendPoolNames(Arrays.asList("ascsBackendPool"))
                                                    .withHealthProbeNames(
                                                        Arrays.asList("ascsHealthProbe", "ersHealthProbe"))))
                                    .withApplicationServer(
                                        new ApplicationServerFullResourceNames()
                                            .withVirtualMachines(
                                                Arrays
                                                    .asList(
                                                        new VirtualMachineResourceNames()
                                                            .withVmName("appvm0")
                                                            .withHostname("apphostName0")
                                                            .withNetworkInterfaces(
                                                                Arrays
                                                                    .asList(
                                                                        new NetworkInterfaceResourceNames()
                                                                            .withNetworkInterfaceName("appnic0")))
                                                            .withOsDiskName("app0osdisk")
                                                            .withDataDiskNames(
                                                                mapOf("default", Arrays.asList("app0disk0"))),
                                                        new VirtualMachineResourceNames()
                                                            .withVmName("appvm1")
                                                            .withHostname("apphostName1")
                                                            .withNetworkInterfaces(
                                                                Arrays
                                                                    .asList(
                                                                        new NetworkInterfaceResourceNames()
                                                                            .withNetworkInterfaceName("appnic1")))
                                                            .withOsDiskName("app1osdisk")
                                                            .withDataDiskNames(
                                                                mapOf("default", Arrays.asList("app1disk0"))))))
                                    .withDatabaseServer(
                                        new DatabaseServerFullResourceNames()
                                            .withVirtualMachines(
                                                Arrays
                                                    .asList(
                                                        new VirtualMachineResourceNames()
                                                            .withVmName("dbvmpr")
                                                            .withHostname("dbprhostName")
                                                            .withNetworkInterfaces(
                                                                Arrays
                                                                    .asList(
                                                                        new NetworkInterfaceResourceNames()
                                                                            .withNetworkInterfaceName("dbprnic")))
                                                            .withOsDiskName("dbprosdisk")
                                                            .withDataDiskNames(
                                                                mapOf(
                                                                    "hanaData",
                                                                    Arrays.asList("hanadatapr0", "hanadatapr1"),
                                                                    "hanaLog",
                                                                    Arrays
                                                                        .asList(
                                                                            "hanalogpr0", "hanalogpr1", "hanalogpr2"),
                                                                    "hanaShared",
                                                                    Arrays.asList("hanasharedpr0", "hanasharedpr1"),
                                                                    "usrSap",
                                                                    Arrays.asList("usrsappr0"))),
                                                        new VirtualMachineResourceNames()
                                                            .withVmName("dbvmsr")
                                                            .withHostname("dbsrhostName")
                                                            .withNetworkInterfaces(
                                                                Arrays
                                                                    .asList(
                                                                        new NetworkInterfaceResourceNames()
                                                                            .withNetworkInterfaceName("dbsrnic")))
                                                            .withOsDiskName("dbsrosdisk")
                                                            .withDataDiskNames(
                                                                mapOf(
                                                                    "hanaData",
                                                                    Arrays.asList("hanadatasr0", "hanadatasr1"),
                                                                    "hanaLog",
                                                                    Arrays
                                                                        .asList(
                                                                            "hanalogsr0", "hanalogsr1", "hanalogsr2"),
                                                                    "hanaShared",
                                                                    Arrays.asList("hanasharedsr0", "hanasharedsr1"),
                                                                    "usrSap",
                                                                    Arrays.asList("usrsapsr0")))))
                                            .withLoadBalancer(
                                                new LoadBalancerResourceNames()
                                                    .withLoadBalancerName("dblb")
                                                    .withFrontendIpConfigurationNames(Arrays.asList("dbip"))
                                                    .withBackendPoolNames(Arrays.asList("dbBackendPool"))
                                                    .withHealthProbeNames(Arrays.asList("dbHealthProbe"))))
                                    .withSharedStorage(
                                        new SharedStorageResourceNames()
                                            .withSharedStorageAccountName("storageacc")
                                            .withSharedStorageAccountPrivateEndPointName("peForxNFS"))))
                    .withOsSapConfiguration(new OsSapConfiguration().withSapFqdn("xyz.test.com")))
            .withTags(mapOf())
            .create();
    }

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
