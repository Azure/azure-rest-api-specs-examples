
import com.azure.resourcemanager.machinelearning.models.FeaturestoreEntityContainerProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for FeaturestoreEntityContainers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturestoreEntityContainer/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Featurestore Entity Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceFeaturestoreEntityContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featurestoreEntityContainers().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(new FeaturestoreEntityContainerProperties().withDescription("string")
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
