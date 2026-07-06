
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileShareSnapshotsList.json
     */
    /**
     * Sample code: ListShareSnapshots.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listShareSnapshots(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().list("res9290", "sto1590", null, null, "snapshots",
            com.azure.core.util.Context.NONE);
    }
}
