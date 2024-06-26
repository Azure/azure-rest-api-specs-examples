import com.azure.resourcemanager.machinelearning.models.UriFileDataVersion;
import java.util.HashMap;
import java.util.Map;

/** Samples for DataVersions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/DataVersionBase/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Data Version Base.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateDataVersionBase(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .dataVersions()
            .define("string")
            .withExistingData("test-rg", "my-aml-workspace", "string")
            .withProperties(
                new UriFileDataVersion()
                    .withDescription("string")
                    .withProperties(mapOf("string", "string"))
                    .withTags(mapOf("string", "string"))
                    .withIsAnonymous(false)
                    .withDataUri("string"))
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
