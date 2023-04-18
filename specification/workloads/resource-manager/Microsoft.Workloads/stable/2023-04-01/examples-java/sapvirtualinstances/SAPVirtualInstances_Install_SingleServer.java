import com.azure.resourcemanager.workloads.models.DeploymentWithOSConfiguration;
import com.azure.resourcemanager.workloads.models.ImageReference;
import com.azure.resourcemanager.workloads.models.LinuxConfiguration;
import com.azure.resourcemanager.workloads.models.OSProfile;
import com.azure.resourcemanager.workloads.models.OsSapConfiguration;
import com.azure.resourcemanager.workloads.models.SapEnvironmentType;
import com.azure.resourcemanager.workloads.models.SapInstallWithoutOSConfigSoftwareConfiguration;
import com.azure.resourcemanager.workloads.models.SapProductType;
import com.azure.resourcemanager.workloads.models.SingleServerConfiguration;
import com.azure.resourcemanager.workloads.models.SshKeyPair;
import com.azure.resourcemanager.workloads.models.VirtualMachineConfiguration;
import java.util.HashMap;
import java.util.Map;

/** Samples for SapVirtualInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Install_SingleServer.json
     */
    /**
     * Sample code: Install SAP Software on Single Server System.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void installSAPSoftwareOnSingleServerSystem(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapVirtualInstances()
            .define("X00")
            .withRegion("eastus2")
            .withExistingResourceGroup("test-rg")
            .withEnvironment(SapEnvironmentType.NON_PROD)
            .withSapProduct(SapProductType.S4HANA)
            .withConfiguration(
                new DeploymentWithOSConfiguration()
                    .withAppLocation("eastus")
                    .withInfrastructureConfiguration(
                        new SingleServerConfiguration()
                            .withAppResourceGroup("test-rg")
                            .withSubnetId(
                                "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/testsubnet")
                            .withVirtualMachineConfiguration(
                                new VirtualMachineConfiguration()
                                    .withVmSize("Standard_E32ds_v4")
                                    .withImageReference(
                                        new ImageReference()
                                            .withPublisher("SUSE")
                                            .withOffer("SLES-SAP")
                                            .withSku("12-sp4-gen2")
                                            .withVersion("2022.02.01"))
                                    .withOsProfile(
                                        new OSProfile()
                                            .withAdminUsername("azureappadmin")
                                            .withOsConfiguration(
                                                new LinuxConfiguration()
                                                    .withDisablePasswordAuthentication(true)
                                                    .withSshKeyPair(
                                                        new SshKeyPair()
                                                            .withPublicKey("fakeTokenPlaceholder")
                                                            .withPrivateKey("fakeTokenPlaceholder"))))))
                    .withSoftwareConfiguration(
                        new SapInstallWithoutOSConfigSoftwareConfiguration()
                            .withBomUrl(
                                "https://teststorageaccount.blob.core.windows.net/sapbits/sapfiles/boms/S41909SPS03_v0011ms/S41909SPS03_v0011ms.yaml")
                            .withSapBitsStorageAccountId(
                                "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/teststorageaccount")
                            .withSoftwareVersion("SAP S/4HANA 1909 SPS 03"))
                    .withOsSapConfiguration(new OsSapConfiguration().withSapFqdn("sap.bpaas.com")))
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
