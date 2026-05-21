
/**
 * Samples for StorageTaskAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/storageTaskAssignmentsList/ListStorageTaskAssignmentsForAccount.json
     */
    /**
     * Sample code: ListStorageTaskAssignmentsForAccount.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listStorageTaskAssignmentsForAccount(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageTaskAssignments().list("res4228", "sto4445", null,
            com.azure.core.util.Context.NONE);
    }
}
