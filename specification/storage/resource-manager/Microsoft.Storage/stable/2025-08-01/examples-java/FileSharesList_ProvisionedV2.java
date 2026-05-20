
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileSharesList_ProvisionedV2.json
     */
    /**
     * Sample code: ListSharesProvisionedV2.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listSharesProvisionedV2(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().list("res9290", "sto1590", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
