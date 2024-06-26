
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/FileShareSnapshotsList.json
     */
    /**
     * Sample code: ListShareSnapshots.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listShareSnapshots(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().list("res9290", "sto1590", null, null,
            "snapshots", com.azure.core.util.Context.NONE);
    }
}
