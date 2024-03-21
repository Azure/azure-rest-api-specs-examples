
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/misc/
     * OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void operationsList(com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
