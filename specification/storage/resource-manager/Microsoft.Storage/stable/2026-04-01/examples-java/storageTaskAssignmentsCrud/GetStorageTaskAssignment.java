
/**
 * Samples for StorageTaskAssignments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/storageTaskAssignmentsCrud/GetStorageTaskAssignment.json
     */
    /**
     * Sample code: GetStorageTaskAssignment.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getStorageTaskAssignment(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageTaskAssignments().getWithResponse("res4228", "sto4445", "myassignment1",
            com.azure.core.util.Context.NONE);
    }
}
