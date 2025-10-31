
/**
 * Samples for SchemaRegistries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_SchemaRegistries_BySubscription.json
     */
    /**
     * Sample code: List_SchemaRegistries_Subscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listSchemaRegistriesSubscription(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaRegistries().list(com.azure.core.util.Context.NONE);
    }
}
