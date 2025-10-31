
/**
 * Samples for SchemaRegistries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Get_SchemaRegistry.json
     */
    /**
     * Sample code: Get_SchemaRegistry.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getSchemaRegistry(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaRegistries().getByResourceGroupWithResponse("myResourceGroup", "my-schema-registry",
            com.azure.core.util.Context.NONE);
    }
}
