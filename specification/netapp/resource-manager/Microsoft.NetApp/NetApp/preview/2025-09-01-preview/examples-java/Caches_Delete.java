
/**
 * Samples for Caches Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Caches_Delete.json
     */
    /**
     * Sample code: Caches_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().delete("myRG", "account1", "pool1", "cache1", com.azure.core.util.Context.NONE);
    }
}
