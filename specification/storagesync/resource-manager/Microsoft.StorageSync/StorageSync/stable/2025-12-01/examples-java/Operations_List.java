
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void operationsList(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
