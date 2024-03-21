
/**
 * Samples for StorageTasks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/
     * storageTasksCrud/DeleteStorageTask.json
     */
    /**
     * Sample code: DeleteStorageTask.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void deleteStorageTask(com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        manager.storageTasks().delete("res4228", "mytask1", com.azure.core.util.Context.NONE);
    }
}
