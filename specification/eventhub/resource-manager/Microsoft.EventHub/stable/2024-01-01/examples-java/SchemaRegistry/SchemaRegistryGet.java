
/**
 * Samples for SchemaRegistry Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/SchemaRegistry/
     * SchemaRegistryGet.json
     */
    /**
     * Sample code: SchemaRegistryGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void schemaRegistryGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getSchemaRegistries().getWithResponse("alitest",
            "ali-ua-test-eh-system-1", "testSchemaGroup1", com.azure.core.util.Context.NONE);
    }
}
