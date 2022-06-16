import com.azure.core.util.Context;

/** Samples for FileShares List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/FileShareSnapshotsList.json
     */
    /**
     * Sample code: ListShareSnapshots.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listShareSnapshots(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .list("res9290", "sto1590", null, null, "snapshots", Context.NONE);
    }
}
