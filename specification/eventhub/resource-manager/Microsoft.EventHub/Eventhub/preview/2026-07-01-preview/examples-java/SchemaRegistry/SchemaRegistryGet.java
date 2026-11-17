
/**
 * Samples for SchemaRegistry Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/SchemaRegistry/SchemaRegistryGet.json
     */
    /**
     * Sample code: SchemaRegistryGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void schemaRegistryGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getSchemaRegistries().getWithResponse("alitest", "ali-ua-test-eh-system-1",
            "testSchemaGroup1", com.azure.core.util.Context.NONE);
    }
}
