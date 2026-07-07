
/**
 * Samples for StorageTaskAssignments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/storageTaskAssignmentsCrud/DeleteStorageTaskAssignment.json
     */
    /**
     * Sample code: DeleteStorageTaskAssignment.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteStorageTaskAssignment(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageTaskAssignments().delete("res4228", "sto4445", "myassignment1",
            com.azure.core.util.Context.NONE);
    }
}
