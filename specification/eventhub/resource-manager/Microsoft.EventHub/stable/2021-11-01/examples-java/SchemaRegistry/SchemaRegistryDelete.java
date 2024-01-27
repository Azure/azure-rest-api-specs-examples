
import com.azure.core.util.Context;

/** Samples for SchemaRegistry Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/SchemaRegistry/
     * SchemaRegistryDelete.json
     */
    /**
     * Sample code: SchemaRegistryDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void schemaRegistryDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getSchemaRegistries().deleteWithResponse("alitest",
            "ali-ua-test-eh-system-1", "testSchemaGroup1", Context.NONE);
    }
}
