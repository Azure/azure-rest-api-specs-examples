
/**
 * Samples for Caches List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/Caches_List.json
     */
    /**
     * Sample code: Caches_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().list("myRG", "account1", "pool1", com.azure.core.util.Context.NONE);
    }
}
