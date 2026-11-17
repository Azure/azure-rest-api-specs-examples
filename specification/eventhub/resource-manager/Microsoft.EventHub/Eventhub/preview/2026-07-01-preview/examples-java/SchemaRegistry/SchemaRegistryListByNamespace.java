
/**
 * Samples for SchemaRegistry ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/SchemaRegistry/SchemaRegistryListByNamespace.json
     */
    /**
     * Sample code: SchemaRegistryListAll.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void schemaRegistryListAll(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getSchemaRegistries().listByNamespace("alitest", "ali-ua-test-eh-system-1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
