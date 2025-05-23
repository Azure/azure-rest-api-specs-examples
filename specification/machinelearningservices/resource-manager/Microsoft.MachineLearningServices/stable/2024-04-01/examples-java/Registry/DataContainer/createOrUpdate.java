
import com.azure.resourcemanager.machinelearning.fluent.models.DataContainerInner;
import com.azure.resourcemanager.machinelearning.models.DataContainerProperties;
import com.azure.resourcemanager.machinelearning.models.DataType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RegistryDataContainers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/DataContainer/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Registry Data Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        createOrUpdateRegistryDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataContainers()
            .createOrUpdate("test-rg", "registryName", "string",
                new DataContainerInner().withProperties(new DataContainerProperties().withDescription("string")
                    .withTags(mapOf("string", "string")).withProperties(mapOf("string", "string")).withIsArchived(false)
                    .withDataType(DataType.URI_FOLDER)),
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
