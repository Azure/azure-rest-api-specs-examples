
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/FileSharesList_ProvisionedV2.
     * json
     */
    /**
     * Sample code: ListSharesProvisionedV2.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSharesProvisionedV2(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().list("res9290", "sto1590", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
