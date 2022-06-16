import com.azure.core.util.Context;

/** Samples for FileShares List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/DeletedFileSharesList.json
     */
    /**
     * Sample code: ListDeletedShares.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDeletedShares(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .list("res9290", "sto1590", null, null, "deleted", Context.NONE);
    }
}
