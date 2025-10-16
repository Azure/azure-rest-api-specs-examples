
/**
 * Samples for Buckets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/Buckets_Delete.json
     */
    /**
     * Sample code: Buckets_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void bucketsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.buckets().delete("myRG", "account1", "pool1", "volume1", "bucket1", com.azure.core.util.Context.NONE);
    }
}
