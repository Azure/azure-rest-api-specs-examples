import com.azure.resourcemanager.machinelearning.models.ComponentContainerProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for ComponentContainers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ComponentContainer/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Component Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateComponentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .componentContainers()
            .define("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new ComponentContainerProperties()
                    .withDescription("string")
                    .withProperties(mapOf("string", "string"))
                    .withTags(mapOf("string", "string")))
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
