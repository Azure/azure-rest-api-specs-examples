
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Pool Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolCreate_Tags.json
     */
    /**
     * Sample code: CreatePool - Tags.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolTags(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().define("testpool").withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withTags(mapOf("TagName1", "TagValue1", "TagName2", "TagValue2")).withVmSize("Standard_d4s_v3")
            .withDeploymentConfiguration(
                new DeploymentConfiguration().withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                    .withImageReference(new ImageReference().withPublisher("Canonical")
                        .withOffer("0001-com-ubuntu-server-jammy").withSku("22_04-lts").withVersion("latest"))
                    .withNodeAgentSkuId("batch.node.ubuntu 22.04")))
            .withScaleSettings(new ScaleSettings()
                .withFixedScale(new FixedScaleSettings().withTargetDedicatedNodes(1).withTargetLowPriorityNodes(0)))
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
