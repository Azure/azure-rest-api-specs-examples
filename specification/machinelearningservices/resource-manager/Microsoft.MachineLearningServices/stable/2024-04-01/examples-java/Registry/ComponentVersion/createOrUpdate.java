
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.fluent.models.ComponentVersionInner;
import com.azure.resourcemanager.machinelearning.models.ComponentVersionProperties;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RegistryComponentVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ComponentVersion/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Registry Component Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateRegistryComponentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) throws IOException {
        manager.registryComponentVersions().createOrUpdate("test-rg", "my-aml-registry", "string", "string",
            new ComponentVersionInner().withProperties(new ComponentVersionProperties().withDescription("string")
                .withTags(mapOf("string", "string")).withProperties(mapOf("string", "string")).withIsAnonymous(false)
                .withComponentSpec(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                    "{\"8ced901b-d826-477d-bfef-329da9672513\":null}", Object.class, SerializerEncoding.JSON))),
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
