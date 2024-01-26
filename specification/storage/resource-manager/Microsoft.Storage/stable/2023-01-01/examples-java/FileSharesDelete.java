
/** Samples for FileShares Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/FileSharesDelete.json
     */
    /**
     * Sample code: DeleteShares.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteShares(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().deleteWithResponse("res4079", "sto4506",
            "share9689", null, null, com.azure.core.util.Context.NONE);
    }
}
