import com.azure.resourcemanager.batch.models.AutoScaleSettings;
import com.azure.resourcemanager.batch.models.CachingType;
import com.azure.resourcemanager.batch.models.DataDisk;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.DiffDiskPlacement;
import com.azure.resourcemanager.batch.models.DiffDiskSettings;
import com.azure.resourcemanager.batch.models.DiskEncryptionConfiguration;
import com.azure.resourcemanager.batch.models.DiskEncryptionTarget;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.InboundEndpointProtocol;
import com.azure.resourcemanager.batch.models.InboundNatPool;
import com.azure.resourcemanager.batch.models.NetworkConfiguration;
import com.azure.resourcemanager.batch.models.NetworkSecurityGroupRule;
import com.azure.resourcemanager.batch.models.NetworkSecurityGroupRuleAccess;
import com.azure.resourcemanager.batch.models.NodePlacementConfiguration;
import com.azure.resourcemanager.batch.models.NodePlacementPolicyType;
import com.azure.resourcemanager.batch.models.OSDisk;
import com.azure.resourcemanager.batch.models.PoolEndpointConfiguration;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.StorageAccountType;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import com.azure.resourcemanager.batch.models.WindowsConfiguration;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pool Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolCreate_VirtualMachineConfiguration.json
     */
    /**
     * Sample code: CreatePool - Full VirtualMachineConfiguration.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolFullVirtualMachineConfiguration(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .define("testpool")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withVmSize("STANDARD_D4")
            .withDeploymentConfiguration(
                new DeploymentConfiguration()
                    .withVirtualMachineConfiguration(
                        new VirtualMachineConfiguration()
                            .withImageReference(
                                new ImageReference()
                                    .withPublisher("MicrosoftWindowsServer")
                                    .withOffer("WindowsServer")
                                    .withSku("2016-Datacenter-SmallDisk")
                                    .withVersion("latest"))
                            .withNodeAgentSkuId("batch.node.windows amd64")
                            .withWindowsConfiguration(new WindowsConfiguration().withEnableAutomaticUpdates(false))
                            .withDataDisks(
                                Arrays
                                    .asList(
                                        new DataDisk()
                                            .withLun(0)
                                            .withCaching(CachingType.READ_WRITE)
                                            .withDiskSizeGB(30)
                                            .withStorageAccountType(StorageAccountType.PREMIUM_LRS),
                                        new DataDisk()
                                            .withLun(1)
                                            .withCaching(CachingType.NONE)
                                            .withDiskSizeGB(200)
                                            .withStorageAccountType(StorageAccountType.STANDARD_LRS)))
                            .withLicenseType("Windows_Server")
                            .withDiskEncryptionConfiguration(
                                new DiskEncryptionConfiguration()
                                    .withTargets(
                                        Arrays
                                            .asList(DiskEncryptionTarget.OS_DISK, DiskEncryptionTarget.TEMPORARY_DISK)))
                            .withNodePlacementConfiguration(
                                new NodePlacementConfiguration().withPolicy(NodePlacementPolicyType.ZONAL))
                            .withOsDisk(
                                new OSDisk()
                                    .withEphemeralOSDiskSettings(
                                        new DiffDiskSettings().withPlacement(DiffDiskPlacement.CACHE_DISK)))))
            .withScaleSettings(
                new ScaleSettings()
                    .withAutoScale(
                        new AutoScaleSettings()
                            .withFormula("$TargetDedicatedNodes=1")
                            .withEvaluationInterval(Duration.parse("PT5M"))))
            .withNetworkConfiguration(
                new NetworkConfiguration()
                    .withEndpointConfiguration(
                        new PoolEndpointConfiguration()
                            .withInboundNatPools(
                                Arrays
                                    .asList(
                                        new InboundNatPool()
                                            .withName("testnat")
                                            .withProtocol(InboundEndpointProtocol.TCP)
                                            .withBackendPort(12001)
                                            .withFrontendPortRangeStart(15000)
                                            .withFrontendPortRangeEnd(15100)
                                            .withNetworkSecurityGroupRules(
                                                Arrays
                                                    .asList(
                                                        new NetworkSecurityGroupRule()
                                                            .withPriority(150)
                                                            .withAccess(NetworkSecurityGroupRuleAccess.ALLOW)
                                                            .withSourceAddressPrefix("192.100.12.45")
                                                            .withSourcePortRanges(Arrays.asList("1", "2")),
                                                        new NetworkSecurityGroupRule()
                                                            .withPriority(3500)
                                                            .withAccess(NetworkSecurityGroupRuleAccess.DENY)
                                                            .withSourceAddressPrefix("*")
                                                            .withSourcePortRanges(Arrays.asList("*"))))))))
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
