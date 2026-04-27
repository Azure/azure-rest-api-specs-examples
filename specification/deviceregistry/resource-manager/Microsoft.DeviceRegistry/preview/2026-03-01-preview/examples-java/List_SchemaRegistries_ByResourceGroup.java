
/**
 * Samples for SchemaRegistries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/List_SchemaRegistries_ByResourceGroup.json
     */
    /**
     * Sample code: List_SchemaRegistries_ResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listSchemaRegistriesResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaRegistries().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
