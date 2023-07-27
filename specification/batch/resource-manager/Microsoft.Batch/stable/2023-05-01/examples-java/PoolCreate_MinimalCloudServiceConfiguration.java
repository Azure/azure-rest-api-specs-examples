import com.azure.resourcemanager.batch.models.CloudServiceConfiguration;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pool Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolCreate_MinimalCloudServiceConfiguration.json
     */
    /**
     * Sample code: CreatePool - Minimal CloudServiceConfiguration.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolMinimalCloudServiceConfiguration(
        com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .define("testpool")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withVmSize("STANDARD_D4")
            .withDeploymentConfiguration(
                new DeploymentConfiguration()
                    .withCloudServiceConfiguration(new CloudServiceConfiguration().withOsFamily("5")))
            .withScaleSettings(new ScaleSettings().withFixedScale(new FixedScaleSettings().withTargetDedicatedNodes(3)))
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
