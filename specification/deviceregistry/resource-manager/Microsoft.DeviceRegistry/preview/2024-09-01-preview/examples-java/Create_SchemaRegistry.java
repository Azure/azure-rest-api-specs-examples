
import com.azure.resourcemanager.deviceregistry.models.SchemaRegistryProperties;
import com.azure.resourcemanager.deviceregistry.models.SystemAssignedServiceIdentity;
import com.azure.resourcemanager.deviceregistry.models.SystemAssignedServiceIdentityType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SchemaRegistries CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Create_SchemaRegistry.json
     */
    /**
     * Sample code: Create_SchemaRegistry.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void createSchemaRegistry(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaRegistries().define("my-schema-registry").withRegion("West Europe")
            .withExistingResourceGroup("myResourceGroup").withTags(mapOf())
            .withProperties(new SchemaRegistryProperties().withNamespace("sr-namespace-001")
                .withDisplayName("Schema Registry namespace 001").withDescription("This is a sample Schema Registry")
                .withStorageAccountContainerUrl("my-blob-storage.blob.core.windows.net/my-container"))
            .withIdentity(new SystemAssignedServiceIdentity().withType(SystemAssignedServiceIdentityType.NONE))
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
