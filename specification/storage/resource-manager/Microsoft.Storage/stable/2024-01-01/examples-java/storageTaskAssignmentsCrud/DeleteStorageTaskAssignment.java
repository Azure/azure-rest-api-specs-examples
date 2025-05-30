
/**
 * Samples for StorageTaskAssignments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/storageTaskAssignmentsCrud/
     * DeleteStorageTaskAssignment.json
     */
    /**
     * Sample code: DeleteStorageTaskAssignment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteStorageTaskAssignment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageTaskAssignments().delete("res4228", "sto4445",
            "myassignment1", com.azure.core.util.Context.NONE);
    }
}
