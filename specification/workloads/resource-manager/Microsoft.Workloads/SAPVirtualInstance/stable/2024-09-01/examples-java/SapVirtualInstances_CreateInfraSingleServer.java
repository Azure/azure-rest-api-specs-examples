
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DeploymentConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ImageReference;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.LinuxConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.NetworkConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.OSProfile;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapVirtualInstanceProperties;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SingleServerConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SshConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SshPublicKey;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.VirtualMachineConfiguration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapVirtualInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateInfraSingleServer.json
     */
    /**
     * Sample code: Create Infrastructure only for Single Server System.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void createInfrastructureOnlyForSingleServerSystem(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().define("X00").withRegion("westcentralus").withExistingResourceGroup("test-rg")
            .withTags(mapOf())
            .withProperties(new SapVirtualInstanceProperties().withEnvironment(SapEnvironmentType.NON_PROD)
                .withSapProduct(SapProductType.S4HANA)
                .withConfiguration(new DeploymentConfiguration().withAppLocation("eastus")
                    .withInfrastructureConfiguration(new SingleServerConfiguration().withAppResourceGroup("X00-RG")
                        .withNetworkConfiguration(new NetworkConfiguration().withIsSecondaryIpEnabled(true))
                        .withDatabaseType(SapDatabaseType.HANA)
                        .withSubnetId(
                            "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                        .withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                            .withVmSize("Standard_E32ds_v4")
                            .withImageReference(new ImageReference().withPublisher("RedHat").withOffer("RHEL-SAP")
                                .withSku("84sapha-gen2").withVersion("latest"))
                            .withOsProfile(new OSProfile().withAdminUsername("{your-username}")
                                .withOsConfiguration(new LinuxConfiguration().withDisablePasswordAuthentication(true)
                                    .withSsh(new SshConfiguration().withPublicKeys(
                                        Arrays.asList(new SshPublicKey().withKeyData("fakeTokenPlaceholder"))))))))))
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
