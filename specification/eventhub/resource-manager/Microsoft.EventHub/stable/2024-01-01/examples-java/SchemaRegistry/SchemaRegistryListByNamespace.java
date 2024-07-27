
/**
 * Samples for SchemaRegistry ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/SchemaRegistry/
     * SchemaRegistryListByNamespace.json
     */
    /**
     * Sample code: SchemaRegistryListAll.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void schemaRegistryListAll(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getSchemaRegistries().listByNamespace("alitest",
            "ali-ua-test-eh-system-1", null, null, com.azure.core.util.Context.NONE);
    }
}
