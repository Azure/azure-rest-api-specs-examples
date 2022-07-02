import com.azure.resourcemanager.workloads.models.DeploymentConfiguration;
import com.azure.resourcemanager.workloads.models.ImageReference;
import com.azure.resourcemanager.workloads.models.LinuxConfiguration;
import com.azure.resourcemanager.workloads.models.NetworkConfiguration;
import com.azure.resourcemanager.workloads.models.OSProfile;
import com.azure.resourcemanager.workloads.models.SapDatabaseType;
import com.azure.resourcemanager.workloads.models.SapEnvironmentType;
import com.azure.resourcemanager.workloads.models.SapProductType;
import com.azure.resourcemanager.workloads.models.SingleServerConfiguration;
import com.azure.resourcemanager.workloads.models.SshConfiguration;
import com.azure.resourcemanager.workloads.models.SshPublicKey;
import com.azure.resourcemanager.workloads.models.VirtualMachineConfiguration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for SapVirtualInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Create_SingleServer.json
     */
    /**
     * Sample code: SAPVirtualInstances_Create_SingleServer.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesCreateSingleServer(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapVirtualInstances()
            .define("X00")
            .withRegion("westcentralus")
            .withExistingResourceGroup("test-rg")
            .withEnvironment(SapEnvironmentType.NON_PROD)
            .withSapProduct(SapProductType.S4HANA)
            .withConfiguration(
                new DeploymentConfiguration()
                    .withAppLocation("eastus")
                    .withInfrastructureConfiguration(
                        new SingleServerConfiguration()
                            .withAppResourceGroup("X00-RG")
                            .withNetworkConfiguration(new NetworkConfiguration().withIsSecondaryIpEnabled(true))
                            .withDatabaseType(SapDatabaseType.HANA)
                            .withSubnetId(
                                "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                            .withVirtualMachineConfiguration(
                                new VirtualMachineConfiguration()
                                    .withVmSize("Standard_E32ds_v4")
                                    .withImageReference(
                                        new ImageReference()
                                            .withPublisher("RedHat")
                                            .withOffer("RHEL-SAP")
                                            .withSku("7.4")
                                            .withVersion("7.4.2019062505"))
                                    .withOsProfile(
                                        new OSProfile()
                                            .withAdminUsername("{your-username}")
                                            .withOsConfiguration(
                                                new LinuxConfiguration()
                                                    .withDisablePasswordAuthentication(true)
                                                    .withSsh(
                                                        new SshConfiguration()
                                                            .withPublicKeys(
                                                                Arrays
                                                                    .asList(
                                                                        new SshPublicKey()
                                                                            .withKeyData("ssh-rsa public key")))))))))
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
