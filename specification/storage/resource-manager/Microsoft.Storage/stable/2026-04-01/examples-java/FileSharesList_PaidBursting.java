
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileSharesList_PaidBursting.json
     */
    /**
     * Sample code: ListSharesPaidBursting.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listSharesPaidBursting(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().list("res9290", "sto1590", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
