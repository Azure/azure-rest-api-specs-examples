
/**
 * Samples for Assets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Get_Asset_With_SyncStatus.json
     */
    /**
     * Sample code: Get_Asset_With_SyncStatus.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getAssetWithSyncStatus(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assets().getByResourceGroupWithResponse("myResourceGroup", "my-asset",
            com.azure.core.util.Context.NONE);
    }
}
