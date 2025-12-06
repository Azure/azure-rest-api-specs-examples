
/**
 * Samples for Caches ListByCapacityPools.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Caches_ListByCapacityPools.json
     */
    /**
     * Sample code: Caches_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().listByCapacityPools("myRG", "account1", "pool1", com.azure.core.util.Context.NONE);
    }
}
