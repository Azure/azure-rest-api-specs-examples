import com.azure.resourcemanager.machinelearning.models.EnvironmentContainerProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for EnvironmentContainers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/EnvironmentContainer/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Environment Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateEnvironmentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .environmentContainers()
            .define("testEnvironment")
            .withExistingWorkspace("testrg123", "testworkspace")
            .withProperties(
                new EnvironmentContainerProperties()
                    .withDescription("string")
                    .withProperties(
                        mapOf("additionalProp1", "string", "additionalProp2", "string", "additionalProp3", "string"))
                    .withTags(
                        mapOf("additionalProp1", "string", "additionalProp2", "string", "additionalProp3", "string")))
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
