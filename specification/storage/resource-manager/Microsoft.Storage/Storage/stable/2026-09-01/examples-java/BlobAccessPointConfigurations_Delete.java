
/**
 * Samples for BlobAccessPointConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/BlobAccessPointConfigurations_Delete.json
     */
    /**
     * Sample code: DeleteBlobAccessPointConfiguration.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteBlobAccessPointConfiguration(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobAccessPointConfigurations().delete("testrg", "teststorageaccount",
            "testaccesspointconfig", com.azure.core.util.Context.NONE);
    }
}
