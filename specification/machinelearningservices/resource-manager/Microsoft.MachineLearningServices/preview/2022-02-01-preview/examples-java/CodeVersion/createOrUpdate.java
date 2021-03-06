import com.azure.resourcemanager.machinelearning.models.CodeVersionDetails;
import java.util.HashMap;
import java.util.Map;

/** Samples for CodeVersions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/CodeVersion/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Code Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdateCodeVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .codeVersions()
            .define("string")
            .withExistingCode("test-rg", "my-aml-workspace", "string")
            .withProperties(
                new CodeVersionDetails()
                    .withDescription("string")
                    .withProperties(mapOf("string", "string"))
                    .withTags(mapOf("string", "string"))
                    .withIsAnonymous(false)
                    .withCodeUri("https://blobStorage/folderName"))
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
