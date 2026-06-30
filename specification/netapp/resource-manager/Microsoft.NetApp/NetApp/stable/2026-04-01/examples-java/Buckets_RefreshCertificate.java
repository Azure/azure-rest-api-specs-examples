
/**
 * Samples for Buckets RefreshCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/Buckets_RefreshCertificate.json
     */
    /**
     * Sample code: Buckets_RefreshCertificate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void bucketsRefreshCertificate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.buckets().refreshCertificate("myRG", "account1", "pool1", "volume1", "bucket1",
            com.azure.core.util.Context.NONE);
    }
}
