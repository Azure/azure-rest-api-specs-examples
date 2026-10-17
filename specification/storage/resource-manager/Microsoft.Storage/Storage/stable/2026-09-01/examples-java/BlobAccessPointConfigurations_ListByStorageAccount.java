
/**
 * Samples for BlobAccessPointConfigurations ListByStorageAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/BlobAccessPointConfigurations_ListByStorageAccount.json
     */
    /**
     * Sample code: ListBlobAccessPointConfigurationsByStorageAccount.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        listBlobAccessPointConfigurationsByStorageAccount(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobAccessPointConfigurations().listByStorageAccount("testrg", "teststorageaccount",
            com.azure.core.util.Context.NONE);
    }
}
