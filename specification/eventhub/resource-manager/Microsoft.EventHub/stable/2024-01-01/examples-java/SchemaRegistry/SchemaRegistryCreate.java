
import com.azure.resourcemanager.eventhubs.fluent.models.SchemaGroupInner;
import com.azure.resourcemanager.eventhubs.models.SchemaCompatibility;
import com.azure.resourcemanager.eventhubs.models.SchemaType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SchemaRegistry CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/SchemaRegistry/
     * SchemaRegistryCreate.json
     */
    /**
     * Sample code: SchemaRegistryCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void schemaRegistryCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getSchemaRegistries().createOrUpdateWithResponse("alitest",
            "ali-ua-test-eh-system-1", "testSchemaGroup1", new SchemaGroupInner().withGroupProperties(mapOf())
                .withSchemaCompatibility(SchemaCompatibility.FORWARD).withSchemaType(SchemaType.AVRO),
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
