
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/FileSharesList.json
     */
    /**
     * Sample code: ListShares.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listShares(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().list("res9290", "sto1590", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
