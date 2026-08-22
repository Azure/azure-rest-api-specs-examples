
/**
 * Samples for PrivateLinkResources ListByStorageSyncService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/PrivateLinkResources_List.json
     */
    /**
     * Sample code: PrivateLinkResources_List.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void privateLinkResourcesList(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.privateLinkResources().listByStorageSyncServiceWithResponse("res6977", "sss2527",
            com.azure.core.util.Context.NONE);
    }
}
