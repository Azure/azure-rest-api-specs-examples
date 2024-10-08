
import com.azure.resourcemanager.machinelearning.models.FlavorData;
import com.azure.resourcemanager.machinelearning.models.ModelVersionProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ModelVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ModelVersion/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Model Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        createOrUpdateWorkspaceModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelVersions().define("string").withExistingModel("test-rg", "my-aml-workspace", "string")
            .withProperties(new ModelVersionProperties().withDescription("string").withTags(mapOf("string", "string"))
                .withProperties(mapOf("string", "string")).withIsAnonymous(false)
                .withFlavors(mapOf("string", new FlavorData().withData(mapOf("string", "string"))))
                .withModelType("CustomModel").withModelUri("string"))
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
