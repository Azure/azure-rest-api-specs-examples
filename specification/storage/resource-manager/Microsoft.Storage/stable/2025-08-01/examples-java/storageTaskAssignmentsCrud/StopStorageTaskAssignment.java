
/**
 * Samples for StorageTaskAssignments StopAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/storageTaskAssignmentsCrud/StopStorageTaskAssignment.json
     */
    /**
     * Sample code: StopStorageTaskAssignment.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void stopStorageTaskAssignment(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageTaskAssignments().stopAssignment("res4228", "sto4445", "myassignment1",
            com.azure.core.util.Context.NONE);
    }
}
