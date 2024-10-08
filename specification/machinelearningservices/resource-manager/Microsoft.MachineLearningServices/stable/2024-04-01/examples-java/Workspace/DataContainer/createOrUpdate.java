
import com.azure.resourcemanager.machinelearning.models.DataContainerProperties;
import com.azure.resourcemanager.machinelearning.models.DataType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DataContainers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/DataContainer/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Data Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        createOrUpdateWorkspaceDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.dataContainers().define("datacontainer123").withExistingWorkspace("testrg123", "workspace123")
            .withProperties(new DataContainerProperties().withDescription("string")
                .withTags(mapOf("tag1", "value1", "tag2", "value2"))
                .withProperties(mapOf("properties1", "value1", "properties2", "value2"))
                .withDataType(DataType.fromString("UriFile")))
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
