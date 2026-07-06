
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void operationsList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
