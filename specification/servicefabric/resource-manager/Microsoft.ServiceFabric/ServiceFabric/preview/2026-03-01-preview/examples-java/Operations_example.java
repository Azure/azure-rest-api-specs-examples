
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Operations_example.json
     */
    /**
     * Sample code: List available operations.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void listAvailableOperations(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
