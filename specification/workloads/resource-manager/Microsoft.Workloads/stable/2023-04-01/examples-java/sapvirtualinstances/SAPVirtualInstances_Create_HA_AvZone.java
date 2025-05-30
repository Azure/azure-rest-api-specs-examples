
import com.azure.resourcemanager.workloads.models.ApplicationServerConfiguration;
import com.azure.resourcemanager.workloads.models.CentralServerConfiguration;
import com.azure.resourcemanager.workloads.models.DatabaseConfiguration;
import com.azure.resourcemanager.workloads.models.DeploymentConfiguration;
import com.azure.resourcemanager.workloads.models.HighAvailabilityConfiguration;
import com.azure.resourcemanager.workloads.models.ImageReference;
import com.azure.resourcemanager.workloads.models.LinuxConfiguration;
import com.azure.resourcemanager.workloads.models.OSProfile;
import com.azure.resourcemanager.workloads.models.SapDatabaseType;
import com.azure.resourcemanager.workloads.models.SapEnvironmentType;
import com.azure.resourcemanager.workloads.models.SapHighAvailabilityType;
import com.azure.resourcemanager.workloads.models.SapProductType;
import com.azure.resourcemanager.workloads.models.SshConfiguration;
import com.azure.resourcemanager.workloads.models.SshPublicKey;
import com.azure.resourcemanager.workloads.models.ThreeTierConfiguration;
import com.azure.resourcemanager.workloads.models.VirtualMachineConfiguration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapVirtualInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/
     * SAPVirtualInstances_Create_HA_AvZone.json
     */
    /**
     * Sample code: Create Infrastructure only for HA System with Availability Zone.
     * 
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createInfrastructureOnlyForHASystemWithAvailabilityZone(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapVirtualInstances().define("X00").withRegion("westcentralus").withExistingResourceGroup("test-rg")
            .withEnvironment(SapEnvironmentType.PROD).withSapProduct(SapProductType.S4HANA)
            .withConfiguration(new DeploymentConfiguration().withAppLocation("eastus")
                .withInfrastructureConfiguration(new ThreeTierConfiguration().withAppResourceGroup("X00-RG")
                    .withCentralServer(new CentralServerConfiguration().withSubnetId(
                        "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                        .withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                            .withVmSize("Standard_E16ds_v4")
                            .withImageReference(new ImageReference().withPublisher("RedHat").withOffer("RHEL-SAP")
                                .withSku("84sapha-gen2").withVersion("latest"))
                            .withOsProfile(new OSProfile().withAdminUsername("{your-username}")
                                .withOsConfiguration(new LinuxConfiguration().withDisablePasswordAuthentication(true)
                                    .withSsh(new SshConfiguration().withPublicKeys(
                                        Arrays.asList(new SshPublicKey().withKeyData("fakeTokenPlaceholder")))))))
                        .withInstanceCount(2L))
                    .withApplicationServer(new ApplicationServerConfiguration().withSubnetId(
                        "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                        .withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                            .withVmSize("Standard_E32ds_v4")
                            .withImageReference(new ImageReference().withPublisher("RedHat").withOffer("RHEL-SAP")
                                .withSku("84sapha-gen2").withVersion("latest"))
                            .withOsProfile(new OSProfile().withAdminUsername("{your-username}")
                                .withOsConfiguration(new LinuxConfiguration().withDisablePasswordAuthentication(true)
                                    .withSsh(new SshConfiguration().withPublicKeys(
                                        Arrays.asList(new SshPublicKey().withKeyData("fakeTokenPlaceholder")))))))
                        .withInstanceCount(6L))
                    .withDatabaseServer(new DatabaseConfiguration().withDatabaseType(SapDatabaseType.HANA).withSubnetId(
                        "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet")
                        .withVirtualMachineConfiguration(new VirtualMachineConfiguration().withVmSize("Standard_M32ts")
                            .withImageReference(new ImageReference().withPublisher("RedHat").withOffer("RHEL-SAP")
                                .withSku("84sapha-gen2").withVersion("latest"))
                            .withOsProfile(new OSProfile().withAdminUsername("{your-username}")
                                .withOsConfiguration(new LinuxConfiguration().withDisablePasswordAuthentication(true)
                                    .withSsh(new SshConfiguration().withPublicKeys(
                                        Arrays.asList(new SshPublicKey().withKeyData("fakeTokenPlaceholder")))))))
                        .withInstanceCount(2L))
                    .withHighAvailabilityConfig(new HighAvailabilityConfiguration()
                        .withHighAvailabilityType(SapHighAvailabilityType.AVAILABILITY_ZONE))))
            .withTags(mapOf()).create();
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
