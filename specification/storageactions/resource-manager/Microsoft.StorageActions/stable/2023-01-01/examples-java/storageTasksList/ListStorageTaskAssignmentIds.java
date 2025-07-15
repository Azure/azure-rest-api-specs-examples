
/**
 * Samples for StorageTaskAssignment List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-01-01/storageTasksList/ListStorageTaskAssignmentIds.json
     */
    /**
     * Sample code: ListStorageTaskAssignmentsByResourceGroup.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void listStorageTaskAssignmentsByResourceGroup(
        com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        manager.storageTaskAssignments().list("rgroup1", "mytask1", null, com.azure.core.util.Context.NONE);
    }
}
