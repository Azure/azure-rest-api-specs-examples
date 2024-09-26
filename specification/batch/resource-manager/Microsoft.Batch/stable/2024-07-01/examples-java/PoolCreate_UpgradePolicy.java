
import com.azure.resourcemanager.batch.models.AutomaticOSUpgradePolicy;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.NodePlacementConfiguration;
import com.azure.resourcemanager.batch.models.NodePlacementPolicyType;
import com.azure.resourcemanager.batch.models.RollingUpgradePolicy;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.UpgradeMode;
import com.azure.resourcemanager.batch.models.UpgradePolicy;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import com.azure.resourcemanager.batch.models.WindowsConfiguration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Pool Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolCreate_UpgradePolicy.json
     */
    /**
     * Sample code: CreatePool - UpgradePolicy.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolUpgradePolicy(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().define("testpool").withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withVmSize("Standard_d4s_v3")
            .withDeploymentConfiguration(
                new DeploymentConfiguration().withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                    .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                        .withOffer("WindowsServer").withSku("2019-datacenter-smalldisk").withVersion("latest"))
                    .withNodeAgentSkuId("batch.node.windows amd64")
                    .withWindowsConfiguration(new WindowsConfiguration().withEnableAutomaticUpdates(false))
                    .withNodePlacementConfiguration(
                        new NodePlacementConfiguration().withPolicy(NodePlacementPolicyType.ZONAL))))
            .withScaleSettings(new ScaleSettings()
                .withFixedScale(new FixedScaleSettings().withTargetDedicatedNodes(2).withTargetLowPriorityNodes(0)))
            .withUpgradePolicy(
                new UpgradePolicy().withMode(UpgradeMode.AUTOMATIC)
                    .withAutomaticOSUpgradePolicy(new AutomaticOSUpgradePolicy().withDisableAutomaticRollback(true)
                        .withEnableAutomaticOSUpgrade(true).withUseRollingUpgradePolicy(true)
                        .withOsRollingUpgradeDeferral(true))
                    .withRollingUpgradePolicy(new RollingUpgradePolicy().withEnableCrossZoneUpgrade(true)
                        .withMaxBatchInstancePercent(20).withMaxUnhealthyInstancePercent(20)
                        .withMaxUnhealthyUpgradedInstancePercent(20).withPauseTimeBetweenBatches("PT0S")
                        .withPrioritizeUnhealthyInstances(false).withRollbackFailedInstancesOnPolicyBreach(false)))
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
