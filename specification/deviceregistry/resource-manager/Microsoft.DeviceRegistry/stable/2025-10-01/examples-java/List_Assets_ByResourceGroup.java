
/**
 * Samples for Assets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_Assets_ByResourceGroup.json
     */
    /**
     * Sample code: List_Assets_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listAssetsByResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assets().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
