
/**
 * Samples for DeletedAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/DeletedAccountList.json
     */
    /**
     * Sample code: DeletedAccountList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletedAccountList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getDeletedAccounts().list(com.azure.core.util.Context.NONE);
    }
}
