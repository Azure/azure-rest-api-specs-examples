
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesList_PaidBursting.
     * json
     */
    /**
     * Sample code: ListSharesPaidBursting.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSharesPaidBursting(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().list("res9290", "sto1590", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
