
/**
 * Samples for Assets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_Assets_ResourceGroup.json
     */
    /**
     * Sample code: List_Assets_ResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listAssetsResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assets().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
