
import com.azure.resourcemanager.machinelearning.models.ModelContainerProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ModelContainers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ModelContainer/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Model Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceModelContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelContainers().define("testContainer").withExistingWorkspace("testrg123", "workspace123")
            .withProperties(new ModelContainerProperties().withDescription("Model container description")
                .withTags(mapOf("tag1", "value1", "tag2", "value2")))
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
