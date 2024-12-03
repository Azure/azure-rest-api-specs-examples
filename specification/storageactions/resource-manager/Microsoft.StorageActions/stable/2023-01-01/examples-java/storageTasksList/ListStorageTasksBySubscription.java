
/**
 * Samples for StorageTasks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/
     * storageTasksList/ListStorageTasksBySubscription.json
     */
    /**
     * Sample code: ListStorageTasksBySubscription.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void
        listStorageTasksBySubscription(com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        manager.storageTasks().list(com.azure.core.util.Context.NONE);
    }
}
