
import com.azure.resourcemanager.machinelearning.models.CodeVersionProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CodeVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/CodeVersion/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        createOrUpdateWorkspaceCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeVersions().define("string").withExistingCode("test-rg", "my-aml-workspace", "string")
            .withProperties(new CodeVersionProperties().withDescription("string").withTags(mapOf("string", "string"))
                .withProperties(mapOf("string", "string")).withIsAnonymous(false).withCodeUri("fakeTokenPlaceholder"))
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
