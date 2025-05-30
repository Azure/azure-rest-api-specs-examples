
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DeploymentWithOSConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DiskConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DiskSku;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DiskSkuName;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DiskVolumeConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ImageReference;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.LinuxConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.NetworkConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.OSProfile;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.OsSapConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapVirtualInstanceProperties;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SingleServerConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SshKeyPair;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.VirtualMachineConfiguration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapVirtualInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateInfraDiskOsSingleServerRecommended.json
     */
    /**
     * Sample code: Create Infrastructure with Disk and OS configurations for Single Server System (Recommended).
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void createInfrastructureWithDiskAndOSConfigurationsForSingleServerSystemRecommended(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().define("X00").withRegion("westcentralus").withExistingResourceGroup("test-rg")
            .withTags(mapOf())
            .withProperties(new SapVirtualInstanceProperties().withEnvironment(SapEnvironmentType.NON_PROD)
                .withSapProduct(SapProductType.S4HANA)
                .withConfiguration(new DeploymentWithOSConfiguration().withAppLocation("eastus")
                    .withInfrastructureConfiguration(new SingleServerConfiguration().withAppResourceGroup("X00-RG")
                        .withNetworkConfiguration(new NetworkConfiguration().withIsSecondaryIpEnabled(true))
                        .withDatabaseType(SapDatabaseType.HANA)
                        .withSubnetId(
                            "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet")
                        .withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                            .withVmSize("Standard_E32ds_v4")
                            .withImageReference(new ImageReference().withPublisher("RedHat").withOffer("RHEL-SAP")
                                .withSku("84sapha-gen2").withVersion("latest"))
                            .withOsProfile(new OSProfile().withAdminUsername("{your-username}")
                                .withOsConfiguration(new LinuxConfiguration().withDisablePasswordAuthentication(true)
                                    .withSshKeyPair(new SshKeyPair().withPublicKey("fakeTokenPlaceholder")
                                        .withPrivateKey("fakeTokenPlaceholder")))))
                        .withDbDiskConfiguration(new DiskConfiguration().withDiskVolumeConfigurations(mapOf("backup",
                            new DiskVolumeConfiguration().withCount(2L).withSizeGB(256L)
                                .withSku(new DiskSku().withName(DiskSkuName.STANDARD_SSD_LRS)),
                            "hana/data",
                            new DiskVolumeConfiguration().withCount(4L).withSizeGB(128L)
                                .withSku(new DiskSku().withName(DiskSkuName.PREMIUM_LRS)),
                            "hana/log",
                            new DiskVolumeConfiguration().withCount(3L).withSizeGB(128L)
                                .withSku(new DiskSku().withName(DiskSkuName.PREMIUM_LRS)),
                            "hana/shared",
                            new DiskVolumeConfiguration().withCount(1L).withSizeGB(256L)
                                .withSku(new DiskSku().withName(DiskSkuName.STANDARD_SSD_LRS)),
                            "os",
                            new DiskVolumeConfiguration().withCount(1L).withSizeGB(64L)
                                .withSku(new DiskSku().withName(DiskSkuName.STANDARD_SSD_LRS)),
                            "usr/sap",
                            new DiskVolumeConfiguration().withCount(1L).withSizeGB(128L)
                                .withSku(new DiskSku().withName(DiskSkuName.PREMIUM_LRS))))))
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
