
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-15/Operations_List.json
     */
    /**
     * Sample code: List operations.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void listOperations(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
