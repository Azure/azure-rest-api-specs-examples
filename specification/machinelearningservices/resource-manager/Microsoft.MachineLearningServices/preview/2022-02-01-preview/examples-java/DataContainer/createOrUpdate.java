import com.azure.resourcemanager.machinelearning.models.DataContainerDetails;
import com.azure.resourcemanager.machinelearning.models.DataType;
import java.util.HashMap;
import java.util.Map;

/** Samples for DataContainers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/DataContainer/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Data Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdateDataContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .dataContainers()
            .define("datacontainer123")
            .withExistingWorkspace("testrg123", "workspace123")
            .withProperties(
                new DataContainerDetails()
                    .withDescription("string")
                    .withProperties(mapOf("properties1", "value1", "properties2", "value2"))
                    .withTags(mapOf("tag1", "value1", "tag2", "value2"))
                    .withDataType(DataType.URI_FILE))
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
