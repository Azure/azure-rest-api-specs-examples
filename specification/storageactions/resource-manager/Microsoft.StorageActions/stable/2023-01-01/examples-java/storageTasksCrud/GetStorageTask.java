
/**
 * Samples for StorageTasks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-01-01/storageTasksCrud/GetStorageTask.json
     */
    /**
     * Sample code: GetStorageTask.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void getStorageTask(com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        manager.storageTasks().getByResourceGroupWithResponse("res4228", "mytask1", com.azure.core.util.Context.NONE);
    }
}
