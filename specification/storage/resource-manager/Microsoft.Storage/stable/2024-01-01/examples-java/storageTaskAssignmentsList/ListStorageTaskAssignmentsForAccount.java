
/**
 * Samples for StorageTaskAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/storageTaskAssignmentsList/
     * ListStorageTaskAssignmentsForAccount.json
     */
    /**
     * Sample code: ListStorageTaskAssignmentsForAccount.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listStorageTaskAssignmentsForAccount(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageTaskAssignments().list("res4228", "sto4445", null,
            com.azure.core.util.Context.NONE);
    }
}
