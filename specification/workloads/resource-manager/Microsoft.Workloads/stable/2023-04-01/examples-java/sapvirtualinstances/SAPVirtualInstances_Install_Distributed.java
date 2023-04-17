import com.azure.resourcemanager.workloads.models.ApplicationServerConfiguration;
import com.azure.resourcemanager.workloads.models.CentralServerConfiguration;
import com.azure.resourcemanager.workloads.models.DatabaseConfiguration;
import com.azure.resourcemanager.workloads.models.DeploymentWithOSConfiguration;
import com.azure.resourcemanager.workloads.models.ImageReference;
import com.azure.resourcemanager.workloads.models.LinuxConfiguration;
import com.azure.resourcemanager.workloads.models.NetworkConfiguration;
import com.azure.resourcemanager.workloads.models.OSProfile;
import com.azure.resourcemanager.workloads.models.OsSapConfiguration;
import com.azure.resourcemanager.workloads.models.SapEnvironmentType;
import com.azure.resourcemanager.workloads.models.SapInstallWithoutOSConfigSoftwareConfiguration;
import com.azure.resourcemanager.workloads.models.SapProductType;
import com.azure.resourcemanager.workloads.models.SshKeyPair;
import com.azure.resourcemanager.workloads.models.ThreeTierConfiguration;
import com.azure.resourcemanager.workloads.models.VirtualMachineConfiguration;
import java.util.HashMap;
import java.util.Map;

/** Samples for SapVirtualInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Install_Distributed.json
     */
    /**
     * Sample code: Install SAP Software on Distributed System.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void installSAPSoftwareOnDistributedSystem(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapVirtualInstances()
            .define("X00")
            .withRegion("eastus2")
            .withExistingResourceGroup("test-rg")
            .withEnvironment(SapEnvironmentType.PROD)
            .withSapProduct(SapProductType.S4HANA)
            .withConfiguration(
                new DeploymentWithOSConfiguration()
                    .withAppLocation("eastus")
                    .withInfrastructureConfiguration(
                        new ThreeTierConfiguration()
                            .withAppResourceGroup("{{resourcegrp}}")
                            .withNetworkConfiguration(new NetworkConfiguration().withIsSecondaryIpEnabled(true))
                            .withCentralServer(
                                new CentralServerConfiguration()
                                    .withSubnetId(
                                        "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app")
                                    .withVirtualMachineConfiguration(
                                        new VirtualMachineConfiguration()
                                            .withVmSize("Standard_E4ds_v4")
                                            .withImageReference(
                                                new ImageReference()
                                                    .withPublisher("RedHat")
                                                    .withOffer("RHEL-SAP-HA")
                                                    .withSku("8.2")
                                                    .withVersion("8.2.2021091201"))
                                            .withOsProfile(
                                                new OSProfile()
                                                    .withAdminUsername("azureuser")
                                                    .withOsConfiguration(
                                                        new LinuxConfiguration()
                                                            .withDisablePasswordAuthentication(true)
                                                            .withSshKeyPair(
                                                                new SshKeyPair()
                                                                    .withPublicKey("fakeTokenPlaceholder")
                                                                    .withPrivateKey("fakeTokenPlaceholder")))))
                                    .withInstanceCount(1L))
                            .withApplicationServer(
                                new ApplicationServerConfiguration()
                                    .withSubnetId(
                                        "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app")
                                    .withVirtualMachineConfiguration(
                                        new VirtualMachineConfiguration()
                                            .withVmSize("Standard_E4ds_v4")
                                            .withImageReference(
                                                new ImageReference()
                                                    .withPublisher("RedHat")
                                                    .withOffer("RHEL-SAP-HA")
                                                    .withSku("8.2")
                                                    .withVersion("8.2.2021091201"))
                                            .withOsProfile(
                                                new OSProfile()
                                                    .withAdminUsername("azureuser")
                                                    .withOsConfiguration(
                                                        new LinuxConfiguration()
                                                            .withDisablePasswordAuthentication(true)
                                                            .withSshKeyPair(
                                                                new SshKeyPair()
                                                                    .withPublicKey("fakeTokenPlaceholder")
                                                                    .withPrivateKey("fakeTokenPlaceholder")))))
                                    .withInstanceCount(2L))
                            .withDatabaseServer(
                                new DatabaseConfiguration()
                                    .withSubnetId(
                                        "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app")
                                    .withVirtualMachineConfiguration(
                                        new VirtualMachineConfiguration()
                                            .withVmSize("Standard_M32ts")
                                            .withImageReference(
                                                new ImageReference()
                                                    .withPublisher("RedHat")
                                                    .withOffer("RHEL-SAP-HA")
                                                    .withSku("8.2")
                                                    .withVersion("8.2.2021091201"))
                                            .withOsProfile(
                                                new OSProfile()
                                                    .withAdminUsername("azureuser")
                                                    .withOsConfiguration(
                                                        new LinuxConfiguration()
                                                            .withDisablePasswordAuthentication(true)
                                                            .withSshKeyPair(
                                                                new SshKeyPair()
                                                                    .withPublicKey("fakeTokenPlaceholder")
                                                                    .withPrivateKey("fakeTokenPlaceholder")))))
                                    .withInstanceCount(1L)))
                    .withSoftwareConfiguration(
                        new SapInstallWithoutOSConfigSoftwareConfiguration()
                            .withBomUrl(
                                "https://teststorageaccount.blob.core.windows.net/sapbits/sapfiles/boms/S41909SPS03_v0011ms/S41909SPS03_v0011ms.yaml")
                            .withSapBitsStorageAccountId(
                                "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/teststorageaccount")
                            .withSoftwareVersion("SAP S/4HANA 1909 SPS 03"))
                    .withOsSapConfiguration(new OsSapConfiguration().withSapFqdn("sap.bpaas.com")))
            .withTags(mapOf("created by", "azureuser"))
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
