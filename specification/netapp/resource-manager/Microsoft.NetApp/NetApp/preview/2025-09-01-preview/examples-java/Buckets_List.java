
/**
 * Samples for Buckets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Buckets_List.json
     */
    /**
     * Sample code: Buckets_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void bucketsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.buckets().list("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
