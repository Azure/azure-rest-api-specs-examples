
import com.azure.resourcemanager.machinelearning.models.FeaturesetContainerProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for FeaturesetContainers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturesetContainer/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Featureset Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceFeaturesetContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featuresetContainers().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(new FeaturesetContainerProperties().withDescription("string")
                .withTags(mapOf("string", "string")).withProperties(mapOf("string", "string")).withIsArchived(false))
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
