
import com.azure.resourcemanager.machinelearning.fluent.models.DataVersionBaseInner;
import com.azure.resourcemanager.machinelearning.models.MLTableData;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RegistryDataVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/DataVersionBase/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Registry Data Version Base.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateRegistryDataVersionBase(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataVersions().createOrUpdate("test-rg", "registryName", "string", "string",
            new DataVersionBaseInner().withProperties(new MLTableData().withDescription("string")
                .withTags(mapOf("string", "string")).withProperties(mapOf("string", "string")).withIsArchived(false)
                .withIsAnonymous(false).withDataUri("string").withReferencedUris(Arrays.asList("string"))),
            com.azure.core.util.Context.NONE);
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
