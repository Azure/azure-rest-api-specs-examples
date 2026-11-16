
/**
 * Samples for SchemaRegistry Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/SchemaRegistry/SchemaRegistryDelete.json
     */
    /**
     * Sample code: SchemaRegistryDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void schemaRegistryDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getSchemaRegistries().deleteWithResponse("alitest", "ali-ua-test-eh-system-1",
            "testSchemaGroup1", com.azure.core.util.Context.NONE);
    }
}
