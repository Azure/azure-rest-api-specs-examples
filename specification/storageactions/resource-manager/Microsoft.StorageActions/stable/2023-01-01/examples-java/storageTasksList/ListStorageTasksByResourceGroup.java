
/**
 * Samples for StorageTasks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-01-01/storageTasksList/ListStorageTasksByResourceGroup.json
     */
    /**
     * Sample code: ListStorageTasksByResourceGroup.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void
        listStorageTasksByResourceGroup(com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        manager.storageTasks().listByResourceGroup("res6117", com.azure.core.util.Context.NONE);
    }
}
