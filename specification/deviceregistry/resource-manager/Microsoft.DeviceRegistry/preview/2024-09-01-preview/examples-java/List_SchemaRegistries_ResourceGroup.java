
/**
 * Samples for SchemaRegistries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_SchemaRegistries_ResourceGroup.json
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
