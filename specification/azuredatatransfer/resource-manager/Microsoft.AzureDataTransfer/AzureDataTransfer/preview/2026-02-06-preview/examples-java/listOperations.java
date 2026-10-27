
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/listOperations.json
     */
    /**
     * Sample code: List operations.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void listOperations(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
